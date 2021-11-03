package configurationmanagement_test

import (
	"api/services/routinemanagement/configurationmanagement"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationManagementService(t *testing.T) {
	cnfgMngr := &configurationmanagement.UnprotectedConfigurationService{}
	offset := new(int)
	*offset = 10
	t.Run("CreateConfiguration", func(t *testing.T) {
		t.Run("should return error if offset is not provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationmanagement.CreateConfigurationRequest{
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if device id is not provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationmanagement.CreateConfigurationRequest{
				Offset: offset,
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.CreateConfiguration(&configurationmanagement.CreateConfigurationRequest{
				Offset:   offset,
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
	t.Run("UpdateConfiguration", func(t *testing.T) {
		t.Run("should return error if configuration id is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationmanagement.UpdateConfigurationRequest{
				Offset:   offset,
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if offset is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationmanagement.UpdateConfigurationRequest{
				ConfigId: "ConfigId",
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should return error if device is not provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationmanagement.UpdateConfigurationRequest{
				ConfigId: "ConfigId",
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.UpdateConfiguration(&configurationmanagement.UpdateConfigurationRequest{
				ConfigId: "ConfigId",
				Offset:   offset,
				DeviceId: "DeviceId",
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
	t.Run("DeleteConfiguration", func(t *testing.T) {
		t.Run("should return error if configuration id is not provided", func(t *testing.T) {
			resp := cnfgMngr.DeleteConfiguration(&configurationmanagement.DeleteConfigurationRequest{})
			assert.Equal(t, resp.Message, "Input field(s) missing")
			assert.Equal(t, resp.Error.Error(), "input field(s) missing")
		})
		t.Run("should not return an error if valid inputs are provided", func(t *testing.T) {
			resp := cnfgMngr.DeleteConfiguration(&configurationmanagement.DeleteConfigurationRequest{
				ConfigId: "ConfigId",
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error.Error(), "not yet implemented")
		})
	})
}
