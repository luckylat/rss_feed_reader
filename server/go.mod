module rss_feed_reader/server

go 1.17

replace rss_feed_reader/server/sql => ./sql

replace rss_feed_reader/server/feed => ./feed

require (
	rss_feed_reader/server/feed v0.0.0-00010101000000-000000000000
	rss_feed_reader/server/sql v0.0.0-00010101000000-000000000000
)

require (
	github.com/PuerkitoBio/goquery v1.5.1 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mmcdole/gofeed v1.1.3 // indirect
	github.com/mmcdole/goxpp v0.0.0-20181012175147-0068e33feabf // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/text v0.3.2 // indirect
)
