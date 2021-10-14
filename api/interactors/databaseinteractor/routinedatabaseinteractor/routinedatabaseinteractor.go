package routinedatabaseinteractor

import (
	"api/postgres"
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

	routine := &model.Routine{}
	routine.SetId(request.Id)
	routine.SetName(request.Name)
	routine.SetUserId(request.UserId)

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.CreateRoutine(&postgres.CreateRoutineDatabaseRequest{Routine: routine})

	return &RoutineCreateDatabaseResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineDBInteractor) UpdateRoutine(request *RoutineUpdateDatabaseRequest) *RoutineUpdateDatabaseResponse {
	if request.Id == "" || request.Configuration == nil {
		return &RoutineUpdateDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	routine := &model.Routine{}
	routine.SetId(request.Id)
	routine.AddToConfiguration(request.Configuration)

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.UpdateRoutine(&postgres.UpdateRoutineDatabaseRequest{Routine: routine})

	return &RoutineUpdateDatabaseResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineDBInteractor) DeleteRoutine(request *RoutineDeleteDatabaseRequest) *RoutineDeleteDatabaseResponse {
	if request.Id == "" {
		return &RoutineDeleteDatabaseResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	routine := &model.Routine{}
	routine.SetId(request.Id)

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.DeleteRoutine(&postgres.DeleteRoutineDatabaseRequest{Routine: routine})

	return &RoutineDeleteDatabaseResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
