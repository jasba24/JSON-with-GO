package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Course struct {
	Title string `json:"title"`
	NumberVideos int `json:"number_videos"`
}

type Courses  []Course

func main() {
	file_data, _ := ioutil.ReadFile("./db.json")

	c := Courses{}

	json.Unmarshal(file_data, &c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c)
	})
	http.ListenAndServe(":8080", nil)
}