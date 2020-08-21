package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/hari.jpg", dogPic)
	http.ListenAndServe(":8000", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="C:/Users/Admin/Desktop/hari.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("hari.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
