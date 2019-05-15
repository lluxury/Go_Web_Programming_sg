package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	// Author  string
	AuthorName string `db:author`

}

// var Db *sql.DB
var Db *sqlx.DB

func init()  {
	var err error
	// Db, err = sql.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	Db, err = sqlx.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	if err != nil {
		panic(err)
	}	// 连接数据库
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	// err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}	
// func (post *Post) Create() (err erros){
func (post *Post) Create() (err error){
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}


func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	post.Create()

	fmt.Println(post)                  // {1 Hello World! Sau Sheong [{1 Good post! Joe 0xc20802a1c0}]}
}