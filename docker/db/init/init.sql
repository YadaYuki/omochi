CREATE DATABASE IF NOT EXISTS omochi_db;
USE omochi_db;
CREATE TABLE IF NOT EXISTS words (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  content VARBINARY(1024) NOT NULL,
  created_at datetime default current_timestamp
) ENGINE = InnoDB DEFAULT CHARSET=utf8;