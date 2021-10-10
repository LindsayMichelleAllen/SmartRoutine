package model

type UserProfile struct {
	/* username used for login */
	userName string
	/* name displayed to user */
	name string
	/* unique id of user */
	id string
}

/* SetUsername overwrites the private username field. */
func (u *UserProfile) SetUsername(newUsername string) {
	u.userName = newUsername
}

/* GetUsername returns the current value of the private username field. */
func (u *UserProfile) GetUsername() string {
	return u.userName
}

/* SetName overwrites the private name field. */
func (u *UserProfile) SetName(newName string) {
	u.name = newName
}

/* GetName returns the current value of the private name field. */
func (u *UserProfile) GetName() string {
	return u.name
}

/* SetId overwrites the private id field. */
func (u *UserProfile) SetId(newId string) {
	u.id = newId
}

/* GetId returns the current value of the private id field. */
func (u *UserProfile) GetId() string {
	return u.id
}
