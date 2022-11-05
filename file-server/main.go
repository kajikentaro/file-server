package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World2")
}

// kajikentaro.clear-net.jp:20768
func main() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":20768", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		http.Error(w, "parse: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, fileheader, err := r.FormFile("upload")
	if err != nil {
		http.Error(w, "file"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	os.MkdirAll("./dest", os.ModePerm)
	destination := fileheader.Filename
	f, err := os.Create(filepath.Join("./dest", destination))
	if err != nil {
		http.Error(w, "create dest"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "copy"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("ok"))
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`<html>
 <head>
     <title>File upload</title>
 </head>
 <body>
 <form method="post" action="/upload" enctype="multipart/form-data">
   <input type="file" name="upload">
   <input type="submit">
 </form>
 </body>
 </html>`))
}
