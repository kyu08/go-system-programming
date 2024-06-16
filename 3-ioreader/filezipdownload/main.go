package main

import (
	"archive/zip"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zw := zip.NewWriter(w)
	defer zw.Close()

	aw, _ := zw.Create("a.txt")
	aw.Write([]byte("content of a"))
	bw, _ := zw.Create("b.txt")
	bw.Write([]byte("content of b"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
