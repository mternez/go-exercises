package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Choice struct {
	Text   string `json:"text"`
	ArcKey string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Choice `json:"options"`
}

func main() {

	arcTemplate, _ := template.ParseFiles("./templates/arc.html")

	stories := parseStories("gopher.json")

	port := "8080"

	serveMux := http.NewServeMux()

	var intro *Arc

	for key, value := range *stories {
		if key == "intro" {
			intro = &value
		}
		serveMux.HandleFunc("/"+key, func(w http.ResponseWriter, r *http.Request) {
			arcTemplate.Execute(w, &value)
		})
	}

	if intro == nil {
		panic("No intro found!")
	}

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arcTemplate.Execute(w, &intro)
	})

	fmt.Println("Listeninig on port : ", port)

	http.ListenAndServe(":"+port, serveMux)
}

func parseStories(filePath string) *map[string]Arc {

	// Read in one go
	file, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	stories := make(map[string]Arc)

	json.Unmarshal(file, &stories)

	return &stories
}
