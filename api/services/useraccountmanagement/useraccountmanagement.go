package useraccountmanagement

import "errors"

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
	User *UserProfile
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

/*
SetUsername overwrites the private username field.
*/
func (u *UserProfile) SetUsername(newUsername string) {
	u.userName = newUsername
}

/*
GetUsername returns the current value of the private username field.
*/
func (u *UserProfile) GetUsername() string {
	return u.userName
}

/*
SetName overwrites the private name field.
*/
func (u *UserProfile) SetName(newName string) {
	u.name = newName
}

/*
GetName returns the current value of the private name field.
*/
func (u *UserProfile) GetName() string {
	return u.name
}

/*
SetId overwrites the private id field.
*/
func (u *UserProfile) SetId(newId string) {
	u.id = newId
}

/*
GetId returns the current value of the private id field.
*/
func (u *UserProfile) GetId() string {
	return u.id
}

func CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse {
	user := &UserProfile{}
	user.SetName(request.Name)
	user.SetUsername(request.Username)
	// TODO update to randomly generate userId
	user.SetId("123456789")
	// pass user object to database handler
	return &UserProfileCreateResponse{
		User:    user,
		Message: "User Profile Successfully Created - Not Yet Added To Database",
		Error:   nil,
	}
}

func UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse {
	return &UserProfileUpdateResponse{
		User:    nil,
		Message: "Method not yet implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse {
	return &UserProfileDeleteResponse{
		User:    nil,
		Message: "Method not yet implemented",
		Error:   errors.New("not yet implemented"),
	}
}
