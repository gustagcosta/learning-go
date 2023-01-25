create database `cli-mysql-crud`;

use `cli-mysql-crud`;

CREATE TABLE `cli-mysql-crud`.`posts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));
