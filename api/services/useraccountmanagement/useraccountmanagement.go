package useraccountmanagement

import (
	"api/interactors/databaseinteractor"
	"errors"
)

type UserProfile struct {
	/* username used for login */
	userName string
	/* name displayed to user */
	name string
	/* unique id of user */
	id string
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

type UserProfileCreateResponse struct {
	User    *UserProfile
	Message string
	Error   error
}

type UserProfileUpdateResponse struct {
	User    *UserProfile
	Message string
	Error   error
}

type UserProfileDeleteResponse struct {
	User    *UserProfile
	Message string
	Error   error
}

/* SetUsername overwrites the private username field. */
func (u *UserProfile) SetUsername(newUsername string) {
	u.userName = newUsername
}

/* GetUsername returns the current value of the private username field. */
func (u *UserProfile) GetUsername() string {
	return u.userName
}

/* SetName overwrites the private name field. */
func (u *UserProfile) SetName(newName string) {
	u.name = newName
}

/* GetName returns the current value of the private name field. */
func (u *UserProfile) GetName() string {
	return u.name
}

/* SetId overwrites the private id field. */
func (u *UserProfile) SetId(newId string) {
	u.id = newId
}

/* GetId returns the current value of the private id field. */
func (u *UserProfile) GetId() string {
	return u.id
}

func CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse {
	if request.Username == "" || request.Name == "" {
		return &UserProfileCreateResponse{
			User:    nil,
			Message: "Error encountered: missing input field: Username: " + request.Username + ", Name: " + request.Name,
			Error:   errors.New("missing input field(s)"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}

	resp := dbInteractor.CreateUserProfile(&userdatabaseinteractor.CreateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       "123456789",
	})

	user := &UserProfile{}
	user.SetName(resp.Name)
	user.SetUsername(resp.Username)
	user.SetId(resp.Id)

	return &UserProfileCreateResponse{
		User:    user,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse {
	if request.Username == "" || request.Name == "" || request.Id == "" {
		return &UserProfileUpdateResponse{
			User:    nil,
			Message: "Error encountered: missing input field: Username: " + request.Username + ", Name: " + request.Name + ", Id: " + request.Id,
			Error:   errors.New("missing input field(s)"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.UpdateUserProfile(&userdatabaseinteractor.UpdateUserDatabaseRequest{
		Username: request.Username,
		Name:     request.Name,
		Id:       request.Id,
	})

	usr := &UserProfile{}
	usr.SetUsername(resp.Username)
	usr.SetName(resp.Name)
	usr.SetId(resp.Id)

	return &UserProfileUpdateResponse{
		User:    usr,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse {
	if request.Id == "" {
		return &UserProfileDeleteResponse{
			User:    nil,
			Message: "Error encountere: missing id input field",
			Error:   errors.New("missing input field"),
		}
	}

	dbInteractor := &userdatabaseinteractor.UserAccountManagementServiceInteractor{}
	resp := dbInteractor.DeleteUserProfile(&userdatabaseinteractor.DeleteUserDatabaseRequest{
		Id: "123456789",
	})

	user := &UserProfile{}
	user.SetUsername(resp.Username)
	user.SetName(resp.Name)
	user.SetId(resp.Id)

	return &UserProfileDeleteResponse{
		User:    user,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
