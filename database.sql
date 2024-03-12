create database golang_database;

use golang_database;

create table users (
	id VARCHAR(100) NOT NULL,
	name VARCHAR(100) NOT NULL,
    primary key (id)
)engine = InnoDB

select * from users;
select * from comments;
desc users;
delete from users;
update users set email = null where id = 2;

ALTER TABLE users 
	ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INTEGER DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false
    ;

create table comments 
(
	id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT,
    PRIMARY KEY(id)
) ENGINE InnoDB;
    