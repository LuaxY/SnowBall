package handler

import (
    "net/http"
    "text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("./public/home.html")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "text/html")

    err = tmpl.ExecuteTemplate(w, "home.html", nil)

    if err != nil {
        panic(err)
    }
}
