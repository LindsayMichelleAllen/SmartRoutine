package devicedatabaseinteractor

import (
	"api/services/model"
	"errors"
)

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

type CreateDeviceResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type UpdateDeviceRespose struct {
	Device  *model.Device
	Message string
	Error   error
}

type DeleteDeviceResponse struct {
	Device  *model.Device
	Message string
	Error   error
}

type DeviceDBInteractor interface {
	CreateDevice(*CreateDeviceRequest) *CreateDeviceResponse
	UpdateDevice(*UpdateDeviceRequest) *UpdateDeviceRespose
	DeleteDevice(*DeleteDeviceRequest) *DeleteDeviceResponse
}

type BasicDeviceDBInteractor struct {
	// intentionally left empty
}

func (d *BasicDeviceDBInteractor) CreateDevice(request *CreateDeviceRequest) *CreateDeviceResponse {
	return &CreateDeviceResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (d *BasicDeviceDBInteractor) UpdateDevice(request *UpdateDeviceRequest) *UpdateDeviceRespose {
	return &UpdateDeviceRespose{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}

func (d *BasicDeviceDBInteractor) DeleteDevice(request *DeleteDeviceRequest) *DeleteDeviceResponse {
	return &DeleteDeviceResponse{
		Message: "Not Yet Implemented",
		Error:   errors.New("not yet implemented"),
	}
}
