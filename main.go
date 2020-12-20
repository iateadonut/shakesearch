package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
		page, ok := r.URL.Query()["page"]
		if !ok || len(query[0]) < 1 {

		}
		pageNumber, error := strconv.Atoi(page[0])
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			pageNumber = 0
		}
		results := json.RawMessage(Search(query[0], pageNumber))
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

func Search(query string, page int) string {

	baseURL := fmt.Sprintf("%s:%s/solr/%s/select?start=%d&", solrAddress, solrPort, solrCore, page*10)
	params := url.Values{}
	params.Add("q", query)

	fmt.Println(baseURL + params.Encode())

	resp, err := http.Get(baseURL + params.Encode())
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
