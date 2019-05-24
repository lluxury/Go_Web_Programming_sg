package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init()  {
	var err error
	Db, err = sql.Open("postgres", "postgres://gwp:gwp@10.65.0.9/gwp?sslmode=disable")
	if err != nil {
		panic(err)
	}	// 连接数据库
	// Db.AutoMigrate(&Post{}, &Comment{})
}

func retrieve(id int) (post Post, err error) {
	// post - Post()
	post = Post{}
	// err = Db.QueryRow("select id, content, author from posts where id = $1,id").Scan(&post.Id, &post.Content, &post.Author)
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	
	return
}

func (post *Post) create()(err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update()(err error) {
	_, err = Db.Exec("update posts set content = $2, author =$3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error)  {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

