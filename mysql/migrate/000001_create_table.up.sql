CREATE TABLE `post` (
    `post_id` bigint PRIMARY KEY AUTO_INCREMENT,
    `post_owner_account` VARCHAR(255) NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `image_url` VARCHAR(255) NOT NULL,
    `text` VARCHAR(255) NOT NULL,
    `like_point` bigint DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `comment` (
    `comment_id` bigint PRIMARY KEY AUTO_INCREMENT,
    `post_id` bigint NOT NULL,
    `comment_owner_account` VARCHAR(255) NOT NULL,
    `text` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX `comment_index_0` ON `comment` (`comment_owner_account`);

ALTER TABLE `comment` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`post_id`) ON DELETE CASCADE;
