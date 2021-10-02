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
			assert.Equal(t, user.GetUsername(), "LJam")
		})
		t.Run("should be able to get the name", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, user.GetName(), "Lindsay")
		})
		t.Run("should be able to get the id", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, user.GetId(), "123456789")
		})
	})
	t.Run("Setters", func(t *testing.T) {
		t.Run("should be able to set the username", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, user.GetUsername(), "LJam")
			user.SetUsername("NewUsername")
			assert.Equal(t, user.GetUsername(), "NewUsername")
		})
		t.Run("should be able to set the name", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, user.GetName(), "Lindsay")
			user.SetName("NewName")
			assert.Equal(t, user.GetName(), "NewName")
		})
		t.Run("should be able to set the id", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, user.GetId(), "123456789")
			user.SetId("987654321")
			assert.Equal(t, user.GetId(), "987654321")
		})
	})
	t.Run("Helper Methods", func(t *testing.T) {
		t.Run("should be able to create user profile", func(t *testing.T) {
			resp := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, resp.Error, nil)
			assert.Equal(t, resp.User.GetUsername(), "LJam")
			assert.Equal(t, resp.User.GetName(), "Lindsay")
			assert.NotEqual(t, resp.User.GetId(), "")
		})
	})
}
