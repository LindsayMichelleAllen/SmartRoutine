package model_test

import (
	"api/services/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRoutine() *model.Routine {
	routine := &model.Routine{}
	routine.SetId("123456789")
	routine.SetConfiguration(make([]*model.Configuration, 0))
	return routine
}

func TestRoutineMethods(t *testing.T) {
	routine := setupRoutine()
	t.Run("Getters", func(t *testing.T) {
		t.Run("should be able to get routine id", func(t *testing.T) {
			assert.Equal(t, routine.GetId(), "123456789")
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
		t.Run("should be able to set routine configuration", func(t *testing.T) {
			config := model.Configuration{}
			slc := make([]*model.Configuration, 0)
			slc = append(slc, &config)
			routine.SetConfiguration(slc)
			assert.Equal(t, routine.GetConfiguration(), slc)
		})
	})
}
