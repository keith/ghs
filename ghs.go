package main

import (
	"flag"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
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

func requestSearch(url string, client *http.Client) (r *http.Response, e error) {
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res.Header.Set("Accept", "application/vnd.github.preview")
	return client.Do(res)
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] query\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

var count int
const countDefault = 10
const countHelp = "The number of results to return"
var lang string
const langHelp = "The language of the repo"

func main() {
	flag.Usage = Usage
	flag.IntVar(&count, "count", countDefault, countHelp)
	flag.IntVar(&count, "c", countDefault, countHelp + " (shorthand)")
	flag.StringVar(&lang, "language", "", langHelp)
	flag.StringVar(&lang, "l", "", langHelp + "(shorthand)")
	flag.Parse()

	if flag.NArg() == 0 || flag.NArg() > 1 {
		flag.PrintDefaults()
		return
	}

	query := flag.Arg(0)
	url, err := searchString(Query{query, lang, count})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := &http.Client{}
	res, err := requestSearch(url, client)
	if err != nil {
		log.Fatal(err)
	}

	buffer, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}
}

