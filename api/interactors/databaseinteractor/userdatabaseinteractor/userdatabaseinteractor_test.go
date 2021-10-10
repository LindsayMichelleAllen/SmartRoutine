package userdatabaseinteractor_test

import (
	"testing"
)

func TestUserAccountManagementDBInteractorMethods(t *testing.T) {
	t.Run("CreateUserProfile", func(t *testing.T) {
		t.Run("should create user profile with username, name, and id", func(t *testing.T) {

		})
		t.Run("should return an error if username is not provided", func(t *testing.T) {

		})
		t.Run("should return an error if name is not provided", func(t *testing.T) {

		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {

		})
	})
	t.Run("UpdateUserProfile", func(t *testing.T) {
		t.Run("should update user profile with username, name, and id", func(t *testing.T) {

		})
		t.Run("should return an error if username is not provided", func(t *testing.T) {

		})
		t.Run("should return an error if name is not provided", func(t *testing.T) {

		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {

		})
	})
	t.Run("DeleteUserProfile", func(t *testing.T) {
		t.Run("should delete user profile given a valid id", func(t *testing.T) {

		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {

		})
	})
}
