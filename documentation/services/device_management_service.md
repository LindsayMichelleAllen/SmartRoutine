## Device Management Service

The Device Management service is responsible for creating, modifying, and deleting devices. It is also responsible for maintaining a device's configuration.

---

### Device Interface

The device interface shall provide methods to:
* Create a new device
* Modify an existing device
* Delete an existing device

Below is an example of the minimum implementation of the device management service. All response objects currently contain the same struct attributes, but these will not be combined to a single response object type to allow for future updates and extensions.

```golang
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

func (d *UnprotectedDeviceService) CreateDevice(request *DeviceCreateRequest) *DeviceCreateResponse {...}

func (d *UnprotectedDeviceService) UpdateDevice(request *DeviceUpdateRequest) *DeviceUpdateResponse {...}

func (d *UnprotectedDeviceService) DeleteDevice(request *DeviceDeleteRequest) *DeviceDeleteResponse {...}
```

### Example Usage

```golang

import (
  dvcMngr "api/services/devicemanagement"
)

func main() {
    /* Create a new device */
    basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
    resp := basicDvcSrvc.CreateDevice(&dvcMngr.DeviceCreateRequest{
		Name:   "DeviceName",
		UserId: "UserId",
	})

	if resp.Error != nil {
		// handle error
	}

    device := resp.Device

    /* Modify an existing device */
    basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
	resp := basicDvcSrvc.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
		Name: "NewName",
		Id:   "DeviceId",
	})

	if resp.Error != nil {
        // handle error
    }

    updatedDev := resp.Device

    /* Delete an existing device */
    basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
	resp := basicDvcSrvc.DeleteDevice(&dvcMngr.DeviceDeleteRequest{
		Id: "DeviceId",
	})

	if resp.Error != nil {
		// handle error
	}

    removedDev := resp.Device
}
```

## Future updates