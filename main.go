package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed routes.json
var routesJson string

type Route struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("t.xk.is"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Environment variable PORT must be set")
	}

	var routes []Route
	err := json.Unmarshal([]byte(routesJson), &routes)
	if err != nil {
		panic(err)
	}

	log.Println("routes.json loaded successfully.")

	for _, route := range routes {
		mux.Handle(route.Path, http.RedirectHandler(route.URL, http.StatusFound))
		log.Printf("Registering route %s -> %s\n", route.Path, route.URL)
	}

	log.Printf("Starting server on :%s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	log.Fatal(err)
}
