package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rverton/webanalyze"
)

func analyze(url string) (tools []string, err error) {
	appsFile, _ := os.Open("apps.json")
	wa, _ := webanalyze.NewWebAnalyzer(appsFile, nil)

	job := webanalyze.NewOnlineJob(url, "", nil, 0, false)
	result, _ := wa.Process(job)

	for _, a := range result.Matches {
		tools = append(tools, a.AppName)
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

	// append https before website if it doesn't exist
	url := body.Url
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	log.Println("Analyzing", body.Url)
	tools, err := analyze(url)
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

	log.Println("app definition file updated from ", webanalyze.WappalyzerURL)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	mode := os.Getenv("mode")
	if mode == "" {
		mode = "development"
	}

	http.HandleFunc("/", handler)

	if mode == "production" {
		log.Fatal(
			http.ListenAndServeTLS(
				":"+port,
				"/etc/letsencrypt/live/workinmena.tech/fullchain.pem",
				"/etc/letsencrypt/live/workinmena.tech/privkey.pem",
				nil,
			),
		)
	} else {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}
