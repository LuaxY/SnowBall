package dofus

import (
    "bufio"
    "fmt"
    "html"
    "net/http"
    "os"
    "strings"

    "github.com/PuerkitoBio/goquery"
    "github.com/pkg/errors"
)

var Base = "https://www.dofus.com"

func GetMessages() ([]string, error) {
    file, _ := os.Open("./messages.txt")

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var messages []string

    for scanner.Scan() {
        messages = append(messages, scanner.Text())
    }

    return messages, nil
}

func GetForums() ([]string, error) {
    res, err := http.Get(fmt.Sprintf("%s%s", Base, "/fr/forum"))

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, errors.Errorf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)

    if err != nil {
        return nil, err
    }

    var forums []string

    doc.Find(".ak-container.ak-table").Each(func(i int, s *goquery.Selection) {
        if s.Find(".ak-lock").Length() > 0 {
            return
        }

        href, _ := s.Find("thead tr th a").Attr("href")

        if len(href) <= 0 {
            return
        }

        if strings.Contains(href, "sujets-fermes") {
            return
        }

        forums = append(forums, href)
    })

    return forums, nil
}

func GetThreads(forum string) ([]string, error) {
    res, err := http.Get(fmt.Sprintf("%s%s", Base, forum))

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, errors.Errorf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)

    if err != nil {
        return nil, err
    }

    var threads []string

    doc.Find(".ak-container tr").Each(func(i int, s *goquery.Selection) {
        if s.Find(".ak-lock").Length() > 0 {
            return
        }

        href, _ := s.Find(".ak-title-topic").Attr("href")

        if len(href) <= 0 {
            return
        }

        threads = append(threads, href)
    })

    return threads, nil
}

func GetPosts(thread string) ([]string, error) {
    res, err := http.Get(fmt.Sprintf("%s%s", Base, thread))

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, errors.Errorf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)

    if err != nil {
        return nil, err
    }

    var posts []string

    doc.Find(".ak-text").Each(func(i int, s *goquery.Selection) {
        text, _ := s.Find("html body p").Html()

        if len(text) <= 0 {
            return
        }

        text = strings.Replace(html.UnescapeString(text), "<br/>", "\n", -1)

        if strings.Contains(text, "<img") || strings.Contains(text, "<a") {
            return
        }

        if len(text) < 5 || len(text) > 500 {
            return
        }

        posts = append(posts, text)
    })

    return posts, nil
}

