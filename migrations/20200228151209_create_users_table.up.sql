CREATE TABLE users(id int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, email VARCHAR(100) UNIQUE, password VARCHAR(100) NOT NULL, VISIBLE BOOLEAN, first_name VARCAHR(40) NOT NULL , last_name VARCAHR(40) NOT NULL)