package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)
	fmt.Println("200")
	result,_ := json.Marshal(``)
	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	srv.ListenAndServe()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
