package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}
var id = 1

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(users)

	case "POST":
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		u.ID = id
		id++
		users = append(users, u)
		json.NewEncoder(w).Encode(u)

	case "PUT", "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid path", 400)
			return
		}
		uid, _ := strconv.Atoi(parts[2])
		for i, u := range users {
			if u.ID == uid {
				if r.Method == "PUT" {
					json.NewDecoder(r.Body).Decode(&users[i])
					users[i].ID = uid
					json.NewEncoder(w).Encode(users[i])
				} else {
					users = append(users[:i], users[i+1:]...)
					json.NewEncoder(w).Encode(map[string]string{"deleted": parts[2]})
				}
				return
			}
		}
		
		http.Error(w, "User not found", 404)
	}
}

func main() {
	http.HandleFunc("/users", handler)
	http.HandleFunc("/users/", handler)
	http.ListenAndServe(":8080", nil)
}
