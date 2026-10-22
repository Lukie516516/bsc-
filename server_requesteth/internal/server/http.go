package server

import (
	"context"
	nethttp "net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/transport/http"
	transaction "requestEth/api/requestEth/v1"
	"requestEth/internal/conf"
	"requestEth/internal/service"

	"github.com/gorilla/handlers"
)

// healthFilter 健康检查：/health 直接返回 200，不经过限流和业务逻辑
func healthFilter(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		if r.URL.Path == "/health" || r.URL.Path == "/health/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(nethttp.StatusOK)
			_, _ = w.Write([]byte(`{"status":"ok"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, transactionService *service.TransactionService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(healthFilter),           // 健康检查优先
		http.Middleware(ratelimit.Server()), // HTTP 限流（BBR 算法）
		http.Filter(handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// 先注册 /health，确保优先匹配（Filter 与路由双重保障）
	r := srv.Route("/")
	r.GET("/health", func(ctx http.Context) error {
		return ctx.JSON(200, map[string]string{"status": "ok"})
	})
	r.GET("/health/", func(ctx http.Context) error {
		return ctx.JSON(200, map[string]string{"status": "ok"})
	})

	// 调试：铸造事件拉取诊断（排查 nft_minted 为空）
	r.GET("/api/debug_mint_sync", func(ctx http.Context) error {
		res := transactionService.DebugMintSync(context.Background())
		return ctx.JSON(200, res)
	})
	// 从合约 view 同步铸造数据，绕过 eth_getLogs 限制（FilterLogs limit exceeded）
	r.GET("/api/sync_mint_from_contract", func(ctx http.Context) error {
		bgCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		res := transactionService.SyncMintedFromContractHandler(bgCtx)
		return ctx.JSON(200, res)
	})
	r.GET("/api/quant_fund_dashboard", func(ctx http.Context) error {
		limit := 20
		if raw := ctx.Request().URL.Query().Get("limit"); raw != "" {
			if v, err := strconv.Atoi(raw); err == nil && v > 0 {
				limit = v
			}
		}
		res := transactionService.QuantFundDashboardHandler(ctx.Request().Context(), limit)
		return ctx.JSON(200, res)
	})
	r.GET("/api/address_box_v2", func(ctx http.Context) error {
		query := ctx.Request().URL.Query()
		req := &transaction.GetAddressBoxRequest{
			Address:    query.Get("address"),
			Page:       uint64(max(1, parseIntWithDefault(query.Get("page"), 1))),
			Num:        uint64(max(0, parseIntWithDefault(query.Get("num"), 2))),
			NumTwo:     uint64(max(0, parseIntWithDefault(query.Get("numTwo"), 0))),
			NumThree:   uint64(max(0, parseIntWithDefault(query.Get("numThree"), 0))),
			NumFour:    uint64(max(0, parseIntWithDefault(query.Get("numFour"), 0))),
			OpenStatus: uint64(max(0, parseIntWithDefault(query.Get("openStatus"), 2))),
			Start:      service.ParseUint64Query(query, "start"),
			End:        service.ParseUint64Query(query, "end"),
		}
		res := transactionService.AddressBoxV2Handler(ctx.Request().Context(), req.Address, req)
		return ctx.JSON(200, res)
	})
	r.GET("/api/user_profit_summary", func(ctx http.Context) error {
		query := ctx.Request().URL.Query()
		res := transactionService.UserProfitSummaryHandler(
			ctx.Request().Context(),
			query.Get("address"),
			service.ParseUint64Query(query, "start"),
			service.ParseUint64Query(query, "end"),
		)
		return ctx.JSON(200, res)
	})
	r.GET("/api/user_relation_summary", func(ctx http.Context) error {
		query := ctx.Request().URL.Query()
		res := transactionService.UserRelationSummaryHandler(ctx.Request().Context(), query.Get("address"))
		return ctx.JSON(200, res)
	})
	r.GET("/api/runtime_config", func(ctx http.Context) error {
		query := ctx.Request().URL.Query()
		res := transactionService.RuntimeConfigListHandler(ctx.Request().Context(), query.Get("prefix"))
		return ctx.JSON(200, res)
	})
	r.GET("/api/runtime_config_upsert", func(ctx http.Context) error {
		query := ctx.Request().URL.Query()
		if !isAdminAuthorized(query.Get("admin_key")) {
			return ctx.JSON(403, map[string]interface{}{
				"code":    403,
				"message": "forbidden",
			})
		}
		res := transactionService.RuntimeConfigUpsertHandler(
			ctx.Request().Context(),
			query.Get("key"),
			query.Get("value"),
			query.Get("description"),
		)
		return ctx.JSON(200, res)
	})

	transaction.RegisterTransactionHTTPServer(srv, transactionService)

	return srv
}

func parseIntWithDefault(raw string, fallback int) int {
	if raw == "" {
		return fallback
	}
	v, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAdminAuthorized(input string) bool {
	expected := os.Getenv("REQUESTETH_ADMIN_API_KEY")
	if expected == "" {
		return false
	}
	return input == expected
}
