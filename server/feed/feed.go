package feed

import (
	"encoding/json"
  "encoding/base64"
	"fmt"
  "os"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/mmcdole/gofeed"
)

type Data struct {
  Rss string `json: "rss"`
  Title string `json: "title"`
  Link string `json: "link"`
  Categories []string `json: "categories"`
  Image string `json: "image"`
}

func ImageLink(guid string) (string){
  now := time.Now()
	var layout = "2006/01/02"
	t := now.Format(layout)
    
	arr := strings.Split(guid, "/")
	url := arr[0] + "//i.gzn.jp/img/" + t + "/" + arr[4][9:] + "/00_m.jpg"
	fmt.Println(url)

  return url
}

func GetImage(url string) (string){
  res, err := http.Get(url)

  if err != nil {
    log.Fatal(err)
  }

  defer res.Body.Close()
  

  body, err := ioutil.ReadAll(res.Body)

  if err != nil {
    log.Fatal(err)
  }
  
  base64Data := base64.StdEncoding.EncodeToString(body)
  

  return base64Data
}


func Feed(links []string) ([]byte){
    var list []Data
    fp := gofeed.NewParser();
    for _, link := range links{
      feed, _ := fp.ParseURL(link)
      fmt.Println(feed)

      items := feed.Items
      for _, item := range items {
        info := Data{}
      
        info.Rss = link
        info.Title = item.Title
        info.Link = item.Link
        info.Categories = item.Categories
      
        if link == "https://gigazine.net/news/rss_2.0/" {
          url := ImageLink(item.GUID)
          info.Image = GetImage(url)
        } else if item.Image != nil {
          fmt.Println(item.Image.URL)
          info.Image = GetImage(item.Image.URL)
        } else {
          file, err := os.Open("./feed/rss.jpg")
        if err != nil {
          log.Fatal(err)
        }

        fi, _ := file.Stat()
        size := fi.Size()

        data := make([]byte, size)

        file.Read(data)

        info.Image = base64.StdEncoding.EncodeToString(data)
      }

      list = append(list, info)
    }
  }
  data, _ := json.MarshalIndent(&list, "", "\t")
  return data
}