package model

type Routine struct {
	id            string
	userId        string
	configuration []*Configuration
}

func (r *Routine) GetId() string {
	return r.id
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

func (r *Routine) SetUserId(newUserId string) {
	r.userId = newUserId
}

func (r *Routine) SetConfiguration(newConfiguration []*Configuration) {
	r.configuration = newConfiguration
}
