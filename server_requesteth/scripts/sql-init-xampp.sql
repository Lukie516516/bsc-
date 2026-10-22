-- =============================================================================
-- 在 phpMyAdmin 用【能进「用户」页面的账号，一般是 root】→ 点「SQL」整段执行。
-- 若仍 1045：说明本机里要么没有 game_fronttest，要么密码不是 1234。下面会删掉再建。
-- =============================================================================

CREATE DATABASE IF NOT EXISTS game_fronttest
  CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 看当前是否已有该用户（可对照 phpMyAdmin「用户」列表）
-- SELECT user, host, plugin FROM mysql.user WHERE user = 'game_fronttest';

-- MariaDB 10.1+ / MySQL 5.7+ 支持 IF EXISTS；老版本报错时去掉 IF EXISTS 一句句执行，或到「用户」里手删
DROP USER IF EXISTS 'game_fronttest'@'127.0.0.1';
DROP USER IF EXISTS 'game_fronttest'@'localhost';

CREATE USER 'game_fronttest'@'127.0.0.1'   IDENTIFIED BY '1234';
CREATE USER 'game_fronttest'@'localhost' IDENTIFIED BY '1234';

GRANT ALL PRIVILEGES ON game_fronttest.* TO 'game_fronttest'@'127.0.0.1';
GRANT ALL PRIVILEGES ON game_fronttest.* TO 'game_fronttest'@'localhost';
FLUSH PRIVILEGES;

-- 执行完后在 PowerShell 验证（应有一行 1）：
-- & "C:\xampp\mysql\bin\mysql.exe" -u game_fronttest -p1234 -h 127.0.0.1 -e "SELECT 1;"

-- =============================================================================
-- 仅本地应急：若 XAMPP 的 root 无密码、又懒得建子账号，可不用 game_fronttest，改用 root（勿用于生产/勿提交到公网）：
--   $env:REQUESTETH_DATABASE_SOURCE = 'root:@tcp(127.0.0.1:3306)/game_fronttest?parseTime=true'
-- 若 root 有密码，把 root: 和 @ 之间写成 root:你的root密码
-- =============================================================================
