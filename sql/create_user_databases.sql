CREATE DATABASE user;
CREATE TABLE user_profile (
    id VARCHAR(9) NOT NULL PRIMARY KEY,
    username VARCHAR(25) NOT NULL,
    displayname VARCHAR(50) NOT NULL,
);