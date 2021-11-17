package model

type UserProfile struct {
	/* username used for login */
	userName string
	/* name displayed to user */
	name string
	/* status of user authenication */
	isAuth bool
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
func (u *UserProfile) SetAuthorizationStatus(stat bool) {
	u.isAuth = stat
}

/* GetId returns the current value of the private id field. */
func (u *UserProfile) GetAuthorizationStatus() bool {
	return u.isAuth
}
