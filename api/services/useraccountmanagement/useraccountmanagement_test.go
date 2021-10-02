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
		t.Run("should be able to create new user with valid input", func(t *testing.T) {
			resp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, resp.Error)
			assert.Equal(t, "LJam", resp.User.GetUsername())
			assert.Equal(t, "Lindsay", resp.User.GetName())
			assert.NotEqual(t, "", resp.User.GetId())
		})
		t.Run("TODO: Create New User Error Testing (blocked by database implementation)", func(t *testing.T) {

		})
	})
	t.Run("Update Existing User", func(t *testing.T) {
		t.Run("should be able to update username", func(t *testing.T) {
			resp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				&userAcctMngr.UserProfile{},
			})
			assert.Equal(t, "Method not yet implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should be able to update name", func(t *testing.T) {
			resp := userAcctMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				&userAcctMngr.UserProfile{},
			})
			assert.Equal(t, "Method not yet implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("TODO: Update Existing User Error Testing (blocked by function not implemented", func(t *testing.T) {

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
		t.Run("TODO: should return error if given id does not match any user accounts", func(t *testing.T) {

		})
		t.Run("TODO: should return error if user is not authorized to delete account", func(t *testing.T) {

		})
	})
}
