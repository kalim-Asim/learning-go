package main

import (
	"io/ioutil"
	"net/http"
)

const url = "https://cses.fi/problemset/task/1069"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	databyes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(databyes))
	println("Status Code:", response.StatusCode)
}
