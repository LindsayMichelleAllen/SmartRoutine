package postgres

import (
	"api/services/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"os"

	_ "github.com/lib/pq"
)

type GetUserDatabaseRequest struct {
	Id string
}

type GetUserDatabaseResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type GetUsersDatabaseResponse struct {
	Users   []*model.UserProfile
	Message string
	Error   error
}

type LoginUserDatabaseRequest struct {
	Username string
	Password string
}

type LoginUserDatabaseResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type CreateUserDatabaseRequest struct {
	Username string
	Password string
	Name     string
}

type CreateUserDatabaseResponse struct {
	Username string
	Name     string
	Message  string
	Error    error
}

type UpdateUserDatabaseRequest struct {
	Username string
	Name     string
}

type UpdateUserDatabaseResponse struct {
	Username string
	Name     string
	Message  string
	Error    error
}

type DeleteUserDatabaseRequest struct {
	Id string
}

type DeleteUserDatabaseResponse struct {
	Username string
	Name     string
	Message  string
	Error    error
}

type UserProfileDB struct {
	// intentionally left empty
}

type GetDeviceDatabaseRequest struct {
	Id string
}

type GetDeviceDatabaseResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type GetDevicesDatabaseResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetUserDevicesDatabaseRequest struct {
	UserId string
}

type GetUserDevicesDatabaseResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetRoutineDevicesDatabaseRequest struct {
	RoutineId string
}

type GetRoutineDevicesDatabaseResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type CreateDeviceDatabaseRequest struct {
	UserId string
	Name   string
}

type UpdateDeviceDatabaseRequest struct {
	Id   string
	Name string
}

type DeleteDeviceDatabaseRequest struct {
	Id string
}

type CreateDeviceDatabaseResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type UpdateDeviceDatabaseResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type DeleteDeviceDatabaseResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type DeviceDBInterface interface {
	GetDevice(request *GetDeviceDatabaseRequest) *GetDeviceDatabaseResponse
	GetDevices() *GetDevicesDatabaseResponse
	GetUserDevices(request *GetUserDevicesDatabaseRequest) *GetUserDevicesDatabaseResponse
	GetRoutineDevices(request *GetRoutineDevicesDatabaseRequest) *GetRoutineDevicesDatabaseResponse
	CreateDevice(request *CreateDeviceDatabaseRequest) *CreateDeviceDatabaseResponse
	UpdateDevice(request *UpdateDeviceDatabaseRequest) *UpdateDeviceDatabaseResponse
	DeleteDevice(request *DeleteDeviceDatabaseRequest) *DeleteDeviceDatabaseResponse
}

type UnprotectedDeviceDB struct {
	// intentionally left empty
}

type GetRoutineDatabaseRequest struct {
	RoutineId string
}
type GetUserRoutinesDatabaseRequest struct {
	UserId string
}
type GetDeviceRoutinesDatabaseRequest struct {
	DeviceId string
}
type CreateRoutineDatabaseRequest struct {
	Routine *model.Routine
}
type UpdateRoutineDatabaseRequest struct {
	Routine *model.Routine
}
type DeleteRoutineDatabaseRequest struct {
	Id string
}
type GetRoutineDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}
type GetRoutinesDatabaseResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}
type GetUserRoutinesDatabaseResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}
type GetDeviceRoutinesDatabaseResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}
type CreateRoutineDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}
type UpdateRoutineDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}
type DeleteRoutineDatabaseResponse struct {
	Id      string
	Message string
	Error   error
}
type RoutineDBInterface interface {
	CreateRoutine(request *CreateRoutineDatabaseRequest) *CreateRoutineDatabaseResponse
	UpdateRoutine(request *UpdateRoutineDatabaseRequest) *UpdateRoutineDatabaseResponse
	DeleteRoutine(request *DeleteRoutineDatabaseRequest) *DeleteRoutineDatabaseResponse
}
type UnprotectedRoutineDB struct {
	/* intentionally left empty */
}

type GetConfigurationDatabaseRequest struct {
	ConfigId string
}

type GetUserConfiguraitonsDatabaseRequest struct {
	UserId string
}

type GetDeviceConfigurationsDatabaseRequest struct {
	DeviceId string
}

type GetRoutineConfigurationsDatabaseRequest struct {
	RoutineId string
}

type CreateConfigurationDatabaseRequest struct {
	Configuration *model.Configuration
}
type UpdateConfigurationDatabaseRequest struct {
	Configuration *model.Configuration
}
type DeleteConfigurationDatabaseRequest struct {
	Id string
}

type GetConfigurationDatabaseResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}

type GetConfigurationsDatabaseResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetUserConfiguraitonsDatabaseResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetDeviceConfigurationsDatabaseResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetRoutineConfigurationsDatabaseResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type CreateConfigurationDatabaseResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type UpdateConfigurationDatabaseResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type DeleteConfigurationDatabaseResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type ConfigurationDBInterface interface {
	GetConfiguration(request *GetConfigurationDatabaseRequest) *GetConfigurationDatabaseResponse
	GetConfigurations() *GetConfigurationsDatabaseResponse
	GetUserConfigurations(request *GetUserConfiguraitonsDatabaseRequest) *GetUserConfiguraitonsDatabaseResponse
	GetDeviceConfigurations(request *GetDeviceConfigurationsDatabaseRequest) *GetDeviceConfigurationsDatabaseResponse
	GetRoutineConfigurations(request *GetRoutineConfigurationsDatabaseRequest) *GetRoutineConfigurationsDatabaseResponse
	CreateConfiguration(request *CreateConfigurationDatabaseRequest) *CreateConfigurationDatabaseResponse
	UpdateConfiguration(request *UpdateConfigurationDatabaseRequest) *UpdateConfigurationDatabaseResponse
	DeleteConfiguration(request *DeleteConfigurationDatabaseRequest) *DeleteConfigurationDatabaseResponse
}
type UnprotectedConfigurationDB struct {
	/* intentionally left empty */
}

func getDatabase() (*sql.DB, error) {

	jsonFile, err := os.Open("../../config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	port, _ := strconv.Atoi(result["port"])

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		result["host"], port, result["user"], result["password"], result["dbname"])
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (u *UserProfileDB) GetUserProfile(request *GetUserDatabaseRequest) *GetUserDatabaseResponse {
	if request.Id == "" {
		return &GetUserDatabaseResponse{
			Message: "Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()
	if err != nil {
		return &GetUserDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}
	resp := &GetUserDatabaseResponse{Message: "Successfully Queried User Profile", Error: nil}
	username := ""
	displayname := ""

	query := "SELECT username, displayname FROM profile_details WHERE id=$1"

	err = db.QueryRow(query, request.Id).Scan(&username, &displayname)

	if err != nil {
		return &GetUserDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	usr := &model.UserProfile{}
	usr.SetUsername(username)
	usr.SetName(displayname)
	resp.User = usr

	return resp
}

func (u *UserProfileDB) GetUserProfiles() *GetUsersDatabaseResponse {
	db, err := getDatabase()
	if err != nil {
		return &GetUsersDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &GetUsersDatabaseResponse{Message: "Successfully Queried All User Profiles", Error: nil}

	query := "SELECT username, displayname FROM profile_details"
	rows, err := db.Query(query)

	if err != nil {
		return &GetUsersDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	usrs := make([]*model.UserProfile, 0)
	for rows.Next() {
		var username string
		var displayname string
		err = rows.Scan(&username, &displayname)
		if err != nil {
			// handle this error
			panic(err)
		}
		usr := &model.UserProfile{}
		usr.SetUsername(username)
		usr.SetName(displayname)
		usrs = append(usrs, usr)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return &GetUsersDatabaseResponse{
			Message: err.Error(),
			Error:   err,
		}
	}

	resp.Users = usrs
	return resp
}

func (u *UserProfileDB) UserProfileLogin(request *LoginUserDatabaseRequest) *LoginUserDatabaseResponse {
	if request.Username == "" || request.Password == "" {
		return &LoginUserDatabaseResponse{
			User:    nil,
			Message: "Input Field(s) Missing",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()
	if err != nil {
		return &LoginUserDatabaseResponse{
			User:    nil,
			Message: "Unable to connect to database",
			Error:   err,
		}
	}
	var username, displayname string
	query := "SELECT username, displayname FROM profile_details WHERE username=$1 AND accountpassword=crypt($2, accountpassword)"
	err = db.QueryRow(query, request.Username, request.Password).Scan(&username, &displayname)
	if err != nil {
		return &LoginUserDatabaseResponse{
			User:    nil,
			Message: "Incorrect Credentials",
			Error:   err,
		}
	}
	user := &model.UserProfile{}
	user.SetUsername(username)
	user.SetName(displayname)
	user.SetAuthorizationStatus(true)
	return &LoginUserDatabaseResponse{
		User:    user,
		Message: "Login Successful",
		Error:   nil,
	}
}

func (u *UserProfileDB) CreateUserProfile(request *CreateUserDatabaseRequest) *CreateUserDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &CreateUserDatabaseResponse{
			Message: "Failed to get database",
			Error:   err,
		}
	}

	resp := &CreateUserDatabaseResponse{Message: "Successfully added user", Error: nil}

	query := `INSERT INTO profile_details (username, accountpassword, displayname) 
			  VALUES ($1, crypt($2, gen_salt('bf',8)), $3) 
			  RETURNING username, displayname`

	err = db.QueryRow(query, request.Username, request.Password, request.Name).Scan(&resp.Username, &resp.Name)

	if err != nil {
		return &CreateUserDatabaseResponse{
			Message: "Failed to perform insert",
			Error:   err,
		}
	}
	return resp
}

func (u *UserProfileDB) UpdateUserProfile(request *UpdateUserDatabaseRequest) *UpdateUserDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &UpdateUserDatabaseResponse{
			Message: "Failed to get database",
			Error:   err,
		}
	}

	resp := &UpdateUserDatabaseResponse{Message: "Successfully updated user profile", Error: nil}

	query := "UPDATE profile_details SET displayname=$1 WHERE username=$2 RETURNING username, displayname"

	err = db.QueryRow(query, request.Name, request.Username).Scan(&resp.Username, &resp.Name)

	if err != nil {
		return &UpdateUserDatabaseResponse{
			Message: "Failed to perform update",
			Error:   err,
		}
	}

	return resp
}

func (u *UserProfileDB) DeleteUserProfile(request *DeleteUserDatabaseRequest) *DeleteUserDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &DeleteUserDatabaseResponse{
			Message: "Failed to get database",
			Error:   err,
		}
	}

	resp := &DeleteUserDatabaseResponse{Message: "Successfully deleted user profile", Error: nil}

	query := "DELETE FROM profile_details WHERE username=$1 RETURNING username, displayname"

	err = db.QueryRow(query, request.Id).Scan(&resp.Username, &resp.Name)

	if err != nil {
		return &DeleteUserDatabaseResponse{
			Message: "Failed to perform deletion",
			Error:   err,
		}
	}

	return resp
}

func (u *UnprotectedDeviceDB) GetDevice(request *GetDeviceDatabaseRequest) *GetDeviceDatabaseResponse {
	if request.Id == "" {
		return &GetDeviceDatabaseResponse{
			Message: "Device ID not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &GetDeviceDatabaseResponse{
			Message: "Unable To Connect To Database",
			Error:   err,
		}
	}

	resp := &GetDeviceDatabaseResponse{Message: "Successfully Queried Device", Error: nil}
	id := ""
	userid := ""
	devicename := ""

	query := "SELECT * FROM device_details WHERE id=$1"
	err = db.QueryRow(query, request.Id).Scan(&id, &devicename, &userid)

	if err != nil {
		return &GetDeviceDatabaseResponse{
			Message: "Device Query Failed",
			Error:   err,
		}
	}

	dev := &model.Device{}
	dev.SetId(id)
	dev.SetUserId(userid)
	dev.SetName(devicename)

	resp.Device = dev

	return resp
}

func (u *UnprotectedDeviceDB) GetDevices() *GetDevicesDatabaseResponse {
	db, err := getDatabase()
	if err != nil {
		return &GetDevicesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &GetDevicesDatabaseResponse{Message: "Successfully Queried All Devices", Error: nil}

	query := "SELECT * FROM device_details"
	rows, err := db.Query(query)

	if err != nil {
		return &GetDevicesDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	devs := make([]*model.Device, 0)
	for rows.Next() {
		var id string
		var name string
		var userid string
		err = rows.Scan(&id, &name, &userid)
		if err != nil {
			// handle this error
			panic(err)
		}
		dev := &model.Device{}
		dev.SetId(id)
		dev.SetName(name)
		dev.SetUserId(userid)
		devs = append(devs, dev)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return &GetDevicesDatabaseResponse{
			Message: err.Error(),
			Error:   err,
		}
	}

	resp.Devices = devs
	return resp
}

func (u *UnprotectedDeviceDB) GetUserDevices(request *GetUserDevicesDatabaseRequest) *GetUserDevicesDatabaseResponse {
	if request.UserId == "" {
		return &GetUserDevicesDatabaseResponse{
			Message: "UserId not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &GetUserDevicesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &GetUserDevicesDatabaseResponse{Message: "Successfully Queried User Devices", Error: nil}
	query := "SELECT * FROM device_details WHERE userid=$1"

	rows, err := db.Query(query, request.UserId)

	if err != nil {
		return &GetUserDevicesDatabaseResponse{
			Message: "User Devices Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	devs := make([]*model.Device, 0)
	for rows.Next() {
		var id string
		var name string
		var userid string
		err := rows.Scan(&id, &name, &userid)
		if err != nil {
			return &GetUserDevicesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		dev := &model.Device{}
		dev.SetId(id)
		dev.SetName(name)
		dev.SetUserId(userid)
		devs = append(devs, dev)
	}

	resp.Devices = devs
	return resp
}

func (u *UnprotectedDeviceDB) GetRoutineDevices(request *GetRoutineDevicesDatabaseRequest) *GetRoutineDevicesDatabaseResponse {
	if request.RoutineId == "" {
		return &GetRoutineDevicesDatabaseResponse{
			Message: "RoutineId not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &GetRoutineDevicesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &GetRoutineDevicesDatabaseResponse{Message: "Successfully Queried Routine Devices", Error: nil}
	query := "SELECT deviceid FROM configuration_details WHERE routineid=$1"

	rows, err := db.Query(query, request.RoutineId)

	if err != nil {
		return &GetRoutineDevicesDatabaseResponse{
			Message: "Query Failed (deviceid from configuration)",
			Error:   err,
		}
	}

	defer rows.Close()
	devids := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return &GetRoutineDevicesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		devids = append(devids, id)
	}

	query = fmt.Sprintf("SELECT * FROM device_details WHERE id IN ('%s')", strings.Join(devids, "', '"))
	rows, err = db.Query(query)

	if err != nil {
		return &GetRoutineDevicesDatabaseResponse{
			Message: "Query Failed (deviceid from configuration)",
			Error:   err,
		}
	}

	defer rows.Close()
	devs := make([]*model.Device, 0)
	for rows.Next() {
		var id string
		var name string
		var userid string
		err := rows.Scan(&id, &name, &userid)
		if err != nil {
			return &GetRoutineDevicesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}

		dev := &model.Device{}
		dev.SetId(id)
		dev.SetName(name)
		dev.SetUserId(userid)

		devs = append(devs, dev)
	}

	resp.Devices = devs
	return resp
}

func (d *UnprotectedDeviceDB) CreateDevice(request *CreateDeviceDatabaseRequest) *CreateDeviceDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &CreateDeviceDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &CreateDeviceDatabaseResponse{Message: "Successfully created device!", Error: nil}
	query := "INSERT INTO device_details (id, userid, devicename) VALUES (gen_random_uuid(), $1, $2) RETURNING id, userid, devicename"
	err = db.QueryRow(query, request.UserId, request.Name).Scan(&resp.Id, &resp.UserId, &resp.Name)

	if err != nil {
		return &CreateDeviceDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	return resp
}

func (d *UnprotectedDeviceDB) UpdateDevice(request *UpdateDeviceDatabaseRequest) *UpdateDeviceDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &UpdateDeviceDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &UpdateDeviceDatabaseResponse{Message: "Successfully updated device!", Error: nil}
	query := "UPDATE device_details SET devicename=$1 WHERE id=$2 RETURNING id, userid, devicename"
	err = db.QueryRow(query, request.Name, request.Id).Scan(&resp.Id, &resp.UserId, &resp.Name)

	if err != nil {
		return &UpdateDeviceDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	return resp
}

func (d *UnprotectedDeviceDB) DeleteDevice(request *DeleteDeviceDatabaseRequest) *DeleteDeviceDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &DeleteDeviceDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &DeleteDeviceDatabaseResponse{Message: "Successfully removed device!", Error: nil}
	query := "DELETE FROM device_details WHERE id=$1 RETURNING id, userid, devicename"
	err = db.QueryRow(query, request.Id).Scan(&resp.Id, &resp.UserId, &resp.Name)

	if err != nil {
		return &DeleteDeviceDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	return resp
}

func (r *UnprotectedRoutineDB) GetRoutine(request *GetRoutineDatabaseRequest) *GetRoutineDatabaseResponse {
	if request.RoutineId == "" {
		return &GetRoutineDatabaseResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db, err := getDatabase()

	if err != nil {
		return &GetRoutineDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	var name, basealarm, userid string

	query := `
SELECT r.routinename, r.basealarm, r.userid, c.id, c.timeoffset, d.id, d.devicename
FROM routine_details r
LEFT JOIN configuration_details c ON r.id = c.routineid
LEFT JOIN device_details d ON d.id = c.deviceid
WHERE r.id = $1`

	rows, err := db.Query(query, request.RoutineId)

	if err != nil {
		return &GetRoutineDatabaseResponse{
			Message: "Routine Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	configs := make([]*model.Configuration, 0)
	for rows.Next() {
		var routineName, baseAlarm, userId, configId, deviceId, deviceName sql.NullString
		var timeoffset sql.NullInt32
		err = rows.Scan(&routineName, &baseAlarm, &userId, &configId, &timeoffset, &deviceId, &deviceName)
		if err != nil {
			return &GetRoutineDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		name = routineName.String
		basealarm = baseAlarm.String
		userid = userId.String
		config := &model.Configuration{}
		dev := &model.Device{}
		dev.SetId(deviceId.String)
		dev.SetName(deviceName.String)
		dev.SetUserId(userId.String)
		config.SetId(configId.String)
		config.SetRoutineId(request.RoutineId)
		config.SetOffset(int(timeoffset.Int32))
		config.SetDevice(dev)
		configs = append(configs, config)
	}

	routine := &model.Routine{}
	routine.PopulateRoutine(request.RoutineId, name, userid, basealarm, configs)

	return &GetRoutineDatabaseResponse{Routine: routine, Message: "Successfully Queried Routine", Error: nil}
}

func (r *UnprotectedRoutineDB) GetRoutines() *GetRoutinesDatabaseResponse {

	db, err := getDatabase()

	if err != nil {
		return &GetRoutinesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := `SELECT r.id, r.routinename, r.basealarm, r.userid, c.id, c.timeoffset, d.id, d.devicename
			  FROM routine_details r, configuration_details c, device_details d
			  WHERE r.id = c.routineid AND c.deviceid = d.id`

	rows, err := db.Query(query)

	if err != nil {
		return &GetRoutinesDatabaseResponse{
			Message: "Routine Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	res := make(map[string][][7]string)
	for rows.Next() {
		var routineId, routineName, basealarm, userId, configId, deviceId, deviceName string
		var timeoffset int
		err = rows.Scan(&routineId, &routineName, &basealarm, &userId, &configId, &timeoffset, &deviceId, &deviceName)
		if err != nil {
			return &GetRoutinesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		if value, ok := res[routineId]; ok {
			temp := [...]string{routineName, basealarm, userId, configId, strconv.Itoa(timeoffset), deviceId, deviceName}
			res[routineId] = append(value, temp)
		} else {
			temp := [...]string{routineName, basealarm, userId, configId, strconv.Itoa(timeoffset), deviceId, deviceName}
			res[routineId] = append(res[routineId], temp)
		}
	}

	routines := make([]*model.Routine, 0)
	for key, val := range res {
		tempRoutine := &model.Routine{}
		tempConfigs := make([]*model.Configuration, 0)
		routineName := ""
		userId := ""
		basealarm := ""
		for i, item := range val {
			if i == 0 {
				routineName = item[0]
				basealarm = item[1]
				userId = item[2]
			}
			tempConfig := &model.Configuration{}
			tempDev := &model.Device{}
			tempDev.SetUserId(item[2])
			tempConfig.SetId(item[3])
			var offset int
			offset, err = strconv.Atoi(item[4])
			if err != nil {
				return &GetRoutinesDatabaseResponse{
					Routines: nil,
					Message:  "Unable to convert offset to integer type",
					Error:    err,
				}
			}
			tempConfig.SetOffset(offset)
			tempConfig.SetRoutineId(key)
			tempDev.SetId(item[5])
			tempDev.SetName(item[6])
			tempConfig.SetDevice(tempDev)
			tempConfigs = append(tempConfigs, tempConfig)
		}
		tempRoutine.PopulateRoutine(key, routineName, userId, basealarm, tempConfigs)
		routines = append(routines, tempRoutine)
	}

	return &GetRoutinesDatabaseResponse{Routines: routines, Message: "Successfully Queried Routines", Error: nil}
}

func (R *UnprotectedRoutineDB) GetUserRoutines(request *GetUserRoutinesDatabaseRequest) *GetUserRoutinesDatabaseResponse {
	if request.UserId == "" {
		return &GetUserRoutinesDatabaseResponse{
			Routines: nil,
			Message:  "User Id not provided",
			Error:    errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &GetUserRoutinesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := `
SELECT r.id, r.routinename, r.basealarm, r.userid, c.id, c.timeoffset, d.id, d.devicename
FROM routine_details r
LEFT JOIN configuration_details c ON r.id = c.routineid
LEFT JOIN device_details d ON d.id = c.deviceid
WHERE r.userid = $1
`

	rows, err := db.Query(query, request.UserId)

	if err != nil {
		return &GetUserRoutinesDatabaseResponse{
			Message: "Routine Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	res := make(map[string][][7]string)
	for rows.Next() {
		var routineId, routineName, basealarm, userId, configId, deviceId, deviceName sql.NullString
		var timeoffset sql.NullInt32
		err = rows.Scan(&routineId, &routineName, &basealarm, &userId, &configId, &timeoffset, &deviceId, &deviceName)
		if err != nil {
			return &GetUserRoutinesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		if value, ok := res[routineId.String]; ok {
			temp := [...]string{
				routineName.String,
				basealarm.String,
				userId.String,
				configId.String,
				strconv.Itoa(int(timeoffset.Int32)),
				deviceId.String,
				deviceName.String,
			}
			res[routineId.String] = append(value, temp)
		} else {
			temp := [...]string{
				routineName.String,
				basealarm.String,
				userId.String,
				configId.String,
				strconv.Itoa(int(timeoffset.Int32)),
				deviceId.String,
				deviceName.String,
			}
			res[routineId.String] = append(res[routineId.String], temp)
		}
	}

	routines := make([]*model.Routine, 0)
	for key, val := range res {
		tempRoutine := &model.Routine{}
		tempConfigs := make([]*model.Configuration, 0)
		routineName := ""
		userId := ""
		basealarm := ""
		for i, item := range val {
			if i == 0 {
				routineName = item[0]
				basealarm = item[1]
				userId = item[2]
			}
			tempConfig := &model.Configuration{}
			tempDev := &model.Device{}
			tempDev.SetUserId(item[2])
			tempConfig.SetId(item[3])
			var offset int
			offset, err = strconv.Atoi(item[4])
			if err != nil {
				return &GetUserRoutinesDatabaseResponse{
					Routines: nil,
					Message:  "Unable to convert offset to integer type",
					Error:    err,
				}
			}
			tempConfig.SetOffset(offset)
			tempConfig.SetRoutineId(key)
			tempDev.SetId(item[5])
			tempDev.SetName(item[6])
			tempConfig.SetDevice(tempDev)
			tempConfigs = append(tempConfigs, tempConfig)
		}
		tempRoutine.PopulateRoutine(key, routineName, userId, basealarm, tempConfigs)
		routines = append(routines, tempRoutine)
	}

	return &GetUserRoutinesDatabaseResponse{Routines: routines, Message: "Successfully Queried User Routines", Error: nil}
}

func (r *UnprotectedRoutineDB) GetDeviceRoutines(request *GetDeviceRoutinesDatabaseRequest) *GetDeviceRoutinesDatabaseResponse {
	if request.DeviceId == "" {
		return &GetDeviceRoutinesDatabaseResponse{
			Routines: nil,
			Message:  "Device Id not provided",
			Error:    errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &GetDeviceRoutinesDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := `SELECT r.id, r.routinename, r.basealarm, r.userid, c.id, c.timeoffset, d.id, d.devicename
			  FROM routine_details r, configuration_details c, device_details d
			  WHERE d.id = $1 AND r.id = c.routineid AND c.deviceid = d.id`

	rows, err := db.Query(query, request.DeviceId)

	if err != nil {
		return &GetDeviceRoutinesDatabaseResponse{
			Message: "Routine Query Failed",
			Error:   err,
		}
	}

	defer rows.Close()
	res := make(map[string][][7]string)
	for rows.Next() {
		var routineId, routineName, basealarm, userId, configId, deviceId, deviceName string
		var timeoffset int
		err = rows.Scan(&routineId, &routineName, &basealarm, &userId, &configId, &timeoffset, &deviceId, &deviceName)
		if err != nil {
			return &GetDeviceRoutinesDatabaseResponse{
				Message: err.Error(),
				Error:   err,
			}
		}
		if value, ok := res[routineId]; ok {
			temp := [...]string{routineName, basealarm, userId, configId, strconv.Itoa(timeoffset), deviceId, deviceName}
			res[routineId] = append(value, temp)
		} else {
			temp := [...]string{routineName, basealarm, userId, configId, strconv.Itoa(timeoffset), deviceId, deviceName}
			res[routineId] = append(res[routineId], temp)
		}
	}

	routines := make([]*model.Routine, 0)
	for key, val := range res {
		tempRoutine := &model.Routine{}
		tempConfigs := make([]*model.Configuration, 0)
		routineName := ""
		basealarm := ""
		userId := ""
		for i, item := range val {
			if i == 0 {
				routineName = item[0]
				basealarm = item[1]
				userId = item[2]
			}
			tempConfig := &model.Configuration{}
			tempDev := &model.Device{}
			tempDev.SetUserId(item[2])
			tempConfig.SetId(item[3])
			var offset int
			offset, err = strconv.Atoi(item[4])
			if err != nil {
				return &GetDeviceRoutinesDatabaseResponse{
					Routines: nil,
					Message:  "Unable to convert offset to integer type",
					Error:    err,
				}
			}
			tempConfig.SetOffset(offset)
			tempConfig.SetRoutineId(key)
			tempDev.SetId(item[5])
			tempDev.SetName(item[6])
			tempConfig.SetDevice(tempDev)
			tempConfigs = append(tempConfigs, tempConfig)
		}
		tempRoutine.PopulateRoutine(key, routineName, userId, basealarm, tempConfigs)
		routines = append(routines, tempRoutine)
	}

	return &GetDeviceRoutinesDatabaseResponse{Routines: routines, Message: "Successfully Queried Device Routines", Error: nil}
}

func (r *UnprotectedRoutineDB) CreateRoutine(request *CreateRoutineDatabaseRequest) *CreateRoutineDatabaseResponse {
	if request.Routine == nil {
		return &CreateRoutineDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &CreateRoutineDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &CreateRoutineDatabaseResponse{Routine: request.Routine, Message: "Successfully Created Routine", Error: nil}
	var id string
	query := "INSERT INTO routine_details (id, basealarm, routinename, userid) VALUES(gen_random_uuid(), $1, $2, $3) RETURNING id"
	err = db.QueryRow(query, request.Routine.GetBaseAlarm(), request.Routine.GetName(), request.Routine.GetUserId()).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return &CreateRoutineDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	request.Routine.SetId(id)

	return resp
}
func (r *UnprotectedRoutineDB) UpdateRoutine(request *UpdateRoutineDatabaseRequest) *UpdateRoutineDatabaseResponse {
	if request.Routine == nil {
		return &UpdateRoutineDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &UpdateRoutineDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := "UPDATE routine_details SET routinename=$1 WHERE id=$2"
	err = db.QueryRow(query, request.Routine.GetName(), request.Routine.GetId()).Scan()

	if err != nil && err != sql.ErrNoRows {
		return &UpdateRoutineDatabaseResponse{
			Message: "Query failed",
			Error:   err,
		}
	}

	resp := &UpdateRoutineDatabaseResponse{Routine: request.Routine, Message: "Successfully Updated Routine", Error: nil}

	return resp
}
func (r *UnprotectedRoutineDB) DeleteRoutine(request *DeleteRoutineDatabaseRequest) *DeleteRoutineDatabaseResponse {
	if request.Id == "" {
		return &DeleteRoutineDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &DeleteRoutineDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := "DELETE FROM routine_details WHERE id=$1"
	err = db.QueryRow(query, request.Id).Scan()

	if err != nil && err != sql.ErrNoRows {
		return &DeleteRoutineDatabaseResponse{
			Message: "Query failed",
			Error:   err,
		}
	}

	return &DeleteRoutineDatabaseResponse{
		Id:      request.Id,
		Message: "Successfully removed routine!",
		Error:   nil,
	}
}

func (c *UnprotectedConfigurationDB) GetConfiguration(request *GetConfigurationDatabaseRequest) *GetConfigurationDatabaseResponse {
	if request.ConfigId == "" {
		return &GetConfigurationDatabaseResponse{
			Configuration: nil,
			Message:       "Config Id not provided",
			Error:         errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()
	if err != nil {
		return &GetConfigurationDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}
	var id, devId, routineId, devName, userId string
	var timeoffset int
	query := `SELECT c.id, c.timeoffset, c.routineid, d.id, d.devicename, d.userid
			  FROM configuration_details c, device_details d 
			  WHERE c.id=$1`
	err = db.QueryRow(query, request.ConfigId).Scan(&id, &timeoffset, &routineId, &devId, &devName, &userId)
	if err != nil {
		return &GetConfigurationDatabaseResponse{
			Configuration: nil,
			Message:       err.Error(),
		}
	}
	config := &model.Configuration{}
	dev := &model.Device{}
	dev.SetId(devId)
	dev.SetName(devName)
	dev.SetUserId(userId)
	config.SetDevice(dev)
	config.SetId(id)
	config.SetRoutineId(routineId)
	config.SetOffset(timeoffset)
	return &GetConfigurationDatabaseResponse{Configuration: config, Message: "Successfully Queried Configuration", Error: nil}
}

func (c *UnprotectedConfigurationDB) GetConfigurations() *GetConfigurationsDatabaseResponse {
	db, err := getDatabase()
	if err != nil {
		return &GetConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        "Unable to connect to database",
			Error:          err,
		}
	}
	// TODO: Query returning duplicates
	query := `SELECT c.id, c.timeoffset, c.routineid, d.id, d.devicename, d.userid
			  FROM configuration_details c, device_details d`
	rows, err := db.Query(query)
	if err != nil {
		return &GetConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        err.Error(),
			Error:          err,
		}
	}

	var id, devId, routineId, devName, userId string
	var timeoffset int
	configs := make([]*model.Configuration, 0)
	for rows.Next() {
		err := rows.Scan(&id, &timeoffset, &routineId, &devId, &devName, &userId)
		if err != nil {
			return &GetConfigurationsDatabaseResponse{
				Configurations: nil,
				Message:        "Issue while scanning row (GetConfgiurations)",
				Error:          err,
			}
		}
		config := &model.Configuration{}
		dev := &model.Device{}
		dev.SetId(devId)
		dev.SetName(devName)
		dev.SetUserId(userId)
		config.SetDevice(dev)
		config.SetId(id)
		config.SetRoutineId(routineId)
		config.SetOffset(timeoffset)
		configs = append(configs, config)
	}
	return &GetConfigurationsDatabaseResponse{Configurations: configs, Message: "Successfully Queried Configurations", Error: nil}
}

func (c *UnprotectedConfigurationDB) GetUserConfigurations(request *GetUserConfiguraitonsDatabaseRequest) *GetUserConfiguraitonsDatabaseResponse {
	if request.UserId == "" {
		return &GetUserConfiguraitonsDatabaseResponse{
			Configurations: nil,
			Message:        "User Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db, err := getDatabase()
	if err != nil {
		return &GetUserConfiguraitonsDatabaseResponse{
			Configurations: nil,
			Message:        "Unable to connect to database",
			Error:          err,
		}
	}
	query := `SELECT c.id, c.timeoffset, c.routineid, d.id, d.devicename, d.userid
			  FROM configuration_details c, device_details d
			  WHERE d.userid = $1`
	rows, err := db.Query(query, request.UserId)
	if err != nil {
		return &GetUserConfiguraitonsDatabaseResponse{
			Configurations: nil,
			Message:        err.Error(),
			Error:          err,
		}
	}

	var id, devId, routineId, devName, userId string
	var timeoffset int
	configs := make([]*model.Configuration, 0)
	for rows.Next() {
		err := rows.Scan(&id, &timeoffset, &routineId, &devId, &devName, &userId)
		if err != nil {
			return &GetUserConfiguraitonsDatabaseResponse{
				Configurations: nil,
				Message:        "Issue while scanning row (GetUserConfgiurations)",
				Error:          err,
			}
		}
		config := &model.Configuration{}
		dev := &model.Device{}
		dev.SetId(devId)
		dev.SetName(devName)
		dev.SetUserId(userId)
		config.SetDevice(dev)
		config.SetId(id)
		config.SetRoutineId(routineId)
		config.SetOffset(timeoffset)
		configs = append(configs, config)
	}
	return &GetUserConfiguraitonsDatabaseResponse{Configurations: configs, Message: "Successfully Queried User Configurations", Error: nil}
}

func (c *UnprotectedConfigurationDB) GetDeviceConfigurations(request *GetDeviceConfigurationsDatabaseRequest) *GetDeviceConfigurationsDatabaseResponse {
	if request.DeviceId == "" {
		return &GetDeviceConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        "Device Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db, err := getDatabase()
	if err != nil {
		return &GetDeviceConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        "Unable to connect to database",
			Error:          err,
		}
	}
	query := `SELECT c.id, c.timeoffset, c.routineid, d.id, d.devicename, d.userid
			  FROM configuration_details c, device_details d
			  WHERE d.id = $1`
	rows, err := db.Query(query, request.DeviceId)
	if err != nil {
		return &GetDeviceConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        err.Error(),
			Error:          err,
		}
	}

	var id, devId, routineId, devName, userId string
	var timeoffset int
	configs := make([]*model.Configuration, 0)
	for rows.Next() {
		err := rows.Scan(&id, &timeoffset, &routineId, &devId, &devName, &userId)
		if err != nil {
			return &GetDeviceConfigurationsDatabaseResponse{
				Configurations: nil,
				Message:        "Issue while scanning row (GetDeviceConfgiurations)",
				Error:          err,
			}
		}
		config := &model.Configuration{}
		dev := &model.Device{}
		dev.SetId(devId)
		dev.SetName(devName)
		dev.SetUserId(userId)
		config.SetDevice(dev)
		config.SetId(id)
		config.SetRoutineId(routineId)
		config.SetOffset(timeoffset)
		configs = append(configs, config)
	}
	return &GetDeviceConfigurationsDatabaseResponse{Configurations: configs, Message: "Successfully Queried Device Configurations", Error: nil}
}

func (c *UnprotectedConfigurationDB) GetRoutineConfigurations(request *GetRoutineConfigurationsDatabaseRequest) *GetRoutineConfigurationsDatabaseResponse {
	if request.RoutineId == "" {
		return &GetRoutineConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        "Routine Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db, err := getDatabase()
	if err != nil {
		return &GetRoutineConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        "Unable to connect to database",
			Error:          err,
		}
	}
	query := `SELECT c.id, c.timeoffset, c.routineid, d.id, d.devicename, d.userid
			  FROM configuration_details c, device_details d
			  WHERE c.routineid = $1`
	rows, err := db.Query(query, request.RoutineId)
	if err != nil {
		return &GetRoutineConfigurationsDatabaseResponse{
			Configurations: nil,
			Message:        err.Error(),
			Error:          err,
		}
	}

	var id, devId, routineId, devName, userId string
	var timeoffset int
	configs := make([]*model.Configuration, 0)
	for rows.Next() {
		err := rows.Scan(&id, &timeoffset, &routineId, &devId, &devName, &userId)
		if err != nil {
			return &GetRoutineConfigurationsDatabaseResponse{
				Configurations: nil,
				Message:        "Issue while scanning row (GetRoutineConfgiurations)",
				Error:          err,
			}
		}
		config := &model.Configuration{}
		dev := &model.Device{}
		dev.SetId(devId)
		dev.SetName(devName)
		dev.SetUserId(userId)
		config.SetDevice(dev)
		config.SetId(id)
		config.SetRoutineId(routineId)
		config.SetOffset(timeoffset)
		configs = append(configs, config)
	}
	return &GetRoutineConfigurationsDatabaseResponse{Configurations: configs, Message: "Successfully Queried Routine Configurations", Error: nil}
}

func (c *UnprotectedConfigurationDB) CreateConfiguration(request *CreateConfigurationDatabaseRequest) *CreateConfigurationDatabaseResponse {
	if request.Configuration == nil {
		return &CreateConfigurationDatabaseResponse{
			Message: "Configuration not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &CreateConfigurationDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	var id string
	query := "INSERT INTO configuration_details (id, timeoffset, deviceid, routineid) VALUES(gen_random_uuid(), $1, $2, $3) RETURNING id"
	err = db.QueryRow(query,
		request.Configuration.GetOffset(),
		request.Configuration.GetDevice().GetId(),
		request.Configuration.GetRoutineId()).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return &CreateConfigurationDatabaseResponse{
			Message: err.Error(),
			Error:   err,
		}
	}

	request.Configuration.SetId(id)
	resp := &CreateConfigurationDatabaseResponse{Configuration: request.Configuration, Message: "Sucessfully Created Configuration", Error: nil}

	return resp
}

func (c *UnprotectedConfigurationDB) UpdateConfiguration(request *UpdateConfigurationDatabaseRequest) *UpdateConfigurationDatabaseResponse {
	if request.Configuration == nil {
		return &UpdateConfigurationDatabaseResponse{
			Message: "Configuration not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &UpdateConfigurationDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	query := "UPDATE configuration_details SET timeoffset=$1 WHERE id=$2"
	err = db.QueryRow(query,
		request.Configuration.GetOffset(),
		request.Configuration.GetId()).Scan()

	if err != nil && err != sql.ErrNoRows {
		return &UpdateConfigurationDatabaseResponse{
			Message: err.Error(),
			Error:   err,
		}
	}

	resp := &UpdateConfigurationDatabaseResponse{Configuration: request.Configuration, Message: "Sucessfully Updated Configuration", Error: nil}

	return resp
}

func (c *UnprotectedConfigurationDB) DeleteConfiguration(request *DeleteConfigurationDatabaseRequest) *DeleteConfigurationDatabaseResponse {
	if request.Id == "" {
		return &DeleteConfigurationDatabaseResponse{
			Message: "Configuration id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db, err := getDatabase()

	if err != nil {
		return &DeleteConfigurationDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &DeleteConfigurationDatabaseResponse{Message: "Sucessfully Removed Configuration", Error: nil}
	configId := ""
	timeoffset := ""
	deviceid := ""
	routineid := ""
	query := "DELETE FROM configuration_details WHERE id=$1 RETURNING id, timeoffset, deviceid, routineid"
	err = db.QueryRow(query, request.Id).Scan(&configId, &timeoffset, &deviceid, &routineid)

	if err != nil {
		return &DeleteConfigurationDatabaseResponse{
			Message: err.Error(),
			Error:   err,
		}
	}

	config := &model.Configuration{}
	dev := &model.Device{}
	dev.SetId(deviceid)
	offset, _ := strconv.Atoi(timeoffset)
	config.SetId(configId)
	config.SetOffset(offset)
	config.SetDevice(dev)
	config.SetRoutineId(routineid)

	resp.Configuration = config

	return resp
}
