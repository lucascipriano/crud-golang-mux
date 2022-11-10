CREATE DATABASE devbook;
CREATE TABLE usuarios(id int auto_increment primary key, name varchar(100) not null,
  email varchar(100) not null) ENGINE=INNODB;