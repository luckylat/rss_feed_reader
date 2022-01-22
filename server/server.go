package main

import (
	"log"
	"net/http"
	"fmt"
	"rss_feed_reader/server/sql"
	"rss_feed_reader/server/feed"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/rss", rssHanadler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	} else {
		fmt.Println("Running server on http://localhost:8080")
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request){
  if(r.Method == "GET"){
		http.ServeFile(w, r, "../front/index.html")
	}
}

func rssHanadler(w http.ResponseWriter, r *http.Request){
	if(r.Method == "GET"){
		var data []byte
		links := sql.GetLink()
		for _, link := range links {
			fmt.Println(link)
      data = feed.Feed(link)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	}
	if(r.Method == "POST"){
		r.ParseForm()
		fmt.Println(r.Form)
		sql.AddLink(r.FormValue("link"))
	}
}
