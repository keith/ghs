package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const baseURL = "https://api.github.com/search/repositories"
const helpers = "&sort=stars&order=desc&page=1&per_page="

type Query struct {
	Q string
	Lang string
	Limit int
}

func escapeSearch(s string) string {
	return strings.Replace(s, " ", "+", -1)
}

func searchString(q Query) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(baseURL)

	if q.Q == "" {
		return "", errors.New("You must enter a search query")
	}

	query := fmt.Sprintf("?q=%s", escapeSearch(q.Q))
	buffer.WriteString(query)

	if q.Lang != "" {
		lang := fmt.Sprintf("+language:%s", q.Lang)
		buffer.WriteString(lang)
	}

	limit := 10
	if q.Limit > 0 {
		limit = q.Limit
	}

	other := fmt.Sprintf("%s%d", helpers, limit)
	buffer.WriteString(other)

	return buffer.String(), nil
}

func requestSearch(url string, client http.Client) (r *http.Response, e error) {
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res.Header.Set("Accept", "application/vnd.github.preview")
	return client.Do(res)
}

func main() {
	fmt.Println(searchString(Query{"foo bar", "", 0}))
}
