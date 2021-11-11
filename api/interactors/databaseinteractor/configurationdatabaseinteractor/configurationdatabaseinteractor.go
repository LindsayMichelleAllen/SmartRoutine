package configurationdatabaseinteractor

import (
	"api/postgres"
	"api/services/model"
	"errors"
)

type GetConfigurationInteractorRequest struct {
	ConfigId string
}

type GetUserConfiguraitonsInteractorRequest struct {
	UserId string
}

type GetDeviceConfigurationsInteractorRequest struct {
	DeviceId string
}

type GetRoutineConfigurationsInteractorRequest struct {
	RoutineId string
}

type CreateConfigurationDBInteractorRequest struct {
	ConfigId  string
	Offset    *int
	DeviceId  string
	RoutineId string
}
type UpdateConfigurationDBInteractorRequest struct {
	ConfigId  string
	Offset    *int
	DeviceId  string
	RoutineId string
}
type DeleteConfigurationDBInteractorRequest struct {
	ConfigId string
}

type GetConfigurationInteractorResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}

type GetConfigurationsInteractorResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetUserConfiguraitonsInteractorResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetDeviceConfigurationsInteractorResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetRoutineConfigurationsInteractorResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
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
	GetConfiguration(request *GetConfigurationInteractorRequest) *GetConfigurationInteractorResponse
	GetConfigurations() *GetConfigurationsInteractorResponse
	GetUserConfigurations(request *GetUserConfiguraitonsInteractorRequest) *GetUserConfiguraitonsInteractorResponse
	GetDeviceConfigurations(request *GetDeviceConfigurationsInteractorRequest) *GetDeviceConfigurationsInteractorResponse
	GetRoutineConfigurations(request *GetRoutineConfigurationsInteractorRequest) *GetRoutineConfigurationsInteractorResponse
	CreateConfiguration(request *CreateConfigurationDBInteractorRequest) *CreateConfigurationDBInteractorResponse
	UpdateConfiguration(request *UpdateConfigurationDBInteractorRequest) *UpdateConfigurationDBInteractorResponse
	DeleteConfiguration(request *DeleteConfigurationDBInteractorRequest) *DeleteConfigurationDBInteractorResponse
}

type UnprotectedConfigurationDBInteractor struct {
	// intentionally left empty
}

func (c *UnprotectedConfigurationDBInteractor) GetConfiguration(request *GetConfigurationInteractorRequest) *GetConfigurationInteractorResponse {
	if request.ConfigId == "" {
		return &GetConfigurationInteractorResponse{
			Configuration: nil,
			Message:       "Config Id not provided",
			Error:         errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.GetConfiguration(&postgres.GetConfigurationDatabaseRequest{
		ConfigId: request.ConfigId,
	})
	return (*GetConfigurationInteractorResponse)(resp)
}

func (c *UnprotectedConfigurationDBInteractor) GetConfigurations() *GetConfigurationsInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.GetConfigurations()
	return (*GetConfigurationsInteractorResponse)(resp)
}

func (c *UnprotectedConfigurationDBInteractor) GetUserConfigurations(request *GetUserConfiguraitonsInteractorRequest) *GetUserConfiguraitonsInteractorResponse {
	if request.UserId == "" {
		return &GetUserConfiguraitonsInteractorResponse{
			Configurations: nil,
			Message:        "Config Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.GetUserConfigurations(&postgres.GetUserConfiguraitonsDatabaseRequest{
		UserId: request.UserId,
	})
	return (*GetUserConfiguraitonsInteractorResponse)(resp)
}

func (c *UnprotectedConfigurationDBInteractor) GetDeviceConfigurations(request *GetDeviceConfigurationsInteractorRequest) *GetDeviceConfigurationsInteractorResponse {
	if request.DeviceId == "" {
		return &GetDeviceConfigurationsInteractorResponse{
			Configurations: nil,
			Message:        "Config Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.GetDeviceConfigurations(&postgres.GetDeviceConfigurationsDatabaseRequest{
		DeviceId: request.DeviceId,
	})
	return (*GetDeviceConfigurationsInteractorResponse)(resp)
}

func (c *UnprotectedConfigurationDBInteractor) GetRoutineConfigurations(request *GetRoutineConfigurationsInteractorRequest) *GetRoutineConfigurationsInteractorResponse {
	if request.RoutineId == "" {
		return &GetRoutineConfigurationsInteractorResponse{
			Configurations: nil,
			Message:        "Config Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.GetRoutineConfigurations(&postgres.GetRoutineConfigurationsDatabaseRequest{
		RoutineId: request.RoutineId,
	})
	return (*GetRoutineConfigurationsInteractorResponse)(resp)
}

func (c *UnprotectedConfigurationDBInteractor) CreateConfiguration(request *CreateConfigurationDBInteractorRequest) *CreateConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	config := &model.Configuration{}
	dev := &model.Device{}
	dev.SetId(request.DeviceId)
	config.SetId(request.ConfigId)
	config.SetOffset(*request.Offset)
	config.SetRoutineId(request.RoutineId)
	config.SetDevice(dev)
	resp := db.CreateConfiguration(&postgres.CreateConfigurationDatabaseRequest{Configuration: config})
	return &CreateConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationDBInteractor) UpdateConfiguration(request *UpdateConfigurationDBInteractorRequest) *UpdateConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	config := &model.Configuration{}
	config.SetId(request.ConfigId)
	config.SetOffset(*request.Offset)
	resp := db.UpdateConfiguration(&postgres.UpdateConfigurationDatabaseRequest{
		Configuration: config,
	})
	return &UpdateConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}

func (c *UnprotectedConfigurationDBInteractor) DeleteConfiguration(request *DeleteConfigurationDBInteractorRequest) *DeleteConfigurationDBInteractorResponse {
	db := &postgres.UnprotectedConfigurationDB{}
	resp := db.DeleteConfiguration(&postgres.DeleteConfigurationDatabaseRequest{
		Id: request.ConfigId,
	})
	return &DeleteConfigurationDBInteractorResponse{
		Configuration: resp.Configuration,
		Message:       resp.Message,
		Error:         resp.Error,
	}
}
