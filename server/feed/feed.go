package feed

import (
  "fmt"
  "encoding/json"
  "github.com/mmcdole/gofeed"
)

type data struct {
  Title string `json: "title"`
  Link string `json: "link"`
  Extensions []Extension `json: "extensions"`
}

type Extension struct {
  Dc Dc `json: "dc"`
}

type Dc struct {
  Date []Date `json: "date"`
  Subject []Subject `json: "subjects"`
}

type Date struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type Subject struct {
    Name string `json:  "name"`
    Value string `json: "value"`
}

func Feed(link string) ([]byte){
    fp := gofeed.NewParser();
    feed, _ := fp.ParseURL(link)
    fmt.Println(feed)

    items := feed.Items
    item := items[0]
    data, _ := json.MarshalIndent(&item, "", "\t")
    return data
}