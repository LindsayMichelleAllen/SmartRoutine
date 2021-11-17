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
	"strings"

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
			id := r.FormValue("id")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.GetUserProfile(&userAcctMngr.UserProfileGetRequest{
				Id: id,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName(), 200)
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
					fmt.Fprintf(w, user.GetUsername()+", "+user.GetName(), 200)
				}
			}
		}
	})

	mux.HandleFunc("/user/login/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			username := r.FormValue("username")
			password := r.FormValue("password")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			loginResponse := basicUsrMngr.UserProfileLogin(&userAcctMngr.UserProfileLoginRequest{
				Username: username,
				Password: password,
			})
			if loginResponse.Error != nil {
				http.Error(w, loginResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, "Success", 200)
			}
		}
	})

	mux.HandleFunc("/user/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			username := r.FormValue("username")
			name := r.FormValue("name")
			password := r.FormValue("password")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
				Username: username,
				Name:     name,
				Password: password,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName(), 200)
			}
		}
	})

	mux.HandleFunc("/user/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			username := r.FormValue("username")
			name := r.FormValue("name")

			basicUsrMngr := userAcctMngr.UnprotectedUserService{}
			userResponse := basicUsrMngr.UpdateUserProfile(&userAcctMngr.UserProfileUpdateRequest{
				Username: username,
				Name:     name,
			})
			if userResponse.Error != nil {
				http.Error(w, userResponse.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName(), 200)
			}
		}
	})

	mux.HandleFunc("/user/delete/", func(w http.ResponseWriter, r *http.Request) {
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
				fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName(), 200)
			}
		}
	})

	mux.HandleFunc("/device/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetDevice(&dvcMngr.GetDeviceRequest{Id: id})
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/devices/", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/device/user/", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/device/routine/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			routineid := r.FormValue("routineid")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.GetRoutineDevices(&dvcMngr.GetRoutineDevicesRequest{RoutineId: routineid})
			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/device/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			userId := r.FormValue("userid")

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

	mux.HandleFunc("/device/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			deviceId := r.FormValue("deviceid")

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

	mux.HandleFunc("/device/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicDvcSrvc := dvcMngr.UnprotectedDeviceService{}
			resp := basicDvcSrvc.DeleteDevice(&dvcMngr.DeviceDeleteRequest{
				Id: id,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			routineid := r.FormValue("routineid")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.GetRoutine(&rtnMngr.GetRoutineRequest{
				RoutineId: routineid,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else if resp.Routine != nil {
				fmt.Fprint(w, resp.Routine.GetJson())
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routines/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.GetRoutines()

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routines/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			userId := r.FormValue("userid")
			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.GetUserRoutines(&rtnMngr.GetUserRoutinesRequest{
				UserId: userId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				routineStrings := []string{}
				for _, r := range resp.Routines {
					routineStrings = append(routineStrings, r.GetJson())
				}

				fmt.Fprintf(w, "[%s]", strings.Join(routineStrings, ","))
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routines/device/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			deviceId := r.FormValue("deviceid")
			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.GetDeviceRoutines(&rtnMngr.GetDeviceRoutinesRequest{
				DeviceId: deviceId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				routineStrings := []string{}
				for _, r := range resp.Routines {
					routineStrings = append(routineStrings, r.GetJson())
				}

				fmt.Fprintf(w, "[%s]", strings.Join(routineStrings, ","))
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			userId := r.FormValue("userid")
			basealarm := r.FormValue("basealarm")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.CreateRoutine(&rtnMngr.RoutineCreateRequest{
				Name:      name,
				UserId:    userId,
				Basealarm: basealarm,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else if resp.Routine != nil {
				fmt.Fprint(w, resp.Routine.GetJson())
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			name := r.FormValue("name")
			routineId := r.FormValue("routineid")
			basealarm := r.FormValue("basealarm")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.UpdateRoutine(&rtnMngr.RoutineUpdateRequest{
				Basealarm: basealarm,
				Name:      name,
				Id:        routineId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/routine/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicRtnMngr := &rtnMngr.UnprotectedRoutineService{}
			resp := basicRtnMngr.DeleteRoutine(&rtnMngr.RoutineDeleteRequest{
				Id: id,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configuration/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicCfgMngr := &cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.GetConfiguration(&cfgMngr.GetConfigurationRequest{
				ConfigId: id,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				fmt.Fprintf(w, "[%s]", resp.Configuration.GetJson())
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configurations/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}

			basicCfgMngr := &cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.GetConfigurations()

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				configStrings := []string{}
				for _, c := range resp.Configurations {
					configStrings = append(configStrings, c.GetJson())
				}

				fmt.Fprintf(w, "[%s]", strings.Join(configStrings, ","))
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configurations/device/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			deviceId := r.FormValue("deviceid")

			basicCfgMngr := &cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.GetDeviceConfigurations(&cfgMngr.GetDeviceConfigurationsRequest{
				DeviceId: deviceId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				configStrings := []string{}
				for _, c := range resp.Configurations {
					configStrings = append(configStrings, c.GetJson())
				}

				fmt.Fprintf(w, "[%s]", strings.Join(configStrings, ","))
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configurations/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			userId := r.FormValue("userid")

			basicCfgMngr := &cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.GetUserConfigurations(&cfgMngr.GetUserConfigurationsRequest{
				UserId: userId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			} else {
				configStrings := []string{}
				for _, c := range resp.Configurations {
					configStrings = append(configStrings, c.GetJson())
				}

				fmt.Fprintf(w, "[%s]", strings.Join(configStrings, ","))
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configurations/routine/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			routineId := r.FormValue("routineid")

			basicCfgMngr := &cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.GetRoutineConfigurations(&cfgMngr.GetRoutineConfigurationsRequest{
				RoutineId: routineId,
			})

			if resp.Error != nil {
				http.Error(w, resp.Error.Error(), 500)
			}
			fmt.Fprint(w, "Success", 200)
		}
	})

	mux.HandleFunc("/configuration/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			offset := new(int)
			offsetInput, _ := strconv.Atoi(r.FormValue("offset"))
			*offset = offsetInput
			deviceId := r.FormValue("deviceid")
			routineId := r.FormValue("routineid")

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

	mux.HandleFunc("/configuration/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			configId := r.FormValue("configid")
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

	mux.HandleFunc("/configuration/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing request", 500)
			}
			id := r.FormValue("id")

			basicCfgMngr := cfgMngr.UnprotectedConfigurationService{}
			resp := basicCfgMngr.DeleteConfiguration(&cfgMngr.DeleteConfigurationRequest{
				ConfigId: id,
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
