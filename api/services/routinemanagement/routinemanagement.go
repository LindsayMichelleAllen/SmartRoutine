package routinemanagement

import (
	"api/services/model"
	"errors"
)

type RoutineCreateRequest struct {
	UserId string
	Name   string
}

type RoutineUpdateRequest struct {
	Id            string
	Configuration *model.Configuration
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
	Routine *model.Routine
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
	return &RoutineCreateResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (r *UnprotectedRoutineService) UpdateRoutine(request *RoutineUpdateRequest) *RoutineUpdateResponse {
	if request.Id == "" || request.Configuration == nil {
		return &RoutineUpdateResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	return &RoutineUpdateResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (r *UnprotectedRoutineService) DeleteRoutine(request *RoutineDeleteRequest) *RoutineDeleteResponse {
	if request.Id == "" {
		return &RoutineDeleteResponse{
			Message: "Input field missing",
			Error:   errors.New("input field missing"),
		}
	}
	return &RoutineDeleteResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
