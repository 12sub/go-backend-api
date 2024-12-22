CREATE TABLE IF NOT EXISTS users (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `firstname` VARCHAR(255) NOT NULL,
    'lastname' VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `createdAt` CURRENT_TIMESTAMP NOT NULL DEFAULT TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE KEY (email)

);