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

	fmt.Println(q)
	if q.Q == "" {
		return "", errors.New("You must enter a search query")
	}

	query := fmt.Sprintf("?q=%s", escapeSearch(q.Q))
	buffer.WriteString(query)
	// return fmt.Sprintf("%s?q=%s+language:assembly&sort=stars&order=desc", baseURL, q)
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
