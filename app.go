package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    var host, _ = os.Hostname()
    fmt.Fprintf(w, "Hi there, I path is %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "\n<h3>Host Name: %s!!</h3>", host)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}