package useraccountmanagement_test

import (
	userAcctMngr "api/services/useraccountmanagement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAccountManagementServiceMethods(t *testing.T) {
	t.Run("Create New User", func(t *testing.T) {
		t.Run("should be able to create new user with valid name and username", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			resp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, resp.Error)
			assert.Equal(t, "LJam", resp.User.GetUsername())
			assert.Equal(t, "Lindsay", resp.User.GetName())
		})
		t.Run("should return error if username is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			resp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Name: "Lindsay",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			resp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
	t.Run("Update Existing User", func(t *testing.T) {
		t.Run("should be able to update username", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Name:     createResp.User.GetName(),
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "LJam Supreme", updateResp.User.GetUsername())
			assert.Equal(t, createResp.User.GetName(), updateResp.User.GetName())
		})
		t.Run("should be able to update name", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: createResp.User.GetUsername(),
				Name:     "Lindsay Allen",
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "Lindsay Allen", updateResp.User.GetName())
			assert.Equal(t, createResp.User.GetUsername(), updateResp.User.GetUsername())
		})
		t.Run("should be able to update username and name", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Name:     "Lindsay Allen",
			})
			assert.Equal(t, "Successfully Updated User Profile", updateResp.Message)
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "Lindsay Allen", updateResp.User.GetName())
			assert.Equal(t, "LJam Supreme", updateResp.User.GetUsername())
		})
		t.Run("should return error if username is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Name: "Lindsay Allen",
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
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
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			resp := basicUsrMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{
				Id: "123456789",
			})
			assert.Equal(t, "Method not yet implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			resp := basicUsrMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("TODO: should return error if given id does not match any user accounts", func(t *testing.T) {

		})
		t.Run("TODO: should return error if user is not authorized to delete account", func(t *testing.T) {

		})
	})
}
