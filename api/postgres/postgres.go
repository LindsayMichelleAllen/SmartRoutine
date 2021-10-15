package postgres

import (
	"api/services/model"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "user_profile"
	TBL_USER    = "profile_details"
)

type CreateUserDatabaseRequest struct {
	Username string
	Name     string
	Id       string
}

type CreateUserDatabaseResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type UpdateUserDatabaseRequest struct {
	Username string
	Name     string
	Id       string
}

type UpdateUserDatabaseResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type DeleteUserDatabaseRequest struct {
	Id string
}

type DeleteUserDatabaseResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type UserProfileDB struct {
	// intentionally left empty
}

type CreateDeviceDatabaseRequest struct {
	Id     string
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

type DeviceDB struct {
	// intentionally left empty
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

func getDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
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

func (u *UserProfileDB) CreateUserProfile(request *CreateUserDatabaseRequest) *CreateUserDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &CreateUserDatabaseResponse{
			Message: "Failed to get database",
			Error:   err,
		}
	}

	resp := &CreateUserDatabaseResponse{Message: "Successfully added user", Error: nil}

	query := "INSERT INTO profile_details (id, username, displayname) VALUES ($1, $2, $3) RETURNING id, username, displayname"

	err = db.QueryRow(query, request.Id, request.Username, request.Name).Scan(&resp.Id, &resp.Username, &resp.Name)

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

	query := "UPDATE profile_details SET username=$1, displayname=$2 WHERE id=$3 RETURNING id, username, displayname"

	err = db.QueryRow(query, request.Username, request.Name, request.Id).Scan(&resp.Id, &resp.Username, &resp.Name)

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

	query := "DELETE FROM profile_details WHERE id=$1 RETURNING id, username, displayname"

	err = db.QueryRow(query, request.Id).Scan(&resp.Id, &resp.Username, &resp.Name)

	if err != nil {
		return &DeleteUserDatabaseResponse{
			Message: "Failed to perform deletion",
			Error:   err,
		}
	}

	return resp
}

func (d *DeviceDB) CreateDevice(request *CreateDeviceDatabaseRequest) *CreateDeviceDatabaseResponse {
	db, err := getDatabase()

	if err != nil {
		return &CreateDeviceDatabaseResponse{
			Message: "Unable to connect to database",
			Error:   err,
		}
	}

	resp := &CreateDeviceDatabaseResponse{Message: "Successfully created device!", Error: nil}
	query := "INSERT INTO device_details (id, userid, devicename) VALUES ($1, $2, $3) RETURNING id, userid, devicename"
	err = db.QueryRow(query, request.Id, request.UserId, request.Name).Scan(&resp.Id, &resp.UserId, &resp.Name)

	if err != nil {
		return &CreateDeviceDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	return resp
}

func (d *DeviceDB) UpdateDevice(request *UpdateDeviceDatabaseRequest) *UpdateDeviceDatabaseResponse {
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

func (d *DeviceDB) DeleteDevice(request *DeleteDeviceDatabaseRequest) *DeleteDeviceDatabaseResponse {
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

	query := "INSERT INTO routine_details (id, routinename, userid) VALUES($1, $2, $3)"
	err = db.QueryRow(query, request.Routine.GetId(), request.Routine.GetName(), request.Routine.GetUserId()).Scan()

	if err != nil && err != sql.ErrNoRows {
		return &CreateRoutineDatabaseResponse{
			Message: "Query Failed",
			Error:   err,
		}
	}

	resp := &CreateRoutineDatabaseResponse{Routine: request.Routine, Message: "Successfully Created Routine", Error: nil}

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
