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
			if err := r.ParseForm(); err != nil {

				userResponse := userAcctMngr.CreateUserProfile(&userAcctMngr.UserProfileCreateRequest{
					Username: r.FormValue("username"),
					Name:     r.FormValue("name"),
				})
				if userResponse.Error != nil {
					http.Error(w, userResponse.Error.Error(), 500)
				} else {
					fmt.Fprintf(w, userResponse.User.GetUsername()+", "+userResponse.User.GetName()+", "+userResponse.User.GetId(), 200)
				}
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
