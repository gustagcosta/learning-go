create database `api-rest-go`;

use `api-rest-go`;

CREATE TABLE `api-rest-go`.`posts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));