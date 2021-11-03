CREATE DATABASE user_profile;
CREATE TABLE profile_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, username VARCHAR(25) NOT NULL, displayname VARCHAR(50) NOT NULL);
CREATE TABLE device_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, devicename VARCHAR(25) NOT NULL, userid VARCHAR(9) REFERENCES profile_details(id));
CREATE TABLE routine_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, routinename VARCHAR(25) NOT NULL, userid VARCHAR(9) NOT NULL REFERENCES profile_details(id));
CREATE TABLE configuration_details ( id VARCHAR(9) NOT NULL PRIMARY KEY, timeoffset INTEGER NOT NULL, deviceid VARCHAR(9) NOT NULL REFERENCES device_details(id), routineid VARCHAR(9) NOT NULL REFERENCES routine_details(id));

/* 
    If device is deleted, delete all configurations and remove from all routines.
    If routine is deleted, all configurations associated with that routine should be removed.
    A device can belong to zero to many configurations
    A routine can have zero to many configurations
    Each configuration can only have one device
    Each configuration can only be associated with one routine
*/

/*
profile_details{
    id      <-----------------
    username       |         |
    displayname    |         |
}                  |         |
                   |         |
device_details{    |         |
    userid  --------         |
    id   <---------------    |
    devicename          |    |
}                       |    |
                        |    |
configuration_details{  |    |
    id                  |    |
    offset              |    |
    deviceid  -----------    |
    routineid ----------     |
}                      |     |
                       |     |
routine_details{       |     |
    id   <--------------     |
    userid -------------------
    name
}
*/