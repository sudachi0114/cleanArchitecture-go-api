---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
CREATE TABLE IF NOT EXISTS `users` (
  `id`          INT AUTO_INCREMENT PRIMARY KEY,
  `first_name`  VARCHAR(20) NOT NULL,
  `last_name`   VARCHAR(20) NOT NULL
);
