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