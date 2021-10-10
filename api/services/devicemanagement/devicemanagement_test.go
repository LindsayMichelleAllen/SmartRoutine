package devicemanagement_test

import (
	dvcMngr "api/services/devicemanagement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceManagementInteractorMethods(t *testing.T) {
	t.Run("Create New Device", func(t *testing.T) {
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
				Name: "DeviceName",
			})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
	t.Run("Update Existing Device", func(t *testing.T) {
		t.Run("should be able to update device name", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
				Name:   "DevName",
			})
			assert.NotEqual(t, nil, createResp.Error)
			/*
				updateResp := basicDvcMngr.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
					Name: "NewDevName",
					Id:   createResp.Device.GetId(),
				})
				assert.NotEqual(t, nil, updateResp.Error)
				assert.Equal(t, "NewDevName", updateResp.Device.GetName())
				assert.Equal(t, createResp.Device.GetId(), updateResp.Device.GetId())
				assert.Equal(t, createResp.Device.GetUserId(), updateResp.Device.GetUserId())
			*/
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				UserId: "123456789",
				Name:   "DevName",
			})
			assert.NotEqual(t, nil, createResp.Error)
			/*
				updateResp := basicDvcMngr.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
					Id: createResp.Device.GetId(),
				})
				assert.NotEqual(t, nil, updateResp.Error)
			*/
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			createResp := basicDvcMngr.CreateDevice(&dvcMngr.DeviceCreateRequest{
				Name:   "Lindsay",
				UserId: "123456789",
			})
			assert.NotEqual(t, nil, createResp.Error)
			/*
				updateResp := basicDvcMngr.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
					Name: "DevName",
				})
				assert.NotEqual(t, nil, updateResp.Error)
			*/
		})
	})
	t.Run("Delete Existing Device", func(t *testing.T) {
		t.Run("should be able to delete user given a valid id", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcMngr.DeleteDevice(&dvcMngr.DeviceDeleteRequest{
				Id: "123456789",
			})
			assert.Equal(t, "Not Yet Implemented", resp.Message)
			assert.NotEqual(t, nil, resp.Error)
		})
		t.Run("should return error if id is not provided", func(t *testing.T) {
			basicDvcMngr := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcMngr.DeleteDevice(&dvcMngr.DeviceDeleteRequest{})
			assert.NotEqual(t, nil, resp.Error)
		})
	})
}
