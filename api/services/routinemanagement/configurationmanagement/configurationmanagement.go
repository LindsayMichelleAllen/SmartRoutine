package configurationmanagement

import (
	"api/interactors/databaseinteractor/configurationdatabaseinteractor"
	"api/services/model"
	"errors"
)

type GetConfigurationRequest struct {
	ConfigId string
}

type GetUserConfigurationsRequest struct {
	UserId string
}

type GetDeviceConfigurationsRequest struct {
	DeviceId string
}

type GetRoutineConfigurationsRequest struct {
	RoutineId string
}

type CreateConfigurationRequest struct {
	Offset    *int
	DeviceId  string
	RoutineId string
}
type UpdateConfigurationRequest struct {
	ConfigId  string
	Offset    *int
	DeviceId  string
	RoutineId string
}
type DeleteConfigurationRequest struct {
	ConfigId string
}

type GetConfigurationResponse struct {
	Configuration *model.Configuration
	Message       string
	Error         error
}

type GetConfigurationsResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetUserConfigurationsResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetDeviceConfigurationsResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
}

type GetRoutineConfigurationsResponse struct {
	Configurations []*model.Configuration
	Message        string
	Error          error
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
	GetConfiguration(request *GetConfigurationRequest) *GetConfigurationResponse
	GetConfigurations() *GetConfigurationsResponse
	GetUserConfigurations(request *GetUserConfigurationsRequest) *GetUserConfigurationsResponse
	GetDeviceConfigurations(request *GetDeviceConfigurationsRequest) *GetDeviceConfigurationsResponse
	GetRoutineConfigurations(request *GetRoutineConfigurationsRequest) *GetRoutineConfigurationsResponse
	CreateConfiguration(request *CreateConfigurationRequest) *CreateConfigurationResponse
	UpdateConfiguration(request *UpdateConfigurationRequest) *UpdateConfigurationResponse
	DeleteConfiguration(request *DeleteConfigurationRequest) *DeleteConfigurationResponse
}

type UnprotectedConfigurationService struct {
	// intentionally left empty
}

func (c *UnprotectedConfigurationService) GetConfiguration(request *GetConfigurationRequest) *GetConfigurationResponse {
	if request.ConfigId == "" {
		return &GetConfigurationResponse{
			Configuration: nil,
			Message:       "Config Id not provided",
			Error:         errors.New("input field(s) missing"),
		}
	}
	dbInt := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := dbInt.GetConfiguration(&configurationdatabaseinteractor.GetConfigurationInteractorRequest{
		ConfigId: request.ConfigId,
	})
	return (*GetConfigurationResponse)(resp)
}

func (c *UnprotectedConfigurationService) GetConfigurations() *GetConfigurationsResponse {
	dbInt := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := dbInt.GetConfigurations()
	return (*GetConfigurationsResponse)(resp)
}

func (c *UnprotectedConfigurationService) GetUserConfigurations(request *GetUserConfigurationsRequest) *GetUserConfigurationsResponse {
	if request.UserId == "" {
		return &GetUserConfigurationsResponse{
			Configurations: nil,
			Message:        "User Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	dbInt := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := dbInt.GetUserConfigurations(&configurationdatabaseinteractor.GetUserConfiguraitonsInteractorRequest{
		UserId: request.UserId,
	})
	return (*GetUserConfigurationsResponse)(resp)
}

func (c *UnprotectedConfigurationService) GetDeviceConfigurations(request *GetDeviceConfigurationsRequest) *GetDeviceConfigurationsResponse {
	if request.DeviceId == "" {
		return &GetDeviceConfigurationsResponse{
			Configurations: nil,
			Message:        "Device Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	dbInt := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := dbInt.GetDeviceConfigurations(&configurationdatabaseinteractor.GetDeviceConfigurationsInteractorRequest{
		DeviceId: request.DeviceId,
	})
	return (*GetDeviceConfigurationsResponse)(resp)
}

func (c *UnprotectedConfigurationService) GetRoutineConfigurations(request *GetRoutineConfigurationsRequest) *GetRoutineConfigurationsResponse {
	if request.RoutineId == "" {
		return &GetRoutineConfigurationsResponse{
			Configurations: nil,
			Message:        "Routine Id not provided",
			Error:          errors.New("input field(s) missing"),
		}
	}
	dbInt := configurationdatabaseinteractor.UnprotectedConfigurationDBInteractor{}
	resp := dbInt.GetRoutineConfigurations(&configurationdatabaseinteractor.GetRoutineConfigurationsInteractorRequest{
		RoutineId: request.RoutineId,
	})
	return (*GetRoutineConfigurationsResponse)(resp)
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
		ConfigId:  "ConfigID",
		Offset:    request.Offset,
		DeviceId:  request.DeviceId,
		RoutineId: request.RoutineId,
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
