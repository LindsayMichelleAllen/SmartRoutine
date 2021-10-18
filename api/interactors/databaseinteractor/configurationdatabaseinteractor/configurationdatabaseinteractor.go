package configurationdatabaseinteractor

import (
	"api/postgres"
	"api/services/model"
)

type CreateConfigurationDBInteractorRequest struct {
	ConfigId string
	Offset   *int
	DeviceId string
}
type UpdateConfigurationDBInteractorRequest struct {
	ConfigId string
	Offset   *int
	DeviceId string
}
type DeleteConfigurationDBInteractorRequest struct {
	ConfigId string
}
type CreateConfigurationDBInteractorResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type UpdateConfigurationDBInteractorResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}
type DeleteConfigurationDBInteractorResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}

type ConfigurationDBInteractor interface {
	CreateConfiguration(request *CreateConfigurationDBInteractorRequest) *CreateConfigurationDBInteractorResponse
	UpdateConfiguration(request *UpdateConfigurationDBInteractorRequest) *UpdateConfigurationDBInteractorResponse
	DeleteConfiguration(request *DeleteConfigurationDBInteractorRequest) *DeleteConfigurationDBInteractorResponse
}

type UnprotectedConfigurationDBInteractor struct {
	// intentionally left empty
}

func (c *UnprotectedConfigurationDBInteractor) CreateConfiguration(request *CreateConfigurationDBInteractorRequest) *CreateConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.CreateConfiguration(&postgres.CreateConfigurationDatabaseRequest{})
	return &CreateConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationDBInteractor) UpdateConfiguration(request *UpdateConfigurationDBInteractorRequest) *UpdateConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.UpdateConfiguration(&postgres.UpdateConfigurationDatabaseRequest{})
	return &UpdateConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationDBInteractor) DeleteConfiguration(request *DeleteConfigurationDBInteractorRequest) *DeleteConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.DeleteConfiguration(&postgres.DeleteConfigurationDatabaseRequest{})
	return &DeleteConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}
