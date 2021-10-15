package main

import (
	dvcMngr "api/services/devicemanagement"
	rtnMngr "api/services/routinemanagement"
	userAcctMngr "api/services/useraccountmanagement"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Landing Page")
	})

	http.HandleFunc("/create/user", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/modify/user", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/delete/user", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/device/create", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/device/update", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/device/delete", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/routine/create", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/routine/update", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/routine/delete", func(w http.ResponseWriter, r *http.Request) {
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
