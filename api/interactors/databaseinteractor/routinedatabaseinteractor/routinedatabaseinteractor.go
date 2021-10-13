package routinedatabaseinteractor

import (
	"api/services/model"
	"errors"
)

type RoutineCreateDatabaseRequest struct {
	Id            string
	Name          string
	UserId        string
	Configuration *model.Configuration
}

type RoutineUpdateDatabaseRequest struct {
	Id            string
	Configuration *model.Configuration
}

type RoutineDeleteDatabaseRequest struct {
	Id string
}

type RoutineCreateDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineUpdateDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineDeleteDatabaseResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineDBInteractor interface {
	CreateRoutine(request *RoutineCreateDatabaseRequest) *RoutineCreateDatabaseResponse
	UpdateRoutine(request *RoutineUpdateDatabaseRequest) *RoutineUpdateDatabaseResponse
	DeleteRoutine(request *RoutineDeleteDatabaseRequest) *RoutineDeleteDatabaseResponse
}

type UnprotectedRoutineDBInteractor struct {
	// intentionally left empty
}

func (r *UnprotectedRoutineDBInteractor) CreateRoutine(request *RoutineCreateDatabaseRequest) *RoutineCreateDatabaseResponse {
	if request.Id == "" || request.Name == "" || request.UserId == "" || request.Configuration == nil {
		return &RoutineCreateDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	return &RoutineCreateDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (r *UnprotectedRoutineDBInteractor) UpdateRoutine(request *RoutineUpdateDatabaseRequest) *RoutineUpdateDatabaseResponse {
	if request.Id == "" || request.Configuration == nil {
		return &RoutineUpdateDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	return &RoutineUpdateDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (r *UnprotectedRoutineDBInteractor) DeleteRoutine(request *RoutineDeleteDatabaseRequest) *RoutineDeleteDatabaseResponse {
	if request.Id == "" {
		return &RoutineDeleteDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	return &RoutineDeleteDatabaseResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
