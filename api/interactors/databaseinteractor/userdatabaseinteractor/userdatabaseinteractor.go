package userdatabaseinteractor

import (
	"api/postgres"
	"api/services/model"
	"errors"
)

type GetUserInteractorRequest struct {
	Id string
}

type GetUserInteractorResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type GetUsersInteractorResponse struct {
	Users   []*model.UserProfile
	Message string
	Error   error
}

type CreateUserInteractorRequest struct {
	Username string
	Name     string
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
	GetUserProfile(request *GetUserInteractorRequest) *GetUserInteractorResponse
	GetUserProfiles() *GetUsersInteractorResponse
	CreateUserProfile(request *CreateUserInteractorRequest) *CreateUserInteractorResponse
	UpdateUserProfile(request *UpdateUserInteractorRequest) *UpdateUserInteractorResponse
	DeleteUserProfile(request *DeleteUserInteractorRequest) *DeleteUserInteractorResponse
}

type UserAccountManagementServiceInteractor struct {
	// intentionally left empty
}

func (u *UserAccountManagementServiceInteractor) GetUserProfile(request *GetUserInteractorRequest) *GetUserInteractorResponse {
	if request.Id == "" {
		return &GetUserInteractorResponse{
			Message: "Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UserProfileDB{}
	resp := db.GetUserProfile(&postgres.GetUserDatabaseRequest{Id: request.Id})
	return &GetUserInteractorResponse{
		User:    resp.User,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UserAccountManagementServiceInteractor) GetUserProfiles() *GetUsersInteractorResponse {
	db := &postgres.UserProfileDB{}
	resp := db.GetUserProfiles()
	return &GetUsersInteractorResponse{
		Users:   resp.Users,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UserAccountManagementServiceInteractor) CreateUserProfile(request *CreateUserInteractorRequest) *CreateUserInteractorResponse {
	db := &postgres.UserProfileDB{}
	resp := db.CreateUserProfile(&postgres.CreateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
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
