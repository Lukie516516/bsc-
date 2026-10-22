# Windows PowerShell：模拟前端关键接口自检
# 用法：
#   .\scripts\smoke-test.ps1
#   $env:API_BASE="http://127.0.0.1:8003"; $env:ADDR="0x你的地址"; .\scripts\smoke-test.ps1

$ErrorActionPreference = "Stop"
$API_BASE = if ($env:API_BASE) { $env:API_BASE } else { "http://127.0.0.1:8003" }
$ADDR = if ($env:ADDR) { $env:ADDR } else { "0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069" }

Write-Host "API_BASE=$API_BASE"
Write-Host "ADDR=$ADDR"

function Get-Json($url) {
    $r = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 60
    return @{ Status = $r.StatusCode; Body = ($r.Content | ConvertFrom-Json) }
}

Write-Host "`n[1] /health"
try {
    $h = Invoke-WebRequest -Uri "$API_BASE/health" -UseBasicParsing -TimeoutSec 10
    Write-Host "HTTP $($h.StatusCode)" $h.Content
} catch { Write-Host "FAIL" $_.Exception.Message }

Write-Host "`n[2] /api/get_address_box"
$url = "$API_BASE/api/get_address_box?address=$ADDR&page=1&numTwo=0&numThree=1&openStatus=0&num=0"
try {
    $r = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 60
    $j = $r.Content | ConvertFrom-Json
    $listLen = @($j.list).Count
    $cnt = [int]$j.count
    $ids = @($j.list | ForEach-Object { $_.tokenId })
    $uniq = ($ids | Select-Object -Unique).Count
    Write-Host "HTTP $($r.StatusCode) count=$cnt list_len=$listLen unique_tokenIds=$uniq"
    if ($cnt -ne $listLen) { Write-Host "FAIL: count 与 list 条数不一致" }
    if ($listLen -gt 0 -and $uniq -ne $listLen) { Write-Host "FAIL: 存在重复 tokenId" }
    if ($cnt -eq $listLen -and ($listLen -eq 0 -or $uniq -eq $listLen)) { Write-Host "OK" }
} catch { Write-Host "FAIL" $_.Exception.Message }

Write-Host "`n[3] /api/debug_mint_sync"
try {
    $r = Invoke-WebRequest -Uri "$API_BASE/api/debug_mint_sync" -UseBasicParsing -TimeoutSec 120
    Write-Host $r.Content
} catch { Write-Host "FAIL" $_.Exception.Message }

Write-Host "`n[4] /api/sync_mint_from_contract"
try {
    $r = Invoke-WebRequest -Uri "$API_BASE/api/sync_mint_from_contract" -UseBasicParsing -TimeoutSec 300
    Write-Host $r.Content
} catch { Write-Host "FAIL" $_.Exception.Message }
