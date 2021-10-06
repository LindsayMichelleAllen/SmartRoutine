package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	data := url.Values{
		"username": {"LJam"},
		"name":     {"Lindsay"},
	}

	resp, err := http.PostForm("http://localhost:8080/create/user", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}
