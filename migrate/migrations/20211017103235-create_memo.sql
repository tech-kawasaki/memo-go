
-- +migrate Up
CREATE TABLE IF NOT EXISTS memo (
    id int(15) AUTO_INCREMENT,
    title varchar(255),
    content text,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id));

-- +migrate Down
DROP TABLE IF EXISTS memo;