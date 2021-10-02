## Database - Postgres
Data to be stored:
* user profile
    * username
    * name
    * userId
* device
    * deviceId
* routine
    * routineId
    * base alarm
    * devices
        * offset

### Postgres vs MySQL
Data being handled will not require complex queries, so MySQL should be adequate for the current stage. However, using Postgres 
will allow for better scalability and data-driven operations in the future if the application is expanded to support more 
smarthome features.

