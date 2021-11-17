package model_test

import (
	"api/services/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func userSetup() *model.UserProfile {
	user := &model.UserProfile{}
	user.SetUsername("LJam")
	user.SetName("Lindsay")
	user.SetAuthorizationStatus(false)
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
		t.Run("should be able to get the auth status", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, false, user.GetAuthorizationStatus())
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
		t.Run("should be able to set the authorization status", func(t *testing.T) {
			user := userSetup()
			assert.Equal(t, false, user.GetAuthorizationStatus())
			user.SetAuthorizationStatus(true)
			assert.Equal(t, true, user.GetAuthorizationStatus())
		})
	})
}
