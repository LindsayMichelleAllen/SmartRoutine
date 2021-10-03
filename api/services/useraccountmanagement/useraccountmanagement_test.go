package useraccountmanagement_test

import (
	userAcctMngr "api/services/useraccountmanagement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func userSetup() *userAcctMngr.UserProfile {
	user := &userAcctMngr.UserProfile{}
	user.SetUsername("LJam")
	user.SetName("Lindsay")
	user.SetId("123456789")
	return user
}

func TestUserProfileMethods(t *testing.T) {
	t.Run("Getters", func(t *testing.T) {
		t.Run("should be able to get the username", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "LJam", user.GetUsername())
		})
		t.Run("should be able to get the name", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "Lindsay", user.GetName())
		})
		t.Run("should be able to get the id", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "123456789", user.GetId())
		})
	})
	t.Run("Setters", func(t *testing.T) {
		t.Run("should be able to set the username", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "LJam", user.GetUsername())
			user.SetUsername("NewUsername")
			assert.Equal(t, "NewUsername", user.GetUsername())
		})
		t.Run("should be able to set the name", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "Lindsay", user.GetName())
			user.SetName("NewName")
			assert.Equal(t, "NewName", user.GetName())
		})
		t.Run("should be able to set the id", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, "123456789", user.GetId())
			user.SetId("987654321")
			assert.Equal(t, "987654321", user.GetId())
		})
	})
}

func TestUserAccountManagementInteractorMethods(t *testing.T) {
	t.Run("Create New User", func(t *testing.T) {
		t.Run("should be able to create new user with valid name and username", func(t *testing.T) {
			resp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, resp.Error)
			assert.Equal(t, "LJam", resp.User.GetUsername())
			assert.Equal(t, "Lindsay", resp.User.GetName())
			assert.NotEqual(t, "", resp.User.GetId())
		})
		t.Run("should return error is username is not provided", func(t *testing.T) {
			resp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Name: "Lindsay",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error is name is not provided", func(t *testing.T) {
			resp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
	t.Run("Update Existing User", func(t *testing.T) {
		t.Run("should be able to update username", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Name:     createResp.User.GetName(),
				Id:       createResp.User.GetId(),
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "LJam Supreme", updateResp.User.GetUsername())
			assert.Equal(t, createResp.User.GetName(), updateResp.User.GetName())
			assert.Equal(t, createResp.User.GetId(), updateResp.User.GetId())
		})
		t.Run("should be able to update name", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: createResp.User.GetUsername(),
				Name:     "Lindsay Allen",
				Id:       createResp.User.GetId(),
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "Lindsay Allen", updateResp.User.GetName())
			assert.Equal(t, createResp.User.GetUsername(), updateResp.User.GetUsername())
			assert.Equal(t, createResp.User.GetId(), updateResp.User.GetId())
		})
		t.Run("should be able to update username and name", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Name:     "Lindsay Allen",
				Id:       createResp.User.GetId(),
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "Lindsay Allen", updateResp.User.GetName())
			assert.Equal(t, "LJam Supreme", updateResp.User.GetUsername())
			assert.Equal(t, createResp.User.GetId(), updateResp.User.GetId())
		})
		t.Run("should return error if username is not provided", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Name: "Lindsay Allen",
				Id:   createResp.User.GetId(),
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Id:       createResp.User.GetId(),
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			createResp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Name:     "Lindsay Allen",
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("TODO should return error if user is not authorized to modify user profile", func(t *testing.T) {

		})
	})
	t.Run("Delete Existing User", func(t *testing.T) {
		t.Run("should be able to delete user given a valid id", func(t *testing.T) {
			resp := userAcctMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{
				Id: "123456789",
			})
			assert.Equal(t, "Method not yet implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			resp := userAcctMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("TODO: should return error if given id does not match any user accounts", func(t *testing.T) {

		})
		t.Run("TODO: should return error if user is not authorized to delete account", func(t *testing.T) {

		})
	})
}
