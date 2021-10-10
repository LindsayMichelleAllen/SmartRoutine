package devicemanagement

import (
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
	return &DeviceCreateResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (d *UnprotectedDeviceService) UpdateDevice(request *DeviceUpdateRequest) *DeviceUpdateResponse {
	return &DeviceUpdateResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (d *UnprotectedDeviceService) DeleteDevice(request *DeviceDeleteRequest) *DeviceDeleteResponse {
	return &DeviceDeleteResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
