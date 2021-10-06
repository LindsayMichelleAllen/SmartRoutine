package userdatabaseinteractor

import (
	"api/postgres"
)

type CreateUserInteractorRequest struct {
	Username string
	Name     string
	Id       string
}

type CreateUserInteractorResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type UpdateUserInteractorRequest struct {
	Username string
	Name     string
	Id       string
}

type UpdateUserInteractorResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type DeleteUserInteractorRequest struct {
	Id string
}

type DeleteUserInteractorResponse struct {
	Username string
	Name     string
	Id       string
	Message  string
	Error    error
}

type UserServiceInteractor interface {
	CreateUserProfile(request *CreateUserInteractorRequest) *CreateUserInteractorResponse
	UpdateUserProfile(request *UpdateUserInteractorRequest) *UpdateUserInteractorResponse
	DeleteUserProfile(request *DeleteUserInteractorRequest) *DeleteUserInteractorResponse
}

type UserAccountManagementServiceInteractor struct {
	// intentionally left empty
}

func (u *UserAccountManagementServiceInteractor) CreateUserProfile(request *CreateUserInteractorRequest) *CreateUserInteractorResponse {
	db := &postgres.UserProfileDB{}
	resp := db.CreateUserProfile(&postgres.CreateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	return &CreateUserInteractorResponse{
		Username: resp.Username,
		Name:     resp.Name,
		Id:       resp.Id,
		Message:  resp.Message,
		Error:    resp.Error,
	}
}

func (u *UserAccountManagementServiceInteractor) UpdateUserProfile(request *UpdateUserInteractorRequest) *UpdateUserInteractorResponse {
	db := &postgres.UserProfileDB{}
	resp := db.UpdateUserProfile(&postgres.UpdateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	return &UpdateUserInteractorResponse{
		Username: resp.Username,
		Name:     resp.Name,
		Id:       resp.Id,
		Message:  resp.Message,
		Error:    resp.Error,
	}
}

func (u *UserAccountManagementServiceInteractor) DeleteUserProfile(request *DeleteUserInteractorRequest) *DeleteUserInteractorResponse {
	db := &postgres.UserProfileDB{}
	resp := db.DeleteUserProfile(&postgres.DeleteUserDatabaseRequest{
		Id: request.Id,
	})

	return &DeleteUserInteractorResponse{
		Username: resp.Username,
		Name:     resp.Name,
		Id:       resp.Id,
		Message:  resp.Message,
		Error:    resp.Error,
	}
}
