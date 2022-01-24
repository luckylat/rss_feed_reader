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

func InsertLink(link string){
	db := ConnectDB()
  stmt, err := db.Prepare(`
	   INSERT INTO links(link, count_clicked)
		 VALUES (?,?)
	`)

	if err != nil {
		log.Fatal(err)
	}
	
	_, err = stmt.Exec(link, 0)
  defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Add link!")
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
  
	defer db.Close()
	
	log.Println(links)
	log.Println("Get links!")
	return links
}

func DeleteLink(link string){
	db := ConnectDB()
	del, err := db.Prepare("DELETE FROM links WHERE link=?")

	if err != nil {
		log.Fatal(err)
	}

	del.Exec(link)
  log.Println("Delete link!")
	defer db.Close()

}
