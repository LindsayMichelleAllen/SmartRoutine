package useraccountmanagement

import (
	"api/interactors/databaseinteractor/userdatabaseinteractor"
	"api/services/model"
	"errors"
)

type UserProfileGetRequest struct {
	/* unique ID of user */
	Id string
}

type UserProfileCreateRequest struct {
	/* username provided by user */
	Username string
	/* name provided by user */
	Name string
}

type UserProfileUpdateRequest struct {
	Username string
	Name     string
	Id       string
}

type UserProfileDeleteRequest struct {
	/* unique id of user to be deleted */
	Id string
}

type UserProfileGetResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type UserProfilesGetResponse struct {
	Users   []*model.UserProfile
	Message string
	Error   error
}

type UserProfileCreateResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type UserProfileUpdateResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type UserProfileDeleteResponse struct {
	User    *model.UserProfile
	Message string
	Error   error
}

type UserService interface {
	GetUserProfile(request *UserProfileGetRequest) *UserProfileGetResponse
	GetUserProfiles() *UserProfilesGetResponse
	CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse
	UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse
	DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse
}

type UnprotectedUserService struct {
	// intentionally left empty
}

func (u *UnprotectedUserService) GetUserProfile(request *UserProfileGetRequest) *UserProfileGetResponse {
	if request.Id == "" {
		return &UserProfileGetResponse{
			Message: "Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.GetUserProfile(&userdatabaseinteractor.GetUserInteractorRequest{Id: request.Id})
	return &UserProfileGetResponse{
		User:    resp.User,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UnprotectedUserService) GetUserProfiles() *UserProfilesGetResponse {
	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.GetUserProfiles()
	return &UserProfilesGetResponse{
		Users:   resp.Users,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UnprotectedUserService) CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse {
	if request.Username == "" || request.Name == "" {
		return &UserProfileCreateResponse{
			User:    nil,
			Message: "Error encountered: missing input field: Username: " + request.Username + ", Name: " + request.Name,
			Error:   errors.New("missing input field(s)"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	// TODO randomly generate the user id
	resp := dbInteractor.CreateUserProfile(&userdatabaseinteractor.CreateUserInteractorRequest{
		Username: request.Username,
		Name:     request.Name,
	})

	user := &model.UserProfile{}
	user.SetName(resp.Name)
	user.SetUsername(resp.Username)
	user.SetId(resp.Id)

	return &UserProfileCreateResponse{
		User:    user,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UnprotectedUserService) UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse {
	if request.Username == "" || request.Name == "" || request.Id == "" {
		return &UserProfileUpdateResponse{
			User:    nil,
			Message: "Error encountered: missing input field: Username: " + request.Username + ", Name: " + request.Name + ", Id: " + request.Id,
			Error:   errors.New("missing input field(s)"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.UpdateUserProfile(&userdatabaseinteractor.UpdateUserInteractorRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	usr := &model.UserProfile{}
	usr.SetUsername(resp.Username)
	usr.SetName(resp.Name)
	usr.SetId(resp.Id)

	return &UserProfileUpdateResponse{
		User:    usr,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (u *UnprotectedUserService) DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse {
	if request.Id == "" {
		return &UserProfileDeleteResponse{
			User:    nil,
			Message: "Error encountere: missing id input field",
			Error:   errors.New("missing input field"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.DeleteUserProfile(&userdatabaseinteractor.DeleteUserInteractorRequest{
		Id: "123456789",
	})

	user := &model.UserProfile{}
	user.SetUsername(resp.Username)
	user.SetName(resp.Name)
	user.SetId(resp.Id)

	return &UserProfileDeleteResponse{
		User:    user,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
