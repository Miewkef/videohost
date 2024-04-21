package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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
	default:
		fmt.Fprint(w, "error on url/path")
	}
}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
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
	// Check if the request method is POST
	if r.Method == http.MethodPost {
		// Parse the form data including the video file
		r.ParseMultipartForm(10 << 20) // 10 MB max file size
		file, _, err := r.FormFile("videoFile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Read the file content
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Handle the file content as needed (e.g., save to disk, process, etc.)
		// Example: Save the file to disk
		err = os.WriteFile("uploaded_video.mp4", fileBytes, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.Write([]byte("File uploaded successfully"))
	}
}
