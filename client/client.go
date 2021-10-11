package main

import (
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
	/*
			data := url.Values{
				"username": {"LJamSupreme"},
				"name":     {"Lindsay Allen"},
				"id":       {"123456789"},
			}

			resp, err := http.PostForm("http://localhost:8080/modify/user", data)

			if err != nil {
				log.Fatal(err)
			}


		data := url.Values{
			"id": {"123456789"},
		}

		resp, err := http.PostForm("http://localhost:8080/delete/user", data)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
	*/
	data = url.Values{
		"name":   {"DeviceName"},
		"userId": {"123456789"},
	}

	resp, err = http.PostForm("http://localhost:8080/device/create", data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
	/*
			data := url.Values{
				"name":     {"DeviceName"},
				"deviceId": {"976431852"},
			}

			resp, err := http.PostForm("http://localhost:8080/device/update", data)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(resp)

		data := url.Values{
			"deviceId": {"976431852"},
		}

		resp, err := http.PostForm("http://localhost:8080/device/delete", data)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	*/
}
