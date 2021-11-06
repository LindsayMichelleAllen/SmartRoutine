package routinedatabaseinteractor

import (
	"api/postgres"
	"api/services/model"
	"errors"
)

type GetRoutineInteractorRequest struct {
	RoutineId string
}
type GetUserRoutinesInteractorRequest struct {
	UserId string
}
type GetDeviceRoutinesInteractorRequest struct {
	DeviceId string
}

type RoutineCreateInteractorRequest struct {
	Id     string
	Name   string
	UserId string
}

type RoutineUpdateInteractorRequest struct {
	Id   string
	Name string
}

type RoutineDeleteInteractorRequest struct {
	Id string
}

type GetRoutineInteractorResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}
type GetRoutinesInteractorResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}
type GetUserRoutinesInteractorResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}
type GetDeviceRoutinesInteractorResponse struct {
	Routines []*model.Routine
	Message  string
	Error    error
}

type RoutineCreateInteractorResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineUpdateInteractorResponse struct {
	Routine *model.Routine
	Message string
	Error   error
}

type RoutineDeleteInteractorResponse struct {
	Id      string
	Message string
	Error   error
}

type RoutineDBInteractor interface {
	GetRoutine(request *GetRoutineInteractorRequest) *GetRoutineInteractorResponse
	GetRoutines() *GetRoutinesInteractorResponse
	GetDeviceRoutines(request *GetDeviceRoutinesInteractorRequest) *GetDeviceRoutinesInteractorResponse
	GetUserRoutines(request *GetUserRoutinesInteractorRequest) *GetUserRoutinesInteractorResponse
	CreateRoutine(request *RoutineCreateInteractorRequest) *RoutineCreateInteractorResponse
	UpdateRoutine(request *RoutineUpdateInteractorRequest) *RoutineUpdateInteractorResponse
	DeleteRoutine(request *RoutineDeleteInteractorRequest) *RoutineDeleteInteractorResponse
}

type UnprotectedRoutineDBInteractor struct {
	// intentionally left empty
}

func (r *UnprotectedRoutineDBInteractor) GetRoutine(request *GetRoutineInteractorRequest) *GetRoutineInteractorResponse {
	if request.RoutineId == "" {
		return &GetRoutineInteractorResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedRoutineDB{}
	resp := db.GetRoutine(&postgres.GetRoutineDatabaseRequest{RoutineId: request.RoutineId})
	return (*GetRoutineInteractorResponse)(resp)
}
func (r *UnprotectedRoutineDBInteractor) GetRoutines() *GetRoutinesInteractorResponse {
	db := &postgres.UnprotectedRoutineDB{}
	resp := db.GetRoutines()
	return (*GetRoutinesInteractorResponse)(resp)
}
func (r *UnprotectedRoutineDBInteractor) GetDeviceRoutines(request *GetDeviceRoutinesInteractorRequest) *GetDeviceRoutinesInteractorResponse {
	if request.DeviceId == "" {
		return &GetDeviceRoutinesInteractorResponse{
			Message: "Device Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedRoutineDB{}
	resp := db.GetDeviceRoutines(&postgres.GetDeviceRoutinesDatabaseRequest{DeviceId: request.DeviceId})
	return (*GetDeviceRoutinesInteractorResponse)(resp)
}
func (r *UnprotectedRoutineDBInteractor) GetUserRoutines(request *GetUserRoutinesInteractorRequest) *GetUserRoutinesInteractorResponse {
	if request.UserId == "" {
		return &GetUserRoutinesInteractorResponse{
			Message: "User Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedRoutineDB{}
	resp := db.GetUserRoutines(&postgres.GetUserRoutinesDatabaseRequest{UserId: request.UserId})
	return (*GetUserRoutinesInteractorResponse)(resp)
}

func (r *UnprotectedRoutineDBInteractor) CreateRoutine(request *RoutineCreateInteractorRequest) *RoutineCreateInteractorResponse {
	if request.Id == "" || request.Name == "" || request.UserId == "" {
		return &RoutineCreateInteractorResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	routine := &model.Routine{}
	routine.PopulateRoutine(request.Id, request.Name, request.UserId, make([]*model.Configuration, 0))

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.CreateRoutine(&postgres.CreateRoutineDatabaseRequest{Routine: routine})

	return &RoutineCreateInteractorResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineDBInteractor) UpdateRoutine(request *RoutineUpdateInteractorRequest) *RoutineUpdateInteractorResponse {
	if request.Id == "" || request.Name == "" {
		return &RoutineUpdateInteractorResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	routine := &model.Routine{}
	routine.PopulateRoutine(request.Id, request.Name, "", make([]*model.Configuration, 0))

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.UpdateRoutine(&postgres.UpdateRoutineDatabaseRequest{Routine: routine})

	return &RoutineUpdateInteractorResponse{
		Routine: resp.Routine,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (r *UnprotectedRoutineDBInteractor) DeleteRoutine(request *RoutineDeleteInteractorRequest) *RoutineDeleteInteractorResponse {
	if request.Id == "" {
		return &RoutineDeleteInteractorResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}

	db := &postgres.UnprotectedRoutineDB{}
	resp := db.DeleteRoutine(&postgres.DeleteRoutineDatabaseRequest{Id: request.Id})

	return &RoutineDeleteInteractorResponse{
		Id:      resp.Id,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
