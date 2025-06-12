package main

import "net/url"

const myUrl = "https://cses.fi/problemset/task/1069?task=1069&lang=en"

func main() {
	result, _ := url.Parse(myUrl)

	println("Scheme:", result.Scheme)
	println("Host:", result.Host)
	println("Path:", result.Path)
	println("Raw Query:", result.RawQuery)
	println(result.Port())

	qparams := result.Query()
	for key, value := range qparams {
		println("Key:", key, "Value:", value[0])
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "cses.fi",
		Path: 	 "/problemset/task/1069",
		RawPath: "user=asim",
	}
	println("Constructed URL:", partsOfUrl.String())
}
