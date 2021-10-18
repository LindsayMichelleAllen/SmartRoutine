package configurationmanagement

import (
	"api/interactors/databaseinteractor/configurationdatabaseinteractor"
	"api/services/model"
	"errors"
)

type CreateConfigurationRequest struct {
	Offset   *int
	DeviceId string
}
type UpdateConfigurationRequest struct {
	ConfigId string
	Offset   *int
	DeviceId string
}
type DeleteConfigurationRequest struct {
	ConfigId string
}
type CreateConfigurationResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type UpdateConfigurationResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type DeleteConfigurationResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}

type ConfigurationService interface {
	CreateConfiguration(request *CreateConfigurationRequest) *CreateConfigurationResponse
	UpdateConfiguration(request *UpdateConfigurationRequest) *UpdateConfigurationResponse
	DeleteConfiguration(request *DeleteConfigurationRequest) *DeleteConfigurationResponse
}

type UnprotectedConfigurationService struct {
	// intentionally left empty
}

func (c *UnprotectedConfigurationService) CreateConfiguration(request *CreateConfigurationRequest) *CreateConfigurationResponse {
	if request.Offset == nil || request.DeviceId == "" {
		return &CreateConfigurationResponse{
			Message: "Input field(s) missing",
			Error:   errors.New("input field(s) missing"),
		}
	}
	db := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := db.CreateConfiguration(&configurationdatabaseinteractor.CreateConfigurationDBInteractorRequest{
		ConfigId: "ConfigID",
		Offset:   request.Offset,
		DeviceId: request.DeviceId,
	})
	return &CreateConfigurationResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationService) UpdateConfiguration(request *UpdateConfigurationRequest) *UpdateConfigurationResponse {
	db := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := db.UpdateConfiguration(&configurationdatabaseinteractor.UpdateConfigurationDBInteractorRequest{
		ConfigId: request.ConfigId,
		Offset:   request.Offset,
		DeviceId: request.DeviceId,
	})
	return &UpdateConfigurationResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationService) DeleteConfiguration(request *DeleteConfigurationRequest) *DeleteConfigurationResponse {
	db := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := db.DeleteConfiguration(&configurationdatabaseinteractor.DeleteConfigurationDBInteractorRequest{
		ConfigId: request.ConfigId,
	})
	return &DeleteConfigurationResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}
