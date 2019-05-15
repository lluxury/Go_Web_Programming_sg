package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	Author  string
}

var Db *sql.DB
func init()  {
	var err error
	Db, err = sql.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	if err != nil {
		panic(err)
	}	// 连接数据库
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return	//获取全部帖子
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}  //获取单个帖子

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}
// 新建帖子

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}
//更新帖子

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
// 删除帖子


func DeleteAll() (err error) {
	_, err = Db.Exec("delete from posts")
	return
}
// 删除全部

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	// Create a post
	fmt.Println(post) // {0 Hello World! Sau Sheong}
	post.Create()
	fmt.Println(post) // {1 Hello World! Sau Sheong}

	// Get one post
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) // {1 Hello World! Sau Sheong}

	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts
	posts, _ := Posts(10)
	fmt.Println(posts) // [{1 Bonjour Monde! Pierre}]

	// Delete the post
	readPost.Delete()

	// Get all posts
	posts, _ = Posts(10)
	fmt.Println(posts) // []

	// Delete all posts
//   DeleteAll()
}