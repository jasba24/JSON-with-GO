package main

import (
	"encoding/json"
	"github.com/jasba24/JSONConverter"
	// "io/ioutil"
	"net/http"
)

type Course struct {
	Title string `json:"title"`
	NumberVideos int `json:"number_videos"`
}

type Courses  []Course

func main() {
	c := JSONConverter.Converter(Courses{}, "db.json")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c)
	})
	http.ListenAndServe(":8080", nil)
}