package main

import (
	dvcMngr "api/services/devicemanagement"
	rtnMngr "api/services/routinemanagement"
	cfgMngr "api/services/routinemanagement/configurationmanagement"
	userAcctMngr "api/services/useraccountmanagement"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Landing Page")
	})

	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id := r.FormValue("userId")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.GetUserProfile(&userAcctMngr.UserProfileGetRequest{
				Id: id,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName()+", "+userResponse.User.GetId(), 200)
			}
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.GetUserProfiles()
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				for _, user := range userResponse.Users {
					fmt.Fprintf(w, user.GetUsername()+", "+user.GetName()+", "+user.GetId(), 200)
				}
			}
		}
	})

	mux.HandleFunc("/create/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			username := r.FormValue("username")
			name := r.FormValue("name")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: username,
				Name:     name,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName()+", "+userResponse.User.GetId(), 200)
			}
		}
	})

	mux.HandleFunc("/modify/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			username := r.FormValue("username")
			name := r.FormValue("name")
			id := r.FormValue("id")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: username,
				Name:     name,
				Id:       id,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName()+", "+userResponse.User.GetId(), 200)
			}
		}
	})

	mux.HandleFunc("/delete/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.DeleteUserProfile(&userAcctMngr.UserProfileDeleteRequest{
				Id: id,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName()+", "+userResponse.User.GetId(), 200)
			}
		}
	})

	mux.HandleFunc("/device/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			deviceid := r.FormValue("deviceid")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetDevice(&dvcMngr.GetDeviceRequest{Id: deviceid})
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetDevices()
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			userid := r.FormValue("userid")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetUserDevices(&dvcMngr.GetUserDevicesRequest{UserId: userid})
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/routine", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			routineid := r.FormValue("routineId")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetRoutineDevices(&dvcMngr.GetRoutineDevicesRequest{RoutineId: routineid})
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			userId := r.FormValue("userId")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.CreateDevice(&dvcMngr.DeviceCreateRequest{
				Name:   name,
				UserId: userId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			deviceId := r.FormValue("deviceId")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.UpdateDevice(&dvcMngr.DeviceUpdateRequest{
				Name: name,
				Id:   deviceId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			deviceId := r.FormValue("deviceId")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.DeleteDevice(&dvcMngr.DeviceDeleteRequest{
				Id: deviceId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			userId := r.FormValue("userId")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.CreateRoutine(&rtnMngr.RoutineCreateRequest{
				Name:   name,
				UserId: userId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			routineId := r.FormValue("routineId")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.UpdateRoutine(&rtnMngr.RoutineUpdateRequest{
				Name: name,
				Id:   routineId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			routineId := r.FormValue("routineId")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.DeleteRoutine(&rtnMngr.RoutineDeleteRequest{
				Id: routineId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/configuration/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			offset := new(int)
			offsetInput, _ := strconv.Atoi(r.FormValue("offset"))
			*offset = offsetInput
			deviceId := r.FormValue("deviceId")
			routineId := r.FormValue("routineId")

			basicCfgMngr := cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.CreateConfiguration(&cfgMngr.CreateConfigurationRequest{
				Offset:    offset,
				DeviceId:  deviceId,
				RoutineId: routineId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/configuration/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			configId := r.FormValue("configId")
			offset := new(int)
			offsetInput, _ := strconv.Atoi(r.FormValue("offset"))
			*offset = offsetInput

			basicCfgMngr := cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.UpdateConfiguration(&cfgMngr.UpdateConfigurationRequest{
				ConfigId: configId,
				Offset:   offset,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/configuration/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			configId := r.FormValue("configId")

			basicCfgMngr := cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.DeleteConfiguration(&cfgMngr.DeleteConfigurationRequest{
				ConfigId: configId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
