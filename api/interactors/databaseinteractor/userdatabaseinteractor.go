package userdatabaseinteractor

import (
	"errors"
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
	resp, err := db.CreateUserProfile(&db.CreateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	if err != nil {
		return &CreateUserInteractorResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return &CreateUserInteractorResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (u *UserAccountManagementServiceInteractor) UpdateUserProfile(request *UpdateUserInteractorRequest) *UpdateUserInteractorResponse {
	return &UpdateUserInteractorResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (u *UserAccountManagementServiceInteractor) DeleteUserProfile(request *DeleteUserInteractorRequest) *DeleteUserInteractorResponse {
	return &DeleteUserInteractorResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
