package models

type User struct {
	/* username used for login */
	userName string
	/* name displayed to user */
	name     string
	/* unique id of user */
	id       string
}

/*
SetUsername overwrites the private username field.
*/
func(u *User) SetUsername(newUsername string) {
	u.userName = newUsername
}

/*
GetUsername returns the current value of the private username field.
*/
func(u *User) GetUsername() string {
	return u.userName
}

/*
SetName overwrites the private name field.
*/
func(u *User) SetName(newName string) {
	u.name = newName
}

/*
GetName returns the current value of the private name field.
*/
func(u *User) GetName() string {
	return u.name
}

/*
SetId overwrites the private id field.
*/
func(u *User) SetId(newId string) {
	u.id = newId
}

/*
GetId returns the current value of the private id field.
*/
func(u *User) GetId() string{
	return u.id
}