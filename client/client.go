package main

import (
	"log"
	"net/http"
	"net/url"
)

/******** USER PROFILE *******/
func getUserProfile(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/user/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getUserProfiles() *http.Response {
	resp, err := http.Get("http://localhost:8080/users/")
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func login(username string, password string) *http.Response {
	data := url.Values{
		"username": {username},
		"password": {password},
	}
	resp, err := http.PostForm("http://localhost:8080/user/login/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func createUserProfile(name string, username string, password string) *http.Response {
	data := url.Values{
		"name":     {name},
		"username": {username},
		"password": {password},
	}
	resp, err := http.PostForm("http://localhost:8080/user/create/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func updateUserProfile(id string, username string, name string) *http.Response {
	data := url.Values{
		"id":       {id},
		"username": {username},
		"name":     {name},
	}
	resp, err := http.PostForm("http://localhost:8080/user/update/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func deleteUserProfile(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/user/delete/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

/******** DEVICES ***********/
func getDevice(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/device/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getDevices() *http.Response {
	resp, err := http.Get("http://localhost:8080/devices/")
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getUserDevices(userid string) *http.Response {
	data := url.Values{
		"userid": {userid},
	}
	resp, err := http.PostForm("http://localhost:8080/device/user/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getRoutineDevices(routineid string) *http.Response {
	data := url.Values{
		"routineid": {routineid},
	}
	resp, err := http.PostForm("http://localhost:8080/device/routine/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func createDevice(name string, userid string) *http.Response {
	data := url.Values{
		"name":   {name},
		"userid": {userid},
	}
	resp, err := http.PostForm("http://localhost:8080/device/create/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func updateDevice(name string, deviceid string) *http.Response {
	data := url.Values{
		"name":     {name},
		"deviceid": {deviceid},
	}
	resp, err := http.PostForm("http://localhost:8080/device/update/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func deleteDevice(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/device/delete/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

/******** ROUTINES **********/
func getRoutine(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/routine/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getRoutines() *http.Response {
	resp, err := http.Get("http://localhost:8080/routines/")

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getUserRoutines(userid string) *http.Response {
	data := url.Values{
		"userid": {userid},
	}
	resp, err := http.PostForm("http://localhost:8080/routines/user/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getDeviceRoutines(deviceid string) *http.Response {
	data := url.Values{
		"deviceid": {deviceid},
	}
	resp, err := http.PostForm("http://localhost:8080/routines/device/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func createRoutine(name string, userid string, basealarm string) *http.Response {
	data := url.Values{
		"name":      {name},
		"userid":    {userid},
		"basealarm": {basealarm},
	}
	resp, err := http.PostForm("http://localhost:8080/routine/create/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func updateRoutine(name string, routineid string, basealarm string) *http.Response {
	data := url.Values{
		"name":      {name},
		"routineid": {routineid},
		"basealarm": {basealarm},
	}
	resp, err := http.PostForm("http://localhost:8080/routine/update/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func deleteRoutine(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/routine/delete/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

/****** CONFIGURATIONS ******/
func getConfiguration(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/configuration/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getConfigurations() *http.Response {
	resp, err := http.Get("http://localhost:8080/configurations/")

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getDeviceConfgiurations(deviceid string) *http.Response {
	data := url.Values{
		"deviceid": {deviceid},
	}
	resp, err := http.PostForm("http://localhost:8080/configurations/device/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getUserConfgiurations(userid string) *http.Response {
	data := url.Values{
		"userid": {userid},
	}
	resp, err := http.PostForm("http://localhost:8080/configurations/user/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getRoutineConfgiurations(routineid string) *http.Response {
	data := url.Values{
		"routineid": {routineid},
	}
	resp, err := http.PostForm("http://localhost:8080/configurations/routine/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func createConfiguration(deviceid string, routineid string) *http.Response {
	data := url.Values{
		"deviceid":  {deviceid},
		"routineid": {routineid},
	}
	resp, err := http.PostForm("http://localhost:8080/configuration/create/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func updateConfiguration(configid string, offset string, deviceid string) *http.Response {
	data := url.Values{
		"configid": {configid},
		"offset":   {offset},
		"deviceid": {deviceid},
	}
	resp, err := http.PostForm("http://localhost:8080/configuration/update/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func deleteConfiguration(id string) *http.Response {
	data := url.Values{
		"id": {id},
	}
	resp, err := http.PostForm("http://localhost:8080/configuration/delete/", data)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func main() {
	//_ = createUserProfile("Lindsay", "LJamSupreme", "Password")
	_ = login("LJamSupreme", "Password")
}
