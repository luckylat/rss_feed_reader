package sql

import (
	"fmt"
	"os"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func loadenv(){
	err := godotenv.Load("./sql/.env")

	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() (*sql.DB){
  loadenv();
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
  dbname := os.Getenv("DBNAME")

	link := username + ":" + password + "@/" + dbname

  db, err := sql.Open("mysql", link);
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open db!")
  return db
}

func AddLink(link string){
	db := ConnectDB()
  stmt, err := db.Prepare(`
	   INSERT INTO links(link, count_clicked)
		 VALUES (?,?)
	`)

	if err != nil {
		log.Fatal(err)
	}
	
	_, err = stmt.Exec(link, 0)
	
	defer stmt.Close()
  defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Add link!")
}

func GetLink() ([]string){
	var links []string
	db := ConnectDB()
	rows, err := db.Query("SELECT link FROM links");

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next(){
		var link string

		err := rows.Scan(&link)

		if err != nil {
			log.Fatal(err)
		}

		links = append(links, link)
	}
  
	fmt.Println(links)
	fmt.Println("Get links!")
	return links
}
