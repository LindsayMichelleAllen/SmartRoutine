## User Account Management Service

The User Account Management service is responsible for creating, modifying, and deleting user accounts and profiles. This service is also be responsible for user authentication.

---

### User Account Interface

The user account interface shall provide methods to:
* Create a new user
* Modify an existing user
* Delete an existing user

Below is an example of the minimum implementation of the user profile interface. All response objects currently contain the same struct attributes, but these will not be combined to a single response object type to allow for future updates and extensions such as specific error types. 

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

func CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse { ... }
func UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse { ... }
func DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse { ... }

// These methods will be available via a services interactor

/*
type UserAccountManagementServiceInteractor interface {
  CreateUserProfile(request *UserProfileCreateRequest) *UserProfileCreateResponse
  UpdateUserProfile(request *UserProfileUpdateRequest) *UserProfileUpdateResponse
  DeleteUserProfile(request *UserProfileDeleteRequest) *UserProfileDeleteResponse
}

type ServicesInteractor struct {
  AccountManager *UserAccountManagementServiceInteractor
  ...
  ...
}

```
