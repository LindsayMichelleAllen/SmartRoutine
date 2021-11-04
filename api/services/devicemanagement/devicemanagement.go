package devicemanagement

import (
	"api/interactors/databaseinteractor/devicedatabaseinteractor"
	"api/services/model"
	"errors"
)

type GetDeviceRequest struct {
	Id string
}

type GetUserDevicesRequest struct {
	UserId string
}

type GetRoutineDevicesRequest struct {
	RoutineId string
}

type DeviceCreateRequest struct {
	Name   string
	UserId string
}

type DeviceUpdateRequest struct {
	Name string
	Id   string
}

type DeviceDeleteRequest struct {
	Id string
}

type GetDeviceResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type GetDevicesResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetUserDevicesResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type GetRoutineDevicesResponse struct {
	Devices []*model.Device
	Message string
	Error   error
}

type DeviceCreateResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type DeviceUpdateResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type DeviceDeleteResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type DeviceService interface {
	GetDevice(request *GetDeviceRequest) *GetDeviceResponse
	GetDevices() *GetDevicesResponse
	GetUserDevices(request *GetUserDevicesRequest) *GetUserDevicesResponse
	GetRoutineDevices(request *GetRoutineDevicesRequest) *GetRoutineDevicesResponse
	CreateDevice(request *DeviceCreateRequest) *DeviceCreateResponse
	UpdateDevice(request *DeviceUpdateRequest) *DeviceUpdateResponse
	DeleteDevice(request *DeviceDeleteRequest) *DeviceDeleteResponse
}

type UnprotectedDeviceService struct {
	// intentionally left empty
}

func (d *UnprotectedDeviceService) GetDevice(request *GetDeviceRequest) *GetDeviceResponse {
	if request.Id == "" {
		return &GetDeviceResponse{
			Message: "Device Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.GetDevice(&devicedatabaseinteractor.GetDeviceInteractorRequest{Id: request.Id})
	return (*GetDeviceResponse)(resp)
}

func (d *UnprotectedDeviceService) GetDevices() *GetDevicesResponse {
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.GetDevices()
	return (*GetDevicesResponse)(resp)
}

func (d *UnprotectedDeviceService) GetUserDevices(request *GetUserDevicesRequest) *GetUserDevicesResponse {
	if request.UserId == "" {
		return &GetUserDevicesResponse{
			Message: "User Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.GetUserDevices(&devicedatabaseinteractor.GetUserDevicesInteractorRequest{UserId: request.UserId})
	return (*GetUserDevicesResponse)(resp)
}

func (d *UnprotectedDeviceService) GetRoutineDevices(request *GetRoutineDevicesRequest) *GetRoutineDevicesResponse {
	if request.RoutineId == "" {
		return &GetRoutineDevicesResponse{
			Message: "Routine Id not provided",
			Error:   errors.New("input field(s) missing"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.GetRoutineDevices(&devicedatabaseinteractor.GetRoutineDevicesInteractorRequest{RoutineId: request.RoutineId})
	return (*GetRoutineDevicesResponse)(resp)
}

func (d *UnprotectedDeviceService) CreateDevice(request *DeviceCreateRequest) *DeviceCreateResponse {
	if request.Name == "" || request.UserId == "" {
		return &DeviceCreateResponse{
			Message: "Device name or user id missing from request",
			Error:   errors.New("missing input field"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	// TODO randomly generate device id
	resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
		Id:     "987654321",
		Name:   request.Name,
		UserId: request.UserId,
	})

	if resp.Error != nil {
		return &DeviceCreateResponse{
			Message: resp.Message,
			Error:   resp.Error,
		}
	}

	device := &model.Device{}
	device.SetId(resp.Id)
	device.SetName(resp.Name)
	device.SetUserId(resp.UserId)

	return &DeviceCreateResponse{
		Device:  device,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceService) UpdateDevice(request *DeviceUpdateRequest) *DeviceUpdateResponse {
	if request.Name == "" || request.Id == "" {
		return &DeviceUpdateResponse{
			Message: "Device name or id missing from request",
			Error:   errors.New("missing input field"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.UpdateDevice(&devicedatabaseinteractor.UpdateDeviceRequest{
		Id:   request.Id,
		Name: request.Name,
	})
	return &DeviceUpdateResponse{
		Device:  nil,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceService) DeleteDevice(request *DeviceDeleteRequest) *DeviceDeleteResponse {
	if request.Id == "" {
		return &DeviceDeleteResponse{
			Message: "Device id missing from request",
			Error:   errors.New("missing input field"),
		}
	}
	dbInt := &devicedatabaseinteractor.UnprotectedDeviceDBInteractor{}
	resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{
		Id: request.Id,
	})
	return &DeviceDeleteResponse{
		Device:  nil,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
