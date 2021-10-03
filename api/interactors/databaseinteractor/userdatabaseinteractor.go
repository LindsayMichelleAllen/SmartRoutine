package userdatabaseinteractor

import (
	//"database/sql"
	"errors"
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

type UserServiceInteractor interface {
	CreateUserProfile(request *CreateUserDatabaseRequest) *CreateUserDatabaseResponse
	UpdateUserProfile(request *UpdateUserDatabaseRequest) *UpdateUserDatabaseResponse
	DeleteUserProfile(request *DeleteUserDatabaseRequest) *DeleteUserDatabaseResponse
}

type UserAccountManagementServiceInteractor struct {
	// intentionally left empty
}

func (u *UserAccountManagementServiceInteractor) CreateUserProfile(request *CreateUserDatabaseRequest) *CreateUserDatabaseResponse {
	return &CreateUserDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (u *UserAccountManagementServiceInteractor) UpdateUserProfile(request *UpdateUserDatabaseRequest) *UpdateUserDatabaseResponse {
	return &UpdateUserDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (u *UserAccountManagementServiceInteractor) DeleteUserProfile(request *DeleteUserDatabaseRequest) *DeleteUserDatabaseResponse {
	return &DeleteUserDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
