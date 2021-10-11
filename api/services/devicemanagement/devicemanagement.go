package devicemanagement

import (
	"api/interactors/databaseinteractor/devicedatabaseinteractor"
	"api/services/model"
	"errors"
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
	if request.Name == "" || request.UserId == "" {
		return &DeviceCreateResponse{
			Message: "Device name or user id missing from request",
			Error:   errors.New("missing input field"),
		}
	}
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
	// TODO randomly generate device id
	resp := dbInt.CreateDevice(&devicedatabaseinteractor.CreateDeviceRequest{
		Id:     "976431852",
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
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
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
	dbInt := &devicedatabaseinteractor.BasicDeviceDBInteractor{}
	resp := dbInt.DeleteDevice(&devicedatabaseinteractor.DeleteDeviceRequest{
		Id: request.Id,
	})
	return &DeviceDeleteResponse{
		Device:  nil,
		Message: resp.Message,
		Error:   resp.Error,
	}
}
