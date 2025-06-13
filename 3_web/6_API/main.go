package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice float64 `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" &&
		c.CourseName == "" &&
		c.CoursePrice == 0 &&
		c.Author == nil
}

func main() {
	fmt.Println("Starting the server on port 8080...")
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(courses) == 0 {
		http.Error(w, "No courses found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(courses)
}
