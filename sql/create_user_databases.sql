CREATE DATABASE smart_routine_db;
CREATE EXTENSION pgcrypto;
CREATE TABLE profile_details ( username VARCHAR(25) NOT NULL PRIMARY KEY, accountpassword text, displayname VARCHAR(50) NOT NULL);
CREATE TABLE device_details ( id uuid DEFAULT gen_random_uuid() PRIMARY KEY, devicename VARCHAR(25) NOT NULL, userid VARCHAR(25) NOT NULL REFERENCES profile_details(username));
CREATE TABLE routine_details ( id uuid DEFAULT gen_random_uuid() PRIMARY KEY, basealarm TIME with time zone NOT NULL, routinename VARCHAR(25) NOT NULL, userid VARCHAR(25) NOT NULL REFERENCES profile_details(username));
CREATE TABLE configuration_details ( id uuid DEFAULT gen_random_uuid() PRIMARY KEY, timeoffset INTEGER NOT NULL, deviceid uuid NOT NULL REFERENCES device_details(id), routineid uuid NOT NULL REFERENCES routine_details(id));


/*
profile_details{
    username  <---------------
    accountpassword |         |
    displayname     |         |
}                   |         |
                    |         |
device_details{     |         |
    userid  ---------         |
    id   <---------------     |
    devicename          |     |
}                       |     |
                        |     |
configuration_details{  |     |
    id                  |     |
    offset              |     |
    deviceid  -----------     |
    routineid ----------      |
}                      |      |
                       |      |
routine_details{       |      |
    id   <--------------      |
    userid --------------------
    name
    basealarm
}
*/