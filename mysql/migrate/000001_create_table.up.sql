CREATE TABLE `post` (
  `post_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `post_owner_account` VARCHAR,
  `title` VARCHAR(255),
  `image_url` VARCHAR(255),
  `text` VARCHAR(255),
  `like_point` bigint,
  `created_at` TIMESTAMP
);

CREATE TABLE `comment` (
  `comment_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `post_id` INT,
  `comment_owner_account` VARCHAR,
  `text` VARCHAR(255),
  `created_at` DATETIME
);

CREATE INDEX `comment_index_0` ON `comment` (`comment_owner_account`);

ALTER TABLE `comment` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`post_id`);
