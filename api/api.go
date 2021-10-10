package main

import (
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
