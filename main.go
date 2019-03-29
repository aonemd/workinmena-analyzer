package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	fmt.Println(analyze("https://github.com"))
}
