# Run from requestEth project root: XAMPP MySQL + BSC (see .env.fronttest for contracts)
# DSN: configs\xampp\config.yaml 与 config.fronttest 一致 (game_fronttest/1234)。1045 时在 phpMyAdmin 执行 scripts\sql-init-xampp.sql
# Optional: $env:REQUESTETH_DATABASE_SOURCE = '...'  (overrides yaml)

$ErrorActionPreference = "Stop"
# REQUESTETH_DATABASE_SOURCE 若在本终端里仍被设为旧 DSN，会完全覆盖 config.yaml；要用 yaml 时先: Remove-Item Env:REQUESTETH_DATABASE_SOURCE
if (Test-Path Env:REQUESTETH_DATABASE_SOURCE) {
  Write-Host "Using REQUESTETH_DATABASE_SOURCE (overrides configs\xampp\config.yaml). To use yaml DSN only, run: Remove-Item Env:REQUESTETH_DATABASE_SOURCE" -ForegroundColor DarkYellow
}
$root = Split-Path -Parent $PSScriptRoot
Set-Location $root
if (-not (Test-Path (Join-Path $root "go.mod"))) {
  Write-Host "Run this script from requestEth project root." -ForegroundColor Red
  exit 1
}
New-Item -ItemType Directory -Force -Path "log" | Out-Null

$envFront = Join-Path $root ".env.fronttest"
if (Test-Path $envFront) {
  foreach ($line in Get-Content -LiteralPath $envFront) {
    $t = $line.Trim()
    if ($t.Length -eq 0) { continue }
    if ($t.StartsWith("#")) { continue }
    $eq = $t.IndexOf("=")
    if ($eq -lt 1) { continue }
    $k = $t.Substring(0, $eq).Trim()
    $v = $t.Substring($eq + 1).Trim()
    [Environment]::SetEnvironmentVariable($k, $v, "Process")
  }
  Write-Host "Loaded .env.fronttest" -ForegroundColor Green
} else {
  Write-Host "No .env.fronttest; set ETH_* env vars yourself." -ForegroundColor Yellow
}

if (-not $env:ETH_BSC_RPC_URLS) {
  $env:ETH_BSC_RPC_URLS = "https://bsc-dataseed1.binance.org/,https://bsc-dataseed2.binance.org/,https://1rpc.io/bnb"
}
if (-not $env:ETH_BSC_RPC_URLS_POLLING) {
  $env:ETH_BSC_RPC_URLS_POLLING = $env:ETH_BSC_RPC_URLS
}

$confXampp = [System.IO.Path]::GetFullPath((Join-Path $root "configs\xampp"))
Write-Host "Starting requestEth  -conf $confXampp  (Ctrl+C to stop)" -ForegroundColor Green
go run ./cmd/requestEth -conf $confXampp
