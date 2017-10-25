package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    var host, _ = os.Hostname()
    fmt.Fprintf(w, "Hi there, I path is %s!", r.URL.Path[1:])
    fmt.Fprint(w, "Host Name: %s!!", host)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}