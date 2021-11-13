package routinemanagement

import (
	"api/interactors/databaseinteractor/routinedatabaseinteractor"
	"api/services/model"
	"errors"
)

type GetRoutineRequest struct {
	RoutineId string
}

type GetUserRoutinesRequest struct {
	UserId string
}

type GetDeviceRoutinesRequest struct {
	DeviceId string
}

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

type GetRoutineResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type GetRoutinesResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}

type GetUserRoutinesResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}

type GetDeviceRoutinesResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
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
	GetRoutine(request *GetRoutineRequest) *GetRoutineResponse
	GetRoutines() *GetRoutinesResponse
	GetDeviceRoutines(request *GetDeviceRoutinesRequest) *GetDeviceRoutinesResponse
	GetUserRoutines(request *GetUserRoutinesRequest) *GetUserRoutinesResponse
	CreateRoutine(request *RoutineCreateRequest) *RoutineCreateResponse
	UpdateRoutine(request *RoutineUpdateRequest) *RoutineUpdateResponse
	DeleteRoutine(request *RoutineDeleteRequest) *RoutineDeleteResponse
}

type UnprotectedRoutineService struct {
	// intentionally left empty
}

func (r *UnprotectedRoutineService) GetRoutine(request *GetRoutineRequest) *GetRoutineResponse {
	if request.RoutineId == "" {
		return &GetRoutineResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.GetRoutine(&routinedatabaseinteractor.GetRoutineInteractorRequest{RoutineId: request.RoutineId})
	return (*GetRoutineResponse)(resp)
}

func (r *UnprotectedRoutineService) GetRoutines() *GetRoutinesResponse {
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.GetRoutines()
	return (*GetRoutinesResponse)(resp)
}

func (r *UnprotectedRoutineService) GetDeviceRoutines(request *GetDeviceRoutinesRequest) *GetDeviceRoutinesResponse {
	if request.DeviceId == "" {
		return &GetDeviceRoutinesResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.GetDeviceRoutines(&routinedatabaseinteractor.GetDeviceRoutinesInteractorRequest{DeviceId: request.DeviceId})
	return (*GetDeviceRoutinesResponse)(resp)
}

func (r *UnprotectedRoutineService) GetUserRoutines(request *GetUserRoutinesRequest) *GetUserRoutinesResponse {
	if request.UserId == "" {
		return &GetUserRoutinesResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.GetUserRoutines(&routinedatabaseinteractor.GetUserRoutinesInteractorRequest{UserId: request.UserId})
	return (*GetUserRoutinesResponse)(resp)
}

func (r *UnprotectedRoutineService) CreateRoutine(request *RoutineCreateRequest) *RoutineCreateResponse {
	if request.Name == "" || request.UserId == "" {
		return &RoutineCreateResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	dbInt := &routinedatabaseinteractor.UnprotectedRoutineDBInteractor{}
	resp := dbInt.CreateRoutine(&routinedatabaseinteractor.RoutineCreateInteractorRequest{
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
	resp := dbInt.UpdateRoutine(&routinedatabaseinteractor.RoutineUpdateInteractorRequest{
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
	resp := dbInt.DeleteRoutine(&routinedatabaseinteractor.RoutineDeleteInteractorRequest{
		Id: request.Id,
	})

	return &RoutineDeleteResponse{
		Id:      resp.Id,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
