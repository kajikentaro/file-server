package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		fs.ServeHTTP(w, r)
	}
}

func main() {
	os.MkdirAll("./file", os.ModePerm)

	fs := http.FileServer(http.Dir("./file"))
	http.Handle("/serve/", http.StripPrefix("/serve", cors(fs)))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	log.Println("Set up file server at port 20768")
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

	destination := fileheader.Filename
	f, err := os.Create(filepath.Join("./file", destination))
	if err != nil {
		http.Error(w, "create file"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "copy"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`
		<html>
			<head>
				<title>File upload</title>
			</head>
			<body>
				<p>success</p>
				<a href="/serve">view files list</a>
				<br/>
				<a href="/">home</a>
			</body>
		</html>`))
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		<html>
			<head>
				<title>File upload</title>
			</head>
			<body>
				<form method="post" action="/upload" enctype="multipart/form-data">
					<input type="file" name="upload">
					<input type="submit">
				</form>
				<a href="/serve">view files list</a>
			</body>
		</html>`))
}
