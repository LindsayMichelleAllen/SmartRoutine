CREATE DATABASE user_profile;
CREATE TABLE profile_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, username VARCHAR(25) NOT NULL, displayname VARCHAR(50) NOT NULL);
CREATE TABLE device_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, devicename VARCHAR(25) NOT NULL, userid VARCHAR(9) REFERENCES profile_details(id));
CREATE TABLE configuration_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, offset INTEGER, deviceid VARCHAR(9) REFERENCES device_details(id), userid VARCHAR(9) REFERENCES profile_details(id));
CREATE TABLE routine_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, configurationid REFERENCES configuration_details(id));