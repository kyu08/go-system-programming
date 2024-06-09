package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	jsonBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jsonBuffer)
	jsonEncoder.Encode(source)
	fmt.Printf("original json: %v\n", jsonBuffer.String())

	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Write(jsonBuffer.Bytes())
	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
