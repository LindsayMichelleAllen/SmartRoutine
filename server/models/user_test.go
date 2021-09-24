package models_test

import (
	"testing"
	"server/models"
	"github.com/stretchr/testify/assert"
)

func userSetup() *models.User {
	user := &models.User{}
	user.SetUsername("LJam")
	user.SetName("Lindsay")
	user.SetId("123456789")
	return user
}

func TestUserModel(t *testing.T) {
	t.Run("should be able to get the username", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetUsername(), "LJam")
	})
	t.Run("should be able to get the name", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetName(), "Lindsay")
	})
	t.Run("should be able to get the id", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetId(), "123456789")
	})
	t.Run("should be able to set the username", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetUsername(), "LJam")
		user.SetUsername("NewUsername")
		assert.Equal(t, user.GetUsername(), "NewUsername")
	})
	t.Run("should be able to set the name", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetName(), "Lindsay")
		user.SetName("NewName")
		assert.Equal(t, user.GetName(), "NewName")
	})
	t.Run("should be able to set the id", func(t *testing.T){
		user := userSetup()
		assert.Equal(t, user.GetId(), "123456789")
		user.SetId("987654321")
		assert.Equal(t, user.GetId(), "987654321")
	})
}