package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "strings"
)

const baseURL = "https://api.github.com/search/repositories"

func escapeSearch(s string) string {
    return strings.Replace(s, " ", "+", -1)
}

func searchString(q string, lang string, limit int) string {
    var buffer bytes.Buffer
    buffer.WriteString(baseURL)

    if q == "" {
        log.Fatal("You must enter a search query")
    }

    query := fmt.Sprintf("?q=%s", escapeSearch(q))
    buffer.WriteString(query)
    // return fmt.Sprintf("%s?q=%s+language:assembly&sort=stars&order=desc", baseURL, q)
    return buffer.String()
}

func requestSearch(url string) (r *Response, e error) {
    res, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }

    res.Header.Set("Accept", "application/vnd.github.preview")
    return http.Get(url)
}

func main() {
    fmt.Println(searchString("foo bar", "", 0))
}

