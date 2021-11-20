package model_test

import (
	"api/services/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRoutine() *model.Routine {
	routine := &model.Routine{}
	routine.SetId("123456789")
	routine.SetName("RoutineName")
	routine.SetConfiguration(make([]*model.Configuration, 0))
	return routine
}

func TestRoutineMethods(t *testing.T) {
	routine := setupRoutine()
	t.Run("Getters", func(t *testing.T) {
		t.Run("should be able to get routine id", func(t *testing.T) {
			assert.Equal(t, routine.GetId(), "123456789")
		})
		t.Run("should be able to get routine name", func(t *testing.T) {
			assert.Equal(t, routine.GetName(), "RoutineName")
		})
		t.Run("should be able to get routine configuration", func(t *testing.T) {
			assert.Equal(t, routine.GetConfiguration(), make([]*model.Configuration, 0))
		})
	})
	t.Run("Setters", func(t *testing.T) {
		t.Run("should be able to set routine id", func(t *testing.T) {
			routine.SetId("NewRoutineId")
			assert.Equal(t, routine.GetId(), "NewRoutineId")
		})
		t.Run("should be able to set routine id", func(t *testing.T) {
			routine.SetName("NewRoutineName")
			assert.Equal(t, routine.GetName(), "NewRoutineName")
		})
		t.Run("should be able to set routine configuration", func(t *testing.T) {
			config := model.Configuration{}
			slc := make([]*model.Configuration, 0)
			slc = append(slc, &config)
			routine.SetConfiguration(slc)
			assert.Equal(t, routine.GetConfiguration(), slc)
			routine.ClearConfigurations()
		})
		t.Run("should be able to add a single routine configuration", func(t *testing.T) {
			config := &model.Configuration{}
			dev := &model.Device{}
			dev.SetId("DeviceID")
			dev.SetName("DeviceName")
			dev.SetUserId("DeviceUserID")
			config.SetId("ConfigID")
			config.SetDevice(dev)
			routine.AddToConfiguration(config)
			assert.Equal(t, len(routine.GetConfiguration()), 1)
			config2 := &model.Configuration{}
			dev2 := &model.Device{}
			dev2.SetId("DeviceID2")
			dev2.SetName("DeviceName2")
			dev2.SetUserId("DeviceUserID2")
			config2.SetId("ConfigID2")
			config2.SetDevice(dev)
			routine.AddToConfiguration(config2)
			assert.Equal(t, len(routine.GetConfiguration()), 2)
			slc := make([]*model.Configuration, 0)
			slc = append(slc, config)
			slc = append(slc, config2)
			assert.Equal(t, routine.GetConfiguration(), slc)
		})
		t.Run("should be able to clear configuration", func(t *testing.T) {
			routine.ClearConfigurations()
			assert.Equal(t, len(routine.GetConfiguration()), 0)
		})
	})
}
