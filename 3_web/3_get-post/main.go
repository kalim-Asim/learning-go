package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "http://localhost:8080/get"

func main() {
	PerformGetRequest(myUrl)
	PerformPostJsonRequest(myUrl)
	PerformPostFormRequest(myUrl)
}

func PerformGetRequest(myUrl string) {
	// one method was in web-request/main.go
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// second method is using strings package

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	println(byteCount) // length of data
	println("Response String:", responseString.String())
}

func PerformPostJsonRequest(myUrl string) {
	requestBody := strings.NewReader(`
		{
			"name": "John Doe",
			"age": 30,
			"city": "New York"
		}
	`)

	response, _ := http.Post(myUrl, "application/json", requestBody)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	println("Response from POST request:", string(content))
}

func PerformPostFormRequest(myUrl string) {
	data := url.Values{}
	data.Add("name", "Jane Doe")
	data.Add("age", "25")

	response, _ := http.PostForm(myUrl, data)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	println("Response from POST Form request:", string(content))
}
