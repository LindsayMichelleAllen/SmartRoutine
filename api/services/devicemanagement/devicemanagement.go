package devicemanagement

import (
	"api/interactors/databaseinteractor/devicedatabaseinteractor"
	"api/services/model"
)

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
	CreateDevice(*DeviceCreateRequest) *DeviceCreateResponse
	UpdateDevice(*DeviceUpdateRequest) *DeviceUpdateResponse
	DeleteDevice(*DeviceDeleteRequest) *DeviceDeleteResponse
}

type UnprotectedDeviceService struct {
	// intentionally left empty
}

func (d *UnprotectedDeviceService) CreateDevice(request *DeviceCreateRequest) *DeviceCreateResponse {
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
	// TODO randomly generate device id
	resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
		Id:     "976431852",
		Name:   request.Name,
		UserId: request.UserId,
	})
	return &DeviceCreateResponse{
		Device:  resp.Device,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceService) UpdateDevice(request *DeviceUpdateRequest) *DeviceUpdateResponse {
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
	resp := dbInt.UpdateDevice(&devicedatabaseinteractor.UpdateDeviceRequest{
		Id:   request.Id,
		Name: request.Name,
	})
	return &DeviceUpdateResponse{
		Device:  resp.Device,
		Message: resp.Message,
		Error:   resp.Error,
	}
}

func (d *UnprotectedDeviceService) DeleteDevice(request *DeviceDeleteRequest) *DeviceDeleteResponse {
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
	resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{
		Id: request.Id,
	})
	return &DeviceDeleteResponse{
		Device:  resp.Device,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
