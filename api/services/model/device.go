package model

import (
	"encoding/json"
)

/* The Device object is used to represent a controllable apparatus */
type Device struct {
	/* unique ID used to identify devices */
	id string
	/* name of device displayed to user */
	name string
	/* unique ID of user who added the device */
	userId string
}

func (d *Device) SetId(newId string) {
	d.id = newId
}

func (d *Device) SetName(newName string) {
	d.name = newName
}

func (d *Device) SetUserId(newUserId string) {
	d.userId = newUserId
}

func (d *Device) GetId() string {
	return d.id
}

func (d *Device) GetName() string {
	return d.name
}

func (d *Device) GetUserId() string {
	return d.userId
}

func (d *Device) GetJsonStruct() interface{} {
	return struct {
		Id     string
		Name   string
		UserID string
	}{
		Id:     d.GetId(),
		Name:   d.GetName(),
		UserID: d.GetUserId(),
	}
}

func (d *Device) GetJson() string {
	bytes, _ := json.Marshal(d.GetJsonStruct())
	return string(bytes)
}
