package model

type Configuration struct {
	id     string
	device *Device
}

func (c *Configuration) GetId() string {
	return c.id
}

func (c *Configuration) GetDevice() *Device {
	return c.device
}

func (c *Configuration) SetId(newId string) {
	c.id = newId
}

func (c *Configuration) SetDevice(newDevice *Device) {
	c.device = newDevice
}
