# XAMPP MySQL quick check. Run: .\scripts\inspect-mysql-xampp.ps1
$ErrorActionPreference = "Continue"
$mysql = "C:\xampp\mysql\bin\mysql.exe"
$pma  = "C:\xampp\phpMyAdmin\config.inc.php"

Write-Host "=== process / port 3306 ===" -ForegroundColor Cyan
Get-Process mysqld -ErrorAction SilentlyContinue | Format-Table Id, Path -AutoSize
netstat -an | findstr "LISTEN" | findstr ":3306"

Write-Host ""
Write-Host "=== phpMyAdmin config.inc.php (default only; may not match real MySQL) ===" -ForegroundColor Cyan
if (Test-Path $pma) {
  $txt = Get-Content $pma -Raw
  if ($txt -match "\['password'\]\s*=\s*'([^']*)'") { Write-Host "File says root password length: $($matches[1].Length) (0=empty in file)" }
  if ($txt -match "\['user'\]\s*=\s*'([^']+)'") { Write-Host "File says user: $($matches[1])" }
} else { Write-Host "Missing: $pma" }

Write-Host ""
Write-Host "=== try root, empty password (mysql CLI) ===" -ForegroundColor Cyan
$null = & $mysql -u root -e "SELECT 1" 2>&1
if ($LASTEXITCODE -eq 0) { Write-Host "root empty password: OK" -ForegroundColor Green }
else { Write-Host "root empty password: 1045 denied" -ForegroundColor Yellow }

$projRoot = Split-Path -Parent $PSScriptRoot
Write-Host ""
Write-Host "=== DSN hint files in repo (open in editor) ===" -ForegroundColor Cyan
foreach ($name in @("configs\config.yaml", "configs\config.fronttest.yaml", "configs\xampp\config.yaml")) {
  $f = Join-Path $projRoot $name
  if (Test-Path $f) { Write-Host "  $f" }
}

Write-Host ""
Write-Host "=== next steps (manual) ===" -ForegroundColor Cyan
Write-Host "1) http://127.0.0.1/phpmyadmin  If login works, open Users to see root / game_fronttest."
Write-Host "2) If denied everywhere: XAMPP stop MySQL, then web search: XAMPP reset MySQL root password skip grant tables"
Write-Host "3) As root, run: scripts\sql-init-xampp.sql  (creates game_fronttest / 1234)"
Write-Host ""
Write-Host "If 1045 on root here but phpMyAdmin still works, config.inc may be out of date; you set a root password later."
