package routinemanagement

import (
	"api/interactors/databaseinteractor/routinedatabaseinteractor"
	"api/services/model"
	"errors"
)

type RoutineCreateRequest struct {
	UserId string
	Name   string
}

type RoutineUpdateRequest struct {
	Id   string
	Name string
}

type RoutineDeleteRequest struct {
	Id string
}

type RoutineCreateResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineUpdateResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineDeleteResponse struct {
	Id      string
	Message string
	Error   error
}

type RoutineService interface {
	CreateRoutine(request *RoutineCreateRequest) *RoutineCreateResponse
	UpdateRoutine(request *RoutineUpdateRequest) *RoutineUpdateResponse
	DeleteRoutine(request *RoutineDeleteRequest) *RoutineDeleteResponse
}

type UnprotectedRoutineService struct {
	// intentionally left empty
}

func (r *UnprotectedRoutineService) CreateRoutine(request *RoutineCreateRequest) *RoutineCreateResponse {
	if request.Name == "" || request.UserId == "" {
		return &RoutineCreateResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateDatabaseRequest{
		Id:     "976431852", // TODO: generate routine id
		Name:   request.Name,
		UserId: request.UserId,
	})

	return &RoutineCreateResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineService) UpdateRoutine(request *RoutineUpdateRequest) *RoutineUpdateResponse {
	if request.Id == "" || request.Name == "" {
		return &RoutineUpdateResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateDatabaseRequest{
		Id:   request.Id,
		Name: request.Name,
	})

	return &RoutineUpdateResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineService) DeleteRoutine(request *RoutineDeleteRequest) *RoutineDeleteResponse {
	if request.Id == "" {
		return &RoutineDeleteResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteDatabaseRequest{
		Id: request.Id,
	})

	return &RoutineDeleteResponse{
		Id:      resp.Id,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
