package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/rverton/webanalyze"
)

func analyze(url string) (tools []string, err error) {
	file := ioutil.NopCloser(strings.NewReader(url))
	results, err := webanalyze.Init(4, file, "apps.json")

	for result := range results {
		for _, a := range result.Matches {
			tools = append(tools, a.AppName)
		}
	}

	return
}

func handler(rw http.ResponseWriter, req *http.Request) {
	body := struct {
		Url string `json:"url"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	tools, err := analyze(body.Url)
	if err != nil {
		panic(err)
	}

	output := struct {
		Tools []string `json:"tools"`
	}{
		tools,
	}

	json.NewEncoder(rw).Encode(output)
}

func init() {
	err := webanalyze.DownloadFile(webanalyze.WappalyzerURL, "apps.json")
	if err != nil {
		log.Fatalf("error: can not update apps file: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
