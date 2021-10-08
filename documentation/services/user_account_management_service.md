## User Account Management Service

The User Account Management service is responsible for creating, modifying, and deleting user accounts and profiles. This service is also be responsible for user authentication.

---

### User Account Interface

The user account interface shall provide methods to:
* Create a new user
* Modify an existing user
* Delete an existing user

Below is an example of the minimum implementation of the user account management service. All response objects currently contain the same struct attributes, but these will not be combined to a single response object type to allow for future updates and extensions.

```golang
type UserProfile struct {
  Username string
  Name     string
  id       string
}

type UserProfileCreateRequest struct {
  Username string
  Name     string
}

type UserProfileUpdateRequest struct {
  User *UserProfile
}

type UserProfileDeleteRequest struct {
  Id string
}

type UserProfileCreateResponse struct {
  User *UserProfile
  Message string
  Error error
}

type UserProfileUpdateResponse struct {
  User *UserProfile
  Message string
  Error error
}

type UserProfileDeleteResponse struct {
  User *UserProfile
  Message string
  Error error
}

// These methods will be available via a services interactor

type UserAccountManagementServiceInteractor struct {}

func (u *UserAccouneManagementServiceInteractor)CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse { ... }
func (u *UserAccouneManagementServiceInteractor)UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse { ... }
func (u *UserAccouneManagementServiceInteractor)DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse { ... }


type ServicesInteractor struct {
  AccountManager *UserAccountManagementServiceInteractor
  ...
  ...
}

```
### Example Usage

```golang

import (
  usrAcctMngr "api/services/useraccountmanagement"
)

func main() {
  /* Create a new user */
  newUserResp := usrAcctMngr.CreateUserProfile(&usrAcctMngr.UserProfileCreateRequest{
    Username: "Example Username",
    Name:     "Example Name",
  })

  if newUserResp.Error != nil {
    // handle error
  }
  
  usr := newUserResp.User

  /* Update an existing user */
  updateUserResp := usrAcctMngr.UpdateUserProfile(&usrAcctMngr.UserProfileUpdateRequest{
    Username: "New Username",
    Name:     "New Name",
    Id:       "UserId",
  })

  if updateUserResp.Error != nil {
    // handle error
  }

  usr := updateUserResp.User

  /* Delete an existing user */
  deleteUserResp := usrAcctMngr.DeleteUserProfile(&usrAcctMngr.UserProfileDeleteRequest{
    Id: "UserId",
  })

  if deleteUserResp.Error != nil {
    // handle error
  }

  removedUser := deleteUserResp.User
}
```


Future updates for the User Account Management Service include 
* Password protection
* Authentication for Update & Delete functionality
