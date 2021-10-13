package routinedatabaseinteractor_test

import (
	"api/interactors/databaseinteractor/routinedatabaseinteractor"
	"api/services/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutineDBInteractor(t *testing.T) {
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	t.Run("CreateRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Name:          "RoutineName",
				UserId:        "RoutineUserID",
				Configuration: &model.Configuration{}})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if routine name is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Id:            "RoutineID",
				UserId:        "RoutineUserID",
				Configuration: &model.Configuration{}})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if user id is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Id:            "RoutineID",
				Name:          "RoutineName",
				Configuration: &model.Configuration{},
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if configuration is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Id:     "RoutineID",
				Name:   "RoutineID",
				UserId: "RoutineUserID",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to create routine with valid input fields", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Id:            "RoutineID",
				Name:          "RoutineName",
				UserId:        "RoutineUserID",
				Configuration: &model.Configuration{},
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
		})
	})
	t.Run("UpdateRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateDatabaseRequest{
				Configuration: &model.Configuration{},
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if routine configuration is not provided", func(t *testing.T) {
			resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateDatabaseRequest{
				Id: "RoutineID",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to create routine with valid input fields", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
				Id:            "RoutineID",
				Configuration: &model.Configuration{},
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
		})
	})
	t.Run("DeleteRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteDatabaseRequest{})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to create routine with valid input fields", func(t *testing.T) {
			resp := dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteDatabaseRequest{
				Id: "RoutineID",
			})
			assert.Equal(t, resp.Message, "Not Yet Implemented")
			assert.Equal(t, resp.Error, errors.New("not yet implemented"))
		})
	})
}
