package main

import "encoding/json"

type course struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

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
