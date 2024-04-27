package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homepage(w, r)
	case "/upload":
		uploadHandler(w, r)
	case "/dance":
		http.ServeFile(w, r, "./static/GangTortureDance(1).mp4")
	case "/video":
		handleVideo(w, r)
	case "/admin":
		fmt.Fprint(w, "u r on admin page")
	case "/video2":
		handleVideo2(w, r)
	default:
		fmt.Fprint(w, "error on url/path")
	}
}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fileName := "gpt.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("error while executing template", err)
		return
	}
}

func handleVideo(w http.ResponseWriter, r *http.Request) {
	fileName := "video.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("error while executing template", err)
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "upload.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("error while executing template", err)
		return
	}
}

func handleVideo2(w http.ResponseWriter, r *http.Request) {
	fileName := "video2.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("error while executing template", err)
		return
	}
}
