package userdatabaseinteractor

import (
	db "postgres"
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
	resp := db.CreateUserProfile(&db.CreateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	if resp.Error != nil {
		return &CreateUserInteractorResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return resp
}

func (u *UserAccountManagementServiceInteractor) UpdateUserProfile(request *UpdateUserInteractorRequest) *UpdateUserInteractorResponse {
	resp := db.UpdateUserProfile(&db.UpdateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	if resp.Error != nil {
		return &UpdateUserInteractorResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return resp
}

func (u *UserAccountManagementServiceInteractor) DeleteUserProfile(request *DeleteUserInteractorRequest) *DeleteUserInteractorResponse {
	resp := db.CreateUserProfile(&db.DeleteUserDatabaseRequest{
		Id: request.Id,
	})

	if resp.Error != nil {
		return &DeleteUserInteractorResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return resp
}
