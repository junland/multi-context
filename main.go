package main

import (
    "os"
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"   
    }
    
    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    
    hostname, _ := os.Hostname()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/App1", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi, this is App 1 on %s\n", hostname)
        fmt.Fprintf(os.Stdout, "Got a request for this app: %s\n", hostname)
    })
    
    http.HandleFunc("/App2", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi, this is App 2 on %s\n", hostname)
        fmt.Fprintf(os.Stdout, "Got a request for this app: %s\n", hostname)
    })
    
    http.HandleFunc("/App3", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi, this is App 3 on %s\n", hostname)
        fmt.Fprintf(os.Stdout, "Got a request for this app: %s\n", hostname)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))

}
