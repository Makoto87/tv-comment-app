USE test_go_tvapp;

CREATE TABLE IF NOT EXISTS `programs` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `program_name` VARCHAR(50) NOT NULL UNIQUE,
      `created_at` DATETIME NOT NULL,
      `updated_at` DATETIME NOT NULL,
      PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `episodes` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `program_id` INT NOT NULL,
      `date` DATE NOT NULL,
      `episode_number` INT,
      `episode_title` VARCHAR(50),
      `created_at` DATETIME NOT NULL,
      `updated_at` DATETIME NOT NULL,
      PRIMARY KEY (`id`),
      FOREIGN KEY (`program_id`) REFERENCES `programs` (`id`)
);

CREATE TABLE IF NOT EXISTS `users` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `user_name` VARCHAR(30) NOT NULL,
      `created_at` DATETIME NOT NULL,
      `updated_at` DATETIME NOT NULL,
      PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `comments` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `comment` VARCHAR(140) NOT NULL,
      `episode_id` INT NOT NULL,
      `user_id` INT NOT NULL,
      `post_date` DATE NOT NULL,
      `likes` INT NOT NULL,
      `created_at` DATETIME NOT NULL,
      `updated_at` DATETIME NOT NULL,
      PRIMARY KEY (`id`),
      FOREIGN KEY (`episode_id`) REFERENCES `episodes` (`id`),
      FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

INSERT INTO users(id, user_name, created_at, updated_at) VALUES 
(1, "ゲストユーザー", 20220425, 20220425);

INSERT INTO programs(id, program_name, created_at, updated_at) VALUES 
(1, "test", 20220428, 20220428), 
(2, "e2e test", 20220429, 20220429);

INSERT INTO episodes(id, program_id, date, episode_number, episode_title, created_at, updated_at) VALUES 
(1, 1, 20220428, NULL, NULL, 20220428, 20220428),
(2, 2, 20220429, NULL, NULL, 20220429, 20220429);

INSERT INTO comments(id, comment, episode_id, user_id, post_date, likes, created_at, updated_at) VALUES 
(1, "this is test comment 1", 1, 1, NOW(), 3, NOW(), NOW()),
(2, "this is test comment 2", 1, 1, NOW(), 0, NOW(), NOW()),
(3, "this is test comment 3", 2, 1, NOW(), 10, NOW(), NOW()),
(4, "this is test comment 4", 2, 1, NOW(), 5, NOW(), NOW());