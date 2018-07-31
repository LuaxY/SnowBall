package main

import (
    "log"
    "net/http"

    "SnowBall/handler"

    "github.com/gorilla/mux"
    "github.com/throttled/throttled"
    "github.com/throttled/throttled/store/memstore"
)

func main() {
    r := mux.NewRouter()

    store, err := memstore.New(65536)

    if err != nil {
        log.Fatal(err)
    }

    quota := throttled.RateQuota{
        MaxRate:  throttled.PerMin(8),
        MaxBurst: 3,
    }

    rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)

    if err != nil {
        log.Fatal(err)
    }

    httpRateLimiter := throttled.HTTPRateLimiter{
        RateLimiter: rateLimiter,
        VaryBy:      &throttled.VaryBy{Path: true},
    }

    //r.HandleFunc("/", handler.Home)
    r.HandleFunc("/form", handler.Form)

    r.Handle("/game", httpRateLimiter.RateLimit(&handler.Game{})).
        Methods("POST")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

    srv := &http.Server{
        Handler: r,
        Addr:    "0.0.0.0:8000",
    }

    log.Fatal(srv.ListenAndServe())
}
