# 在 requestEth 已监听后执行：快速检查「扫链 + 同步」相关 HTTP 是否正常（仅看状态码与简短 body）
# 用法：.\scripts\verify-api-flow.ps1  或  $env:API_BASE = "http://127.0.0.1:8003"; .\scripts\verify-api-flow.ps1

$Base = if ($env:API_BASE) { $env:API_BASE.TrimEnd("/") } else { "http://127.0.0.1:8003" }
$long = 120
$syncMint = 300

function Test-Get([string]$Path, [int]$MaxSec = 60) {
  $url = "$Base$Path"
  try {
    $r = Invoke-WebRequest -Uri $url -Method GET -UseBasicParsing -TimeoutSec $MaxSec
    $code = [int]$r.StatusCode
    $ok = $code -ge 200 -and $code -lt 300
    $snippet = if ($r.Content.Length -gt 200) { $r.Content.Substring(0, 200) + "..." } else { $r.Content }
    [PSCustomObject]@{ Path = $Path; Code = $code; Ok = $ok; Snippet = $snippet }
  } catch {
    [PSCustomObject]@{ Path = $Path; Code = -1; Ok = $false; Snippet = $_.Exception.Message }
  }
}

Write-Host "API_BASE=$Base" -ForegroundColor Cyan
$rows = @()
$rows += Test-Get "/health" 5
# 与 sync-events-fast + sync-events-rest 对齐的关键路径
$rows += Test-Get "/api/debug_mint_sync" $long
$rows += Test-Get "/api/sync_mint_from_contract" $syncMint
$rows += Test-Get "/api/get_box_mint_event" $long
$rows += Test-Get "/api/get_box_transfer_event" $long
$rows += Test-Get "/api/get_box_open_event" $long
$rows += Test-Get "/api/update_box" $long
$rows += Test-Get "/api/get_box_list_event" $long
$rows += Test-Get "/api/get_box_unlist_event" $long
$rows += Test-Get "/api/get_box_buy_event" $long
$rows += Test-Get "/api/get_user_r_event" $long
$rows += Test-Get "/api/get_exchange_event" $long
# 常用查询
$rows += Test-Get "/api/get_address_box?address=0x0000000000000000000000000000000000000000" $long
$rows | Format-Table -AutoSize
$bad = $rows | Where-Object { -not $_.Ok }
if ($bad) {
  Write-Host "部分请求失败。若数据库/RPC/合约未配置，属预期；修正 REQUESTETH_DATABASE_SOURCE 与 .env 后重试。" -ForegroundColor Yellow
  exit 1
}
Write-Host "全部返回 2xx。" -ForegroundColor Green
