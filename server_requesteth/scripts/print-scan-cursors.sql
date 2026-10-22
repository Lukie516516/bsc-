-- 各链扫块游标（与表 Get*Last: ORDER BY block_number DESC, id DESC LIMIT 1 一致）
USE game_fronttest;

SELECT 'swap_trade'            AS scan_type, (SELECT block_number FROM swap_trade            ORDER BY block_number DESC, id DESC LIMIT 1) AS last_block
UNION ALL SELECT 'primary_buy',            (SELECT block_number FROM primary_buy            ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'primary_sell',           (SELECT block_number FROM primary_sell           ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'reward_notified',      (SELECT block_number FROM reward_notified        ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_market_purchase',  (SELECT block_number FROM nft_market_purchase   ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_minted',            (SELECT block_number FROM nft_minted              ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_market_listed',     (SELECT block_number FROM nft_market_listed     ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_market_unlisted',  (SELECT block_number FROM nft_market_unlisted   ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_opened',           (SELECT block_number FROM nft_opened             ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'nft_transfer',         (SELECT block_number FROM nft_transfer            ORDER BY block_number DESC, id DESC LIMIT 1)
UNION ALL SELECT 'user_registered',      (SELECT block_number FROM user_registered         ORDER BY block_number DESC, id DESC LIMIT 1);
