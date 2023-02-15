CREATE TABLE IF NOT EXISTS users (
    id bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_name varchar(255) NOT NULL,
    designation varchar(255) NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);