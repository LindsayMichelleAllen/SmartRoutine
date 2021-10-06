package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "user"
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

func getDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%ssslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, errors.New("cannot open database")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("cannot connect to database")
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

	query := "INSERT INTO user_profile (id, username, displayname) VALUES ($1, $2, $3) RETURNING id, username, displayname"

	err = db.QueryRow(query, request.Id, request.Username, request.Name).Scan(&resp.Id, &resp.Username, &resp.Name)

	if err != nil {
		return &CreateUserDatabaseResponse{
			Message: "Faile to perform insert",
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

	query := "UPDATE user_profile SET username=$1 displayname=$2 WHERE id=$3 RETURNING id, username, displayname"

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

	query := "DELETE FROM user_profile WHERE id=$1 RETURNING id, username, displayname"

	err = db.QueryRow(query, request.Id).Scan(&resp.Id, &resp.Username, &resp.Name)

	if err != nil {
		return &DeleteUserDatabaseResponse{
			Message: "Failed to perform deletion",
			Error:   err,
		}
	}

	return resp
}
