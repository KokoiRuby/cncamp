// create a HTTP server
// 1. recv req from client, write from req header to res header
// 2. read ENV VERSION and write into res header
// 3. server log - client ip, http status code to stdout
// 4. return 200 when access localhost/healthz

package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// global log level
	flag.Set("v", "0")
	glog.V(2).Info("Starting HTTP Server...")
	glog.V(2).Info("Listening on http://localhost:8080")

	http.HandleFunc("/", rootHandler) // reg handler to endpoint
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil) // listen & service on http://localhost:8080
	if err != nil {
		log.Fatal(err) // log if err
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 3. server log - client ip, http status code to stdout
	clientIP := r.RemoteAddr
	statusCode := http.StatusOK
	fmt.Fprintf(w, "Client IP: %s\nHTTP Status Code: %d\n\n", clientIP, statusCode)

	// 1. recv req from client, write from req header to res header
	io.WriteString(w, "=================== Details of HTTP Request Headers ============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	// 2. read ENV VERSION and write into res header
	// If using WSL:
	//     a. export VERSION=1.0
	//     b. go run exe2.go
	ver := os.Getenv("VERSION")
	//all := os.Environ()
	//fmt.Println("ENV - VERSION is:", ver)
	//fmt.Println("All:", all)

	// set response header
	w.Header().Set("Version", ver)

	io.WriteString(w, "\n=================== Details of HTTP Response Headers ============\n")
	for k, v := range w.Header() {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

// 4. return 200 when access localhost/healthz
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}
