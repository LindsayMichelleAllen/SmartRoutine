package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	/*
			data := url.Values{
				"username": {"LJam"},
				"name":     {"Lindsay"},
			}

			_, err := http.PostForm("http://localhost:8080/create/user", data)

			if err != nil {
				log.Fatal(err)
			}

				data = url.Values{
					"userId": {"123456789"},
				}

				_, err = http.PostForm("http://localhost:8080/user/", data)

				if err != nil {
					log.Fatal(err)
				}

				_, err = http.Get("http://localhost:8080/users/")

				if err != nil {
					log.Fatal(err)
				}




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

			data = url.Values{
				"name":   {"DeviceName"},
				"userId": {"123456789"},
			}

			_, err = http.PostForm("http://localhost:8080/device/create", data)

			if err != nil {
				log.Fatal(err)
			}

		data := url.Values{
			"deviceid": {"987654321"},
		}

		_, err := http.PostForm("http://localhost:8080/device/", data)

		if err != nil {
			log.Fatal(err)
		}

		_, err = http.Get("http://localhost:8080/device/all")

		if err != nil {
			log.Fatal(err)
		}

		data = url.Values{
			"userid": {"123456789"},
		}

		_, err = http.PostForm("http://localhost:8080/device/user", data)
	*/
	/* Create routine given userid
	Create configuration with routineid and deviceid */

	data := url.Values{
		"userId": {"123456789"},
		"name":   {"RoutineName"},
	}

	_, err := http.PostForm("http://localhost:8080/routine/create", data)

	if err != nil {
		log.Fatal(err)
	}

	data = url.Values{
		"offset":    {"10"},
		"deviceId":  {"987654321"},
		"routineId": {"976431852"},
	}

	_, err = http.PostForm("http://localhost:8080/routine/configuration/create", data)

	if err != nil {
		log.Fatal(err)
	}

	data = url.Values{
		"routineId": {"976431852"},
	}

	_, err = http.PostForm("http://localhost:8080/device/routine", data)

	if err != nil {
		log.Fatal(err)
	}
	/*

						data := url.Values{
							"name":     {"NewDeviceName"},
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


				data = url.Values{
					"userId": {"123456789"},
					"name":   {"RoutineName"},
				}

				resp, err = http.PostForm("http://localhost:8080/routine/create", data)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)

						data := url.Values{
							"routineId": {"RoutineID"},
							"name":      {"NewRoutineName1"},
						}

						resp, err := http.PostForm("http://localhost:8080/routine/update", data)

						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp)

					data := url.Values{
						"routineId": {"RoutineID"},
					}

					resp, err := http.PostForm("http://localhost:8080/routine/delete", data)

					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resp)

				data = url.Values{
					"offset":    {"10"},
					"deviceId":  {"976431852"},
					"routineId": {"RoutineID"},
				}

				resp, err = http.PostForm("http://localhost:8080/routine/configuration/create", data)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)

			data := url.Values{
				"configId": {"ConfigID"},
				"offset":   {"20"},
			}

			resp, err := http.PostForm("http://localhost:8080/routine/configuration/update", data)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)

		data := url.Values{
			"configId": {"ConfigID"},
		}

		resp, err := http.PostForm("http://localhost:8080/routine/configuration/delete", data)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	*/
}
