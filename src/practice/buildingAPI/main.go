package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var courses []Course

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome)
	r.HandleFunc("/{id}",getCourseByID).Methods("GET")
	r.HandleFunc("/add", CreateCourse).Methods("POST")

	http.ListenAndServe(":4000", r)
}

type Course struct {
	ID   int64  `json:"id"`
	Name string `json:"course_name"`
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home buddy"))
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	dataBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var course Course
	_ = json.Unmarshal(dataBytes, &course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	w.Header().Add("content-type", "application/json")
}

func getCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, _ := vars["id"]
	id, _ := strconv.ParseInt(idStr, 10, 0)
	for _, course := range courses {
		if course.ID == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
	w.Header().Add("content-type", "application/json")
}

func init() {
	courses = []Course{
		Course{ID: 1, Name: "RBPB"},
		Course{ID: 2, Name: "secondBook"},
	}
}
