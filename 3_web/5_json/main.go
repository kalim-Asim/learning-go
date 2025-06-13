package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
	DecodeJson()
}

// produce json data
func EncodeJson() {
	courses := []course{
		{
			Name:     "Go Programming",
			Password: "go123",
			Tags:     []string{"programming", "go", "backend"},
		},
		{
			Name:     "Python Programming",
			Password: "python123",
			Tags:     nil,
		},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	println(string(finalJson))
}

// consume json data
func DecodeJson() {
	jsonData := []byte(`
		{
				"name": "Go Programming",
				"tags": [
								"programming",
								"go",
								"backend"
				]
		}
	`)

	var Course course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("json is not valid")
	}

	// better way to decode json data
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}
}
