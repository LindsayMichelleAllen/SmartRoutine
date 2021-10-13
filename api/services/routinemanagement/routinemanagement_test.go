package routinemanagement_test

import (
	"api/services/model"
	"api/services/routinemanagement"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutineManagementService(t *testing.T) {
	routineMngmt := &routinemanagement.UnprotectedRoutineService{}
	t.Run("CreateRoutine", func(t *testing.T) {
		t.Run("should return error if user id is not provided", func(t *testing.T) {
			resp := routineMngmt.CreateRoutine(&routinemanagement.RoutineCreateRequest{Name: "RoutineName"})
			assert.Equal(t, resp.Error, errors.New("input field missing"))
			assert.Equal(t, resp.Message, "Input field missing")
		})
		t.Run("should return error if name is not provided", func(t *testing.T) {
			resp := routineMngmt.CreateRoutine(&routinemanagement.RoutineCreateRequest{UserId: "123456789"})
			assert.Equal(t, resp.Error, errors.New("input field missing"))
			assert.Equal(t, resp.Message, "Input field missing")
		})
		t.Run("should create routine given valid input fields", func(t *testing.T) {
			resp := routineMngmt.CreateRoutine(&routinemanagement.RoutineCreateRequest{UserId: "123456789", Name: "RoutineName"})
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			assert.Equal(t, resp.Message, "Not Yet Implemented")
		})
	})
	t.Run("UpdateRoutine", func(t *testing.T) {
		t.Run("should return error if routine id is not provided", func(t *testing.T) {
			resp := routineMngmt.UpdateRoutine(&routinemanagement.RoutineUpdateRequest{Configuration: &model.Configuration{}})
			assert.Equal(t, resp.Error, errors.New("input field missing"))
			assert.Equal(t, resp.Message, "Input field missing")
		})
		t.Run("should return error if routine configuration is not provided", func(t *testing.T) {
			resp := routineMngmt.UpdateRoutine(&routinemanagement.RoutineUpdateRequest{Id: "RoutineID"})
			assert.Equal(t, resp.Error, errors.New("input field missing"))
			assert.Equal(t, resp.Message, "Input field missing")
		})
		t.Run("should create routine given valid input fields", func(t *testing.T) {
			resp := routineMngmt.UpdateRoutine(&routinemanagement.RoutineUpdateRequest{Id: "RoutineID", Configuration: &model.Configuration{}})
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			assert.Equal(t, resp.Message, "Not Yet Implemented")
		})
	})
	t.Run("DeleteRoutine", func(t *testing.T) {
		t.Run("should return error if routine id is not provided", func(t *testing.T) {
			resp := routineMngmt.DeleteRoutine(&routinemanagement.RoutineDeleteRequest{})
			assert.Equal(t, resp.Error, errors.New("input field missing"))
			assert.Equal(t, resp.Message, "Input field missing")
		})
		t.Run("should create routine given valid input fields", func(t *testing.T) {
			resp := routineMngmt.DeleteRoutine(&routinemanagement.RoutineDeleteRequest{Id: "RoutineID"})
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			assert.Equal(t, resp.Message, "Not Yet Implemented")
		})
	})
}
