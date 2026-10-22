# Compare SHA-256 of runtime bytecode: BSC vs local Hardhat artifact.
# Usage: cd box; .\scripts\compare-onchain-local.ps1
# Requires: PowerShell 5+

$ErrorActionPreference = "Stop"
$rpc = "https://bsc-dataseed1.binance.org"

function Get-CodeSha256([string]$address) {
  $body = @{
    jsonrpc = "2.0"
    id      = 1
    method  = "eth_getCode"
    params  = @($address, "latest")
  } | ConvertTo-Json -Compress
  $r = Invoke-RestMethod -Uri $rpc -Method Post -Body $body -ContentType "application/json" -TimeoutSec 60
  $code = $r.result
  if (-not $code -or $code -eq "0x") { return @{ len = 0; sha = $null; raw = $code } }
  $hex = $code.TrimStart("0x")
  $bytes = New-Object byte[] ($hex.Length / 2)
  for ($i = 0; $i -lt $hex.Length; $i += 2) {
    $bytes[$i / 2] = [Convert]::ToByte($hex.Substring($i, 2), 16)
  }
  $sha = [System.BitConverter]::ToString([System.Security.Cryptography.SHA256]::Create().ComputeHash($bytes)).Replace("-", "").ToLowerInvariant()
  return @{ len = $bytes.Length; sha = $sha; raw = $code }
}

function Get-ArtifactSha256([string]$artifactPath) {
  $full = Join-Path (Join-Path $PSScriptRoot "..") $artifactPath | Resolve-Path
  $j = Get-Content $full -Raw | ConvertFrom-Json
  $hex = $j.deployedBytecode.TrimStart("0x")
  if ($hex.Length -eq 0) { return @{ len = 0; sha = $null } }
  if ($hex -match "[^0-9a-fA-F]") {
    # Hardhat library placeholders (__$...$__) — cannot hash locally; compare via BscScan or eth_getCode only.
    return @{ len = -1; sha = $null }
  }
  $bytes = New-Object byte[] ($hex.Length / 2)
  for ($i = 0; $i -lt $hex.Length; $i += 2) {
    $bytes[$i / 2] = [Convert]::ToByte($hex.Substring($i, 2), 16)
  }
  $sha = [System.BitConverter]::ToString([System.Security.Cryptography.SHA256]::Create().ComputeHash($bytes)).Replace("-", "").ToLowerInvariant()
  return @{ len = $bytes.Length; sha = $sha }
}

# BlindBoxMainnetTest 2026-04-20 部署（与 box/ignition/deployments/chain-56/deployed_addresses.json 一致）
$pairs = @(
  @{ name = "MockUSDT"; addr = "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4"; art = "artifacts/contracts/MockUSDT.sol/MockUSDT.json" }
  @{ name = "Piece"; addr = "0xe5638DCDbE52E3AD3D6b253713Ea66637bb63313"; art = "artifacts/contracts/Piece.sol/Piece.json" }
  @{ name = "Hrx"; addr = "0xE3198bf629C91bd678190D22393Ba5DEc40E9106"; art = "artifacts/contracts/Hrx.sol/Hrx.json" }
  @{ name = "BondingCurvePrimaryMarket"; addr = "0xE9E4C6c8Eb8b80420D94475C7278c693dd70b41a"; art = "artifacts/contracts/Market.sol/BondingCurvePrimaryMarket.json" }
  @{ name = "UmbrellaStakeDividend"; addr = "0xBd86DD7a5eAdAEbF7E7Fa1a258Dc39318b364f00"; art = "artifacts/contracts/Stake.sol/UmbrellaStakeDividend.json" }
  @{ name = "BlindBoxNFT"; addr = "0x73F76b11ebF03718B71c163BE2B5445ECFC25Ce6"; art = "artifacts/contracts/Nft.sol/BlindBoxNFT.json" }
  @{ name = "SwapHrxUsdt"; addr = "0x3Fc7A08C7af8B8300a6C56475cc32c72534eE7E4"; art = "artifacts/contracts/SwapHrxUsdt.sol/SwapHrxUsdt.json" }
  @{ name = "DealRouter"; addr = "0x552F6Fd1c1725AEe50dA39922507D7687c9D3112"; art = "artifacts/contracts/Router.sol/DealRouter.json" }
  @{ name = "BlindBoxPrincipalVault"; addr = "0x1fcEfea5d9Ac20C40441793d2B8BcbE22694891C"; art = "artifacts/contracts/StakeTwo.sol/BlindBoxPrincipalVault.json" }
  @{ name = "BlindBoxVaultFactory"; addr = "0x52FE50ae93CFE985032BE5059cc60d3b6348c1dc"; art = "artifacts/contracts/FactoryList.sol/BlindBoxVaultFactory.json" }
  @{ name = "BlindBoxUserVault"; addr = "0xaf08c8f214d0b7A74c3620EB432789f291402f82"; art = "artifacts/contracts/ListTemplate.sol/BlindBoxUserVault.json" }
)

$root = Split-Path $PSScriptRoot -Parent
Set-Location $root

Write-Host "Comparing on-chain vs local artifacts (SHA-256 of deployed bytecode)`n"

foreach ($p in $pairs) {
  $chain = Get-CodeSha256 $p.addr
  $local = Get-ArtifactSha256 $p.art
  if ($local.len -lt 0) {
    Write-Host ("{0,-28} SKIP-ARTIFACT  chainBytes={1}  (local artifact has unlinked library placeholders)" -f $p.name, $chain.len)
    Write-Host ("  address: {0}" -f $p.addr)
    continue
  }
  $match = $chain.sha -and $local.sha -and ($chain.sha -eq $local.sha)
  $status = if ($match) { "MATCH" } else { "DIFF" }
  Write-Host ("{0,-28} {1}  chainBytes={2} localBytes={3}" -f $p.name, $status, $chain.len, $local.len)
  if (-not $match) {
    Write-Host ("  chain sha256: {0}" -f $chain.sha)
    Write-Host ("  local sha256: {0}" -f $local.sha)
    Write-Host ("  address: {0}" -f $p.addr)
  }
}
