
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int, name string, email string, password string);
-- +migrate Down
DROP TABLE IF EXISTS users;