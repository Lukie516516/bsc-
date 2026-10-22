从 requestEth_ZTKtx.tar.gz 解压后说明（2026-04-21）

1) docker-compose.yml 里写的是：env_file: .env
   因此：只改 .env.fronttest 时，Docker 里的 api/worker 不会读到，除非：
   - 把 .env.fronttest 复制覆盖为 .env：  cp .env.fronttest .env
   - 或改 docker-compose.yml 为 env_file: .env.fronttest 后再 compose up

2) 你原来的 .env.fronttest 第 21 行把两行粘在一起，已在本目录副本里拆开为：
   ETH_AICAT_TOKEN_CONTRACT=...
   ETH_DEPLOY_BLOCK_NFT_LISTED=...

3) 若要与本机 nft-blind 仓库「新部署」一致，合约与块高见 docs/前端合约地址-BlindBoxMainnetTest.md，勿与旧地址混用。
