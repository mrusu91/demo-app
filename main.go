package main

import (
    "fmt"
    "net/http"
    "log"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
    log.Println("Responding to request")
    fmt.Fprintf(w, "Hello World!\n")
}

func main() {
    log.Println("App started on port 8888")
    panic("Let's crash :D")
    http.HandleFunc("/", sayHelloWorld)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal(err)
    }
}
