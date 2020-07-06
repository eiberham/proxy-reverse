package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRequest)

	fmt.Println("Running reverse proxy ...")

	log.Fatal(http.ListenAndServeTLS(":3000", "./certs/server.cert", "./certs/server.key", r))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	if (*r).Method == "OPTIONS" {
		return
	}

	url := r.URL.Query().Get("url")
	if url != "" {
		httpClient := http.Client{
			Timeout: time.Second * 10,
		}

		resp, err := httpClient.Get(url)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		w.Write(body)
	}
}
