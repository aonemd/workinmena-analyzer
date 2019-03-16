package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/rverton/webanalyze"
)

func main() {
	file := ioutil.NopCloser(strings.NewReader("https://github.com"))
	results, err := webanalyze.Init(4, file, "apps.json")
	if err != nil {
		log.Fatal("error initializing:", err)
	}

	for result := range results {
		fmt.Println(result)
	}
}
