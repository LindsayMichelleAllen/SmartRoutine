package main

import (
	"fmt"
	"log"
	"net/http"
	"server/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Landing Page")
	})

	http.HandleFunc("/create/user", func(w http.ResponseWriter, r *http.Request) {
		user, err := models.CreateUser(&models.CreateUserRequest{
			Username: "LJam",
			Name:     "Lindsay",
			Id:       "123456789",
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			fmt.Fprintf(w, user.GetUsername()+", "+user.GetName()+", "+user.GetId(), 200)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
