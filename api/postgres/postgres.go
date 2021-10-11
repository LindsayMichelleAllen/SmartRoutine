package postgres

import (
	"database/sql"
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
type UpdateDeviceDatabaseRequest struct{}
type DeleteDeviceDatabaseRequest struct{}
type CreateDeviceDatabaseResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}
type UpdateDeviceDatabaseResponse struct {
	DeviceId string
	Name     string
	UserId   string
	Message  string
	Error    error
}
type DeleteDeviceDatabaseResponse struct {
	DeviceId string
	Name     string
	UserId   string
	Message  string
	Error    error
}

type DeviceDB struct {
	// intentionally left empty
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
	return &UpdateDeviceDatabaseResponse{}
}

func (d *DeviceDB) DeleteDevice(request *DeleteDeviceDatabaseRequest) *DeleteDeviceDatabaseResponse {
	return &DeleteDeviceDatabaseResponse{}
}
