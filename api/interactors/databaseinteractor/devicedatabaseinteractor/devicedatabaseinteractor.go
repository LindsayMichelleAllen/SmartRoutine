package devicedatabaseinteractor

import (
	"api/postgres"
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
	CreateDevice(*CreateDeviceRequest) *CreateDeviceResponse
	UpdateDevice(*UpdateDeviceRequest) *UpdateDeviceRespose
	DeleteDevice(*DeleteDeviceRequest) *DeleteDeviceResponse
}

type BasicDeviceDBInteractor struct {
	// intentionally left empty
}

func (d *BasicDeviceDBInteractor) CreateDevice(request *CreateDeviceRequest) *CreateDeviceResponse {
	dbInt := &postgres.DeviceDB{}
	resp := dbInt.CreateDevice(&postgres.CreateDeviceDatabaseRequest{
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
