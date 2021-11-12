package configurationdatabaseinteractor_test

import (
	"api/interactors/databaseinteractor/configurationdatabaseinteractor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationDBInteractor(t *testing.T) {
	cnfgMngr := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	offset := new(int)
	*offset = 10
	t.Run("CreateConfiguration", func(t *testing.T) {
		t.Run("should return error if configuration id is not provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationdatabaseinteractor.CreateConfigurationDBInteractorRequest{
				Offset:   offset,
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if offset is not provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationdatabaseinteractor.CreateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if device is not provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationdatabaseinteractor.CreateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationdatabaseinteractor.CreateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
	t.Run("UpdateConfiguration", func(t *testing.T) {
		t.Run("should return error if configuration id is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationdatabaseinteractor.UpdateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if offset is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationdatabaseinteractor.UpdateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if device is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationdatabaseinteractor.UpdateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationdatabaseinteractor.UpdateConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
	t.Run("DeleteConfiguration", func(t *testing.T) {
		t.Run("should return error if configuration id is not provided", func(t *testing.T) {
			resp := cnfgMngr.DeleteConfiguration(&configurationdatabaseinteractor.DeleteConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.DeleteConfiguration(&configurationdatabaseinteractor.DeleteConfigurationDBInteractorRequest{})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
}
