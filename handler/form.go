package handler

import (
    "math/rand"
    "net/http"
    "text/template"
    "time"

    "SnowBall/dofus"
)

var common = `

Si non vous avez tester XXX ? https://example.com`

type FormData struct {
    Base    string
    Thread  string
    Message string
    Common  string
}

func Form(w http.ResponseWriter, r *http.Request) {
    var err error
    var forum, thread, message string

    rand.Seed(time.Now().Unix())

    for retry := 3; retry > 0; retry-- {
        forums, err := dofus.GetForums()

        if err != nil {
            panic(err)
        }

        if len(forums) <= 0 {
            continue
        }

        forum = forums[rand.Intn(len(forums))]
        //log.Print(forum)

        threads, err := dofus.GetThreads(forum)

        if err != nil {
            panic(err)
        }

        if len(threads) <= 0 {
            continue
        }

        thread = threads[rand.Intn(len(threads))]
        //log.Print(thread)

        messages, err := dofus.GetMessages(thread)

        if err != nil {
            panic(err)
        }

        if len(messages) <= 0 {
            continue
        }

        message = messages[rand.Intn(len(messages))]
        //log.Print(message)

        break
    }

    if len(message) <= 0 {
        panic("no message after 3 retry")
    }

    tmpl, err := template.ParseFiles("./public/post.html")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "text/html")

    err = tmpl.ExecuteTemplate(w, "post.html", FormData{
        Base:    dofus.Base,
        Thread:  thread,
        Message: message,
        Common:  common,
    })

    if err != nil {
        panic(err)
    }
}
