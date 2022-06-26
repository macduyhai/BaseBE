CREATE TABLE staffs (  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    username varchar(255) UNIQUE,
    password varchar(255)  NOT NULL,
    fullname varchar(255)  NOT NULL,
    salary bigint(20) DEFAULT NULL,
    create_time DATETIME COMMENT 'Create Time',
    update_time DATETIME COMMENT 'Update Time'

) DEFAULT CHARSET UTF8 COMMENT '';