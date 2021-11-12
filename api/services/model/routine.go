package model

import "encoding/json"

type Routine struct {
	id            string
	name          string
	userId        string
	configuration []*Configuration
}

func (r *Routine) PopulateRoutine(id string, name string, userId string, configurations []*Configuration) {
	r.SetId(id)
	r.SetName(name)
	r.SetUserId(userId)
	r.SetConfiguration(configurations)
}

func (r *Routine) GetId() string {
	return r.id
}

func (r *Routine) GetName() string {
	return r.name
}

func (r *Routine) GetUserId() string {
	return r.userId
}

func (r *Routine) GetConfiguration() []*Configuration {
	return r.configuration
}

func (r *Routine) SetId(newId string) {
	r.id = newId
}

func (r *Routine) SetName(newName string) {
	r.name = newName
}

func (r *Routine) SetUserId(newUserId string) {
	r.userId = newUserId
}

func (r *Routine) SetConfiguration(newConfiguration []*Configuration) {
	r.configuration = newConfiguration
}

func (r *Routine) AddToConfiguration(config *Configuration) {
	r.configuration = append(r.configuration, config)
}

func (r *Routine) ClearConfiguration() {
	r.configuration = nil
}

func (r *Routine) GetJsonStruct() interface{} {
	configs := []interface{}{}
	for _, c := range r.configuration {
		configs = append(configs, c.GetJsonStruct())
	}

	return struct {
		Id            string
		Name          string
		UserId        string
		Configuration []interface{}
	}{
		Id:            r.id,
		Name:          r.name,
		UserId:        r.userId,
		Configuration: configs,
	}
}

func (r *Routine) GetJson() string {
	bytes, _ := json.Marshal(r.GetJsonStruct())
	return string(bytes)
}
