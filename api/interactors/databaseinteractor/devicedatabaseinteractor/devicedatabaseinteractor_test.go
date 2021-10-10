package devicedatabaseinteractor_test

import (
	"api/interactors/databaseinteractor/devicedatabaseinteractor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceManagementDBInteractorMethods(t *testing.T) {
	t.Run("CreateDevice", func(t *testing.T) {
		dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
		t.Run("should create device with userid, name, and id", func(t *testing.T) {
			resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
				Id:     "TestDeviceId",
				Name:   "TestDeviceName",
				UserId: "TestUserId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if userid is not provided", func(t *testing.T) {
			resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
				Id:   "TestDeviceId",
				Name: "TestDeviceName",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if name is not provided", func(t *testing.T) {
			resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
				Id:     "TestDeviceId",
				UserId: "TestUserId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {
			resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
				Name:   "TestDeviceName",
				UserId: "TestUserId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
	})
	t.Run("UpdateDevice", func(t *testing.T) {
		dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
		t.Run("should update device with name and id", func(t *testing.T) {
			resp := dbInt.UpdateDevice(&devicedatabaseinteractor.UpdateDeviceRequest{
				Id:   "TestDeviceId",
				Name: "TestDeviceName",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if name is not provided", func(t *testing.T) {
			resp := dbInt.UpdateDevice(&devicedatabaseinteractor.UpdateDeviceRequest{
				Id: "TestDeviceId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {
			resp := dbInt.UpdateDevice(&devicedatabaseinteractor.UpdateDeviceRequest{
				Name: "TestDeviceName",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
	})
	t.Run("DeleteDevice", func(t *testing.T) {
		dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
		t.Run("should delete device with valid id", func(t *testing.T) {
			resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{
				Id: "TestDeviceId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if id is not provided", func(t *testing.T) {
			resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{})
			assert.NotEqual(t, resp.Error, nil)
		})
		t.Run("should return an error if id does not match a known device", func(t *testing.T) {
			resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{
				Id: "UnknownTestDeviceId",
			})
			assert.NotEqual(t, resp.Error, nil)
		})
	})
}
