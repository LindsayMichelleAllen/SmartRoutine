package devicemanagement_test

import (
	dvcMngr "api/services/devicemanagement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAccountManagementInteractorMethods(t *testing.T) {
	t.Run("Create New User", func(t *testing.T) {
		t.Run("should be able to create new device with valid device name and user id", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				Name:   "DevName",
				UserId: "123456789",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if user id is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				Name: "LJam",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
	t.Run("Update Existing User", func(t *testing.T) {
		t.Run("should be able to update username", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
				Name:   "DevName",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicDvcMngr.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
				Name: "NewDevName",
				Id:   createResp.Device.GetId(),
			})
			assert.Equal(t, nil, updateResp.Error)
			assert.Equal(t, "LJam Supreme", updateResp.Device.GetName())
			assert.Equal(t, createResp.Device.GetName(), updateResp.Device.GetName())
			assert.Equal(t, createResp.Device.GetId(), updateResp.Device.GetId())
		})
		t.Run("should be able to update device name", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
				Name:   "DevName",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicDvcMngr.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
				Name: "NewDevName",
				Id:   createResp.Device.GetId(),
			})
			assert.Equal(t, nil, updateResp.Error)
		})
		t.Run("should be able to update username and name", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
				Name:   "DevName",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
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
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Name: "Lindsay Allen",
				Id:   createResp.User.GetId(),
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: "LJam",
				Name:     "Lindsay",
			})
			assert.Equal(t, nil, createResp.Error)

			updateResp := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: "LJam Supreme",
				Id:       createResp.User.GetId(),
			})
			assert.NotEqual(t, nil, updateResp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
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
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicUsrMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{
				Id: "123456789",
			})
			assert.Equal(t, "Method not yet implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicUsrMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
}
