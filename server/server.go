package main

import (
	"bytes"
	"log"
	"net/http"
	"fmt"
	"rss_feed_reader/server/sql"
	"rss_feed_reader/server/feed"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/rss/select", selectHandler)
	http.HandleFunc("/rss/insert", insertHandler)
	http.HandleFunc("/rss/delete", deleteHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
		http.ServeFile(w, r, "../front/index.html")
	}
}

func selectHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		var data []byte
		links := sql.GetLink()
		for _, link := range links {
			fmt.Println(link)
      data = feed.Feed(link)
		}
    

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		fmt.Fprint(w, string(data))
	}

	http.Redirect(w, r, "/", 301)
}

func insertHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		sql.InsertLink(body)
	}
	http.Redirect(w, r, "/", 301)
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		sql.DeleteLink(body)
	}
	http.Redirect(w, r, "/", 301)
}
