## User Account Management Service

The User Account Management service is responsible for creating, modifying, and deleting user accounts and profiles. This service is also be responsible for user authentication.

---

### User Account Interface

The user account interface provides the following methods:
```golang
GetUserProfile(request *UserProfileGetRequest) *UserProfileGetResponse
GetUserProfiles() *UserProfilesGetResponse
UserProfileLogin(request *UserProfileLoginRequest) *UserProfileLoginResponse
CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse
UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse
DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse
 ```


#### UserProfile Type
```golang
type UserProfile struct {
	/* username used for login */
	userName string
	/* name displayed to user */
	name string
	/* status of user authenication */
	isAuth bool
}
```
#### Request Types
```golang
type UserProfileGetRequest struct {
	/* unique ID of user */
	Id string
}

type UserProfileLoginRequest struct {
  /* unique ID of user */
	Username string
  /* account password */
	Password string
}

type UserProfileCreateRequest struct {
	/* unique identifier - username provided by user */
	Username string
	/* name displayed to user */
	Name string
	/* password for login */
	Password string
}

type UserProfileUpdateRequest struct {
  /* unique ID of user */
	Username string
  /* name displayed to user */
	Name     string
}

type UserProfileDeleteRequest struct {
	/* unique id of user to be deleted */
	Id string
}

```
#### Response Types
Note: Most of the response types currently contain identical attributes. These were intentionally separated into different types to allow for better adaptability as the application evolves. 
```golang
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

type UserProfileLoginResponse struct {
	User    *model.UserProfile
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
```
### Example Usage

```golang

import (
  usrAcctMngr "api/services/useraccountmanagement"
)

func main() {
  /* Find a specific user */
  
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  userResponse := basicUsrMngr.GetUserProfile(GetUserProfile(&userAcctMngr.UserProfileGetRequest{
    Id: "username",
  })
  
  if userResponse.Error != nil {
    // handle error
  }
  
  usr := userResponse.User
  
  
  /* Get all users */
  
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  userResponse := basicUsrMngr.GetUserProfiles()  
  if userResponse.Error != nil {
    // handle error
  }
  
  usr := userResponse.User
  
  
  /* Login */
  
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  loginResponse := basicUsrMngr.UserProfileLogin(&userAcctMngr.UserProfileLoginRequest{
    Username: "username",
    Password: "password",
  })
  
  if loginResponse.Error != nil {
    // handle error
  }
  
  usr := loginResponse.User
      
  /* Create a new user */
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  newUserResp := basicUsrMngr.CreateUserProfile(&usrAcctMngr.UserProfileCreateRequest{
    Username: "Example Username",
    Password: "Password",
    Name:     "Example Name",
  })

  if newUserResp.Error != nil {
    // handle error
  }
  
  usr := newUserResp.User

  /* Update an existing user */
  
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  updateUserResp := basicUsrMngr.UpdateUserProfile(&usrAcctMngr.UserProfileUpdateRequest{
    Username: "New Username",
    Name:     "New Name",
  })

  if updateUserResp.Error != nil {
    // handle error
  }

  usr := updateUserResp.User

  /* Delete an existing user */
  
  basicUsrMngr := userAcctMngr.UnprotectedUserService{}
  deleteUserResp := basicUsrMngr.DeleteUserProfile(&usrAcctMngr.UserProfileDeleteRequest{
    Id: "username",
  })

  if deleteUserResp.Error != nil {
    // handle error
  }

  removedUser := deleteUserResp.User
}
```


## Future updates
* Maintain authentication state throughout user session
* Authentication for Update & Delete functionality
