package main

import (
    "log"
    "net/http"

    "SnowBall/handler"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", handler.Home)
    r.HandleFunc("/form", handler.Form)

    srv := &http.Server{
        Handler: r,
        Addr:    "127.0.0.1:8000",
    }

    log.Fatal(srv.ListenAndServe())
}
