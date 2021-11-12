package devicedatabaseinteractor

import (
	"api/postgres"
	"api/services/model"
	"errors"
)

type GetDeviceInteractorRequest struct {
	Id string
}

type GetUserDevicesInteractorRequest struct {
	UserId string
}

type GetRoutineDevicesInteractorRequest struct {
	RoutineId string
}

type CreateDeviceRequest struct {
	Id     string
	Name   string
	UserId string
}

type UpdateDeviceRequest struct {
	Id   string
	Name string
}

type DeleteDeviceRequest struct {
	Id string
}

type GetDeviceInteractorResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type GetDevicesInteractorResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetUserDevicesInteractorResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetRoutineDevicesInteractorResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type CreateDeviceResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type UpdateDeviceRespose struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type DeleteDeviceResponse struct {
	Id      string
	Name    string
	UserId  string
	Message string
	Error   error
}

type DeviceDBInteractor interface {
	GetDevice(request *GetDeviceInteractorRequest) *GetDeviceInteractorResponse
	GetDevices() *GetDevicesInteractorResponse
	GetUserDevices(request *GetUserDevicesInteractorRequest) *GetUserDevicesInteractorResponse
	GetRoutineDevices(request *GetRoutineDevicesInteractorRequest) *GetRoutineDevicesInteractorResponse
	CreateDevice(request *CreateDeviceRequest) *CreateDeviceResponse
	UpdateDevice(request *UpdateDeviceRequest) *UpdateDeviceRespose
	DeleteDevice(request *DeleteDeviceRequest) *DeleteDeviceResponse
}

type UnprotectedDeviceDBInteractor struct {
	// intentionally left empty
}

func (d *UnprotectedDeviceDBInteractor) GetDevice(request *GetDeviceInteractorRequest) *GetDeviceInteractorResponse {
	if request.Id == "" {
		return &GetDeviceInteractorResponse{
			Message: "Device Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedDeviceDB{}
	resp := db.GetDevice(&postgres.GetDeviceDatabaseRequest{Id: request.Id})

	return (*GetDeviceInteractorResponse)(resp)
}

func (d *UnprotectedDeviceDBInteractor) GetDevices() *GetDevicesInteractorResponse {
	db := &postgres.UnprotectedDeviceDB{}
	resp := db.GetDevices()
	return (*GetDevicesInteractorResponse)(resp)
}

func (d *UnprotectedDeviceDBInteractor) GetUserDevices(request *GetUserDevicesInteractorRequest) *GetUserDevicesInteractorResponse {
	if request.UserId == "" {
		return &GetUserDevicesInteractorResponse{
			Message: "User Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db := &postgres.UnprotectedDeviceDB{}
	resp := db.GetUserDevices(&postgres.GetUserDevicesDatabaseRequest{UserId: request.UserId})
	return (*GetUserDevicesInteractorResponse)(resp)
}

func (d *UnprotectedDeviceDBInteractor) GetRoutineDevices(request *GetRoutineDevicesInteractorRequest) *GetRoutineDevicesInteractorResponse {
	if request.RoutineId == "" {
		return &GetRoutineDevicesInteractorResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}

	db := &postgres.UnprotectedDeviceDB{}
	resp := db.GetRoutineDevices(&postgres.GetRoutineDevicesDatabaseRequest{RoutineId: request.RoutineId})
	return (*GetRoutineDevicesInteractorResponse)(resp)
}

func (d *UnprotectedDeviceDBInteractor) CreateDevice(request *CreateDeviceRequest) *CreateDeviceResponse {
	db := &postgres.UnprotectedDeviceDB{}
	resp := db.CreateDevice(&postgres.CreateDeviceDatabaseRequest{
		Id:     request.Id,
		Name:   request.Name,
		UserId: request.UserId,
	})

	if resp.Error != nil {
		return &CreateDeviceResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return &CreateDeviceResponse{
		Id:      resp.Id,
		Name:    resp.Name,
		UserId:  resp.UserId,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceDBInteractor) UpdateDevice(request *UpdateDeviceRequest) *UpdateDeviceRespose {
	db := &postgres.UnprotectedDeviceDB{}
	resp := db.UpdateDevice(&postgres.UpdateDeviceDatabaseRequest{
		Id:   request.Id,
		Name: request.Name,
	})

	if resp.Error != nil {
		return &UpdateDeviceRespose{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return &UpdateDeviceRespose{
		Id:      resp.Id,
		Name:    resp.Name,
		UserId:  resp.UserId,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceDBInteractor) DeleteDevice(request *DeleteDeviceRequest) *DeleteDeviceResponse {
	db := &postgres.UnprotectedDeviceDB{}
	resp := db.DeleteDevice(&postgres.DeleteDeviceDatabaseRequest{
		Id: request.Id,
	})

	if resp.Error != nil {
		return &DeleteDeviceResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	return &DeleteDeviceResponse{
		Id:      resp.Id,
		Name:    resp.Name,
		UserId:  resp.UserId,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
