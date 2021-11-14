package model

type Routine struct {
	id            string
	name          string
	userId        string
	basealarm     string
	configuration []*Configuration
}

func (r *Routine) PopulateRoutine(id string, name string, userId string, basealarm string, configurations []*Configuration) {
	r.SetId(id)
	r.SetName(name)
	r.SetUserId(userId)
	r.SetBaseAlarm(basealarm)
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

func (r *Routine) GetBaseAlarm() string {
	return r.basealarm
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

func (r *Routine) SetBaseAlarm(newBaseAlarm string) {
	r.basealarm = newBaseAlarm
}

func (r *Routine) SetConfiguration(newConfiguration []*Configuration) {
	r.configuration = newConfiguration
}

func (r *Routine) AddToConfiguration(config *Configuration) {
	r.configuration = append(r.configuration, config)
}

func (r *Routine) ClearConfigurations() {
	r.configuration = nil
}
