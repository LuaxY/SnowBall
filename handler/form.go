package handler

import (
    "math/rand"
    "net/http"
    "text/template"

    "SnowBall/dofus"
)

type FormData struct {
    Base    string
    Thread  string
    Post    string
    Message string
}

func Form(w http.ResponseWriter, r *http.Request) {
    var err error
    var forum, thread string

    messages, err := dofus.GetMessages()

    if err != nil {
        panic(err)
    }

    if len(messages) <= 0 {
        panic("no message available")
    }

    message := messages[rand.Intn(len(messages))]

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

        /*posts, err := dofus.GetPosts(thread)

        if err != nil {
            panic(err)
        }

        if len(posts) <= 0 {
            continue
        }

        post = posts[rand.Intn(len(posts))]
        //log.Print(posts)*/

        break
    }

    /*if len(post) <= 0 {
        log.Print("no posts after 3 retry")
        return
    }*/

    tmpl, err := template.ParseFiles("./public/post.html")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "text/html")

    err = tmpl.ExecuteTemplate(w, "post.html", FormData{
        Base:   dofus.Base,
        Thread: thread,
        //Post:    post,
        Message: message,
    })

    if err != nil {
        panic(err)
    }
}
