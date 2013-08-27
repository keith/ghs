package main

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh/terminal"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const baseURL = "https://api.github.com/search/repositories"
const helpers = "&sort=stars&order=desc&page=1&per_page="

type Query struct {
	Q     string
	Lang  string
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

func printFromJSON(n int, b []byte) error {
	var j map[string]interface{}
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	items := j["items"].([]interface{})
	if len(items) < 1 {
		return errors.New("No matching repositories")
	}

	if len(items) < n {
		n = len(items)
	}

	for i := 0; i < n; i++ {
		repo := items[i].(map[string]interface{})
		// name := repo["name"].(string)
		url := repo["html_url"].(string)
		stars := int(repo["watchers"].(float64))
		lang := repo["language"].(string)
		fmt.Print(repoString(url, stars, lang))
	}

	return nil
}

func repoString(u string, s int, l string) string {
	url := strings.TrimPrefix(u, "https://github.com/")
	w, _, _ := terminal.GetSize(0)
	urlLen := utf8.RuneCountInString(url)
	starLen := utf8.RuneCountInString(strconv.Itoa(s))
	langLen := utf8.RuneCountInString(l)

	spaceLen := w - urlLen - starLen - langLen - 1
	if spaceLen < 1 {
		spaceLen := w - starLen - langLen - 1
		spaces := strings.Repeat(" ", spaceLen)
		return fmt.Sprintf("%s\n%s%d %s\n", url, spaces, s, l)
	}

	spaces := strings.Repeat(" ", spaceLen)
	return fmt.Sprintf("%s%s%d %s\n", url, spaces, s, l)
}

var count int

const countDefault = 10
const countHelp = "The number of results to return"

var lang string

const langHelp = "The language of the repo"

func main() {
	flag.Usage = Usage
	flag.IntVar(&count, "count", countDefault, countHelp)
	flag.IntVar(&count, "c", countDefault, countHelp+" (shorthand)")
	flag.StringVar(&lang, "language", "", langHelp)
	flag.StringVar(&lang, "l", "", langHelp+"(shorthand)")
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

	buffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = printFromJSON(count, buffer)
	if err != nil {
		log.Fatal(err)
	}
}
