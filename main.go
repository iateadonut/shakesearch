package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	solrPort    = "8983"
	solrAddress = "http://solr.100wires.com"
	solrCore    = "shakespeare"
)

func main() {

	//http://127.0.0.1:8983/solr/shakespeare/select?q=oh%20my%20son

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSearch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := json.RawMessage(Search(query[0]))
		//fmt.Println(results)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func Search(query string) string {
	//fmt.Println(fmt.Sprintf("%s:%s/solr/%s/select?q=%s", solrAddress, solrPort, solrCore, query))

	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%s/solr/%s/select?q=%s", solrAddress, solrPort, solrCore, query), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("q", query)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := http.Get(req.URL.String())
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
	return string(body)
}
