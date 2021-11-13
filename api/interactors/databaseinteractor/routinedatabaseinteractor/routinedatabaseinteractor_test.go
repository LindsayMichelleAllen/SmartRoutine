package routinedatabaseinteractor_test

import (
	"api/interactors/databaseinteractor/routinedatabaseinteractor"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutineDBInteractor(t *testing.T) {
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	t.Run("CreateRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateInteractorRequest{
				Name:   "RoutineName",
				UserId: "RoutineUserID",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if routine name is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateInteractorRequest{
				UserId: "RoutineUserID",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if user id is not provided", func(t *testing.T) {
			resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateInteractorRequest{
				Name: "RoutineName",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to create routine with valid input fields", func(t *testing.T) {
			_ = dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateInteractorRequest{
				Name:   "RoutineName",
				UserId: "RoutineUserID",
			})
			/*
				assert.Equal(t, resp.Message, "Not Yet Implemented")
				assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			*/
		})
	})
	t.Run("UpdateRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateInteractorRequest{
				Name: "NewRoutineName",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should return an error if routine name is not provided", func(t *testing.T) {
			resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateInteractorRequest{
				Id: "RoutineID",
			})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to update routine with valid input fields", func(t *testing.T) {
			_ = dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateInteractorRequest{
				Id:   "RoutineID",
				Name: "RoutineName",
			})
			/*
				assert.Equal(t, resp.Message, "Not Yet Implemented")
				assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			*/
		})
	})
	t.Run("DeleteRoutine", func(t *testing.T) {
		t.Run("should return an error if routine id is not provided", func(t *testing.T) {
			resp := dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteInteractorRequest{})
			assert.Equal(t, resp.Message, "Input field missing")
			assert.Equal(t, resp.Error, errors.New("input field missing"))
		})
		t.Run("should be able to create routine with valid input fields", func(t *testing.T) {
			_ = dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteInteractorRequest{
				Id: "RoutineID",
			})
			/*
				assert.Equal(t, resp.Message, "Not Yet Implemented")
				assert.Equal(t, resp.Error, errors.New("not yet implemented"))
			*/
		})
	})
}
