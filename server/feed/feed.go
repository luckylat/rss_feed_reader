package feed

import (
  "fmt"
  "encoding/json"
  "github.com/mmcdole/gofeed"
)

type Data struct {
  Title string `json: "title"`
  Link string `json: "link"`
  Categories []string `json: "categories"`
}

func Feed(link string) ([]byte){
    fp := gofeed.NewParser();
    feed, _ := fp.ParseURL(link)
    fmt.Println(feed)

    items := feed.Items
    var list []Data
    for _, item := range items {
      info := Data{}
      
      fmt.Println(item)
      info.Title = item.Title
      info.Link = item.Link
      info.Categories = item.Categories

      list = append(list, info)
    }
    data, _ := json.MarshalIndent(&list, "", "\t")
    return data
}