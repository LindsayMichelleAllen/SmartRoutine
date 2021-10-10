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
