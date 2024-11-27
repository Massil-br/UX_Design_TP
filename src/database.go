package src

import (
	"UX_Design_TP/class"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDb(){
	var err error
	db, err = sql.Open("sqlite3", "./UX_Design.db")
	if err != nil{
		log.Fatal(err)
	}
	createTable()
}

func createTable(){
	Table := `CREATE TABLE IF NOT EXISTS Products(
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(100) NOT NULL ,
	"description" TEXT,
	"price" REAL DEFAULT 0.0,
	"image_url" TEXT
	);`
	statement, err := db.Prepare(Table)
	if err != nil{
		log.Fatal(err)
	}
	statement.Exec()
	fmt.Println("Products table created successfully.")
}

func GetPosts(page int, limit int) ([]class.Post, error){
	offset := (page -1) * limit
	rows, err := db.Query("SELECT id, name, description, price, image_url FROM Products LIMIT ? OFFSET ?" , limit, offset)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var posts []class.Post

	for rows.Next() {
		var post class.Post
		 err := rows.Scan(post.GetIDAdress(), post.GetNameAdress(), post.GetDescriptionAdress(),post.GetPriceAdress(), post.GetImageUrlAdress())
		 if err != nil {
			return nil, err
			}
			//fmt.Printf("Post:%v\n", post)  // working
		posts = append(posts, post)
	}
	return posts, nil
}