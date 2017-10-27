package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/go-redis/redis"
)

var client *redis.Client

func handler(w http.ResponseWriter, r *http.Request) {
    /*client = redis.NewClient(&redis.Options{
        Addr: "redis:6379"
    })
    */
    
    result, err := client.Incr("counter").Result();
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    var host, _ = os.Hostname()
    fmt.Fprintf(w, "Hi there, I path is %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "\n<h3>Host Name: %s!!</h3>", host)
    fmt.Fprintf(w, "count: ", result, " err: ", err)
}

func main() {

    client := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    fmt.Println("Starting service")
    pong, err := client.Ping().Result();
    fmt.Println("pong: ", pong, " err: ", err)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}