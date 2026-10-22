# requestEth: API + Workers
# 默认运行 API，覆盖 CMD 可运行 Worker，例如：
#   docker run ... your-image ./blind_settle_worker -conf /data/conf
#
# 说明：不在镜像内 apt install protoc —— api 下 *.pb.go 已提交仓库，本地改 proto 后执行 make api 再构建即可。
# 避免部分服务器上 deb.debian.org 出现 “invalid signature / not signed” 导致 docker build 失败。

FROM golang:1.22 AS builder

WORKDIR /src

ENV GOPROXY=https://goproxy.cn,direct
# 镜像内构建不依赖宿主机 Git；避免 dubious ownership / 无 .git 时 go build 因 VCS stamping 失败
ENV GOFLAGS=-buildvcs=false

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/ ./cmd/...

# 运行阶段用 alpine + apk，减少对 Debian apt 源的依赖（与 builder 中 apt-get 报错环境解耦）
FROM alpine:3.19

RUN apk add --no-cache ca-certificates wget

COPY --from=builder /src/bin /app
WORKDIR /app
RUN mkdir -p /log

EXPOSE 8000 9000
VOLUME ["/data/conf"]

CMD ["./requestEth", "-conf", "/data/conf"]
