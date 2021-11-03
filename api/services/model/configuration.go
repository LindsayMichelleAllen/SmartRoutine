package model

type Configuration struct {
	id        string
	routineId string
	offset    int
	device    *Device
}

func (c *Configuration) GetId() string {
	return c.id
}

func (c *Configuration) GetDevice() *Device {
	return c.device
}

func (c *Configuration) GetOffset() int {
	return c.offset
}

func (c *Configuration) GetRoutineId() string {
	return c.routineId
}

func (c *Configuration) SetId(newId string) {
	c.id = newId
}

func (c *Configuration) SetDevice(newDevice *Device) {
	c.device = newDevice
}

func (c *Configuration) SetOffset(newOffset int) {
	c.offset = newOffset
}

func (c *Configuration) SetRoutineId(newRoutineId string) {
	c.routineId = newRoutineId
}
