package model_test

import (
	"api/services/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupConfiguration() *model.Configuration {
	config := &model.Configuration{}
	device := &model.Device{}
	device.SetId("DeviceID")
	device.SetName("DeviceName")
	device.SetUserId("DeviceUserID")
	config.SetId("ConfigurationID")
	config.SetRoutineId("RoutineID")
	config.SetOffset(5)
	config.SetDevice(device)
	return config
}

func TestConfigurationMethods(t *testing.T) {
	config := setupConfiguration()
	t.Run("Getters", func(t *testing.T) {
		t.Run("should be able to get configuration id", func(t *testing.T) {
			assert.Equal(t, config.GetId(), "ConfigurationID")
		})
		t.Run("should be able to get configuration offset", func(t *testing.T) {
			assert.Equal(t, config.GetOffset(), 5)
		})
		t.Run("should be able to get configuration device", func(t *testing.T) {
			assert.Equal(t, config.GetDevice().GetId(), "DeviceID")
			assert.Equal(t, config.GetDevice().GetName(), "DeviceName")
			assert.Equal(t, config.GetDevice().GetUserId(), "DeviceUserID")
		})
		t.Run("should be able to get routine id", func(t *testing.T) {
			assert.Equal(t, config.GetRoutineId(), "RoutineID")
		})
	})
	t.Run("Setters", func(t *testing.T) {
		t.Run("should be able to set configuration id", func(t *testing.T) {
			config.SetId("NewConfigurationID")
			assert.Equal(t, config.GetId(), "NewConfigurationID")
		})
		t.Run("should be able to set configuration offset", func(t *testing.T) {
			config.SetOffset(10)
			assert.Equal(t, config.GetOffset(), 10)
		})
		t.Run("should be able to set configuration device", func(t *testing.T) {
			device := &model.Device{}
			device.SetId("NewlyCreatedDeviceID")
			device.SetName("NewlyCreatedDeviceName")
			device.SetUserId("NewlyCreatedDeviceUserID")
			config.SetDevice(device)
			assert.Equal(t, config.GetDevice().GetId(), "NewlyCreatedDeviceID")
			assert.Equal(t, config.GetDevice().GetName(), "NewlyCreatedDeviceName")
			assert.Equal(t, config.GetDevice().GetUserId(), "NewlyCreatedDeviceUserID")
		})
		t.Run("should be able to set routine id", func(t *testing.T) {
			assert.Equal(t, config.GetRoutineId(), "RoutineID")
			config.SetRoutineId("NewRoutineID")
			assert.Equal(t, config.GetRoutineId(), "NewRoutineID")
		})
	})
}
