package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	Author  string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB
func init()  {
	var err error
	Db, err = sql.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	if err != nil {
		panic(err)
	}	// 连接数据库
}

	// func (comment *Comment) create() (err error)  {
	func (comment *Comment) Create() (err error) {
	if comment.Post == nil{
		err = errors.New("Post not found")
		return 
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2,$3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return 
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	  post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	
	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}
	for rows.Next(){
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return 
}  //创建评论

// func (post *Post) Create() (err erros){
func (post *Post) Create() (err error){
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}


func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)                  // {1 Hello World! Sau Sheong [{1 Good post! Joe 0xc20802a1c0}]}
	fmt.Println(readPost.Comments)         // [{1 Good post! Joe 0xc20802a1c0}]
	fmt.Println(readPost.Comments[0].Post) // &{1 Hello World! Sau Sheong [{1 Good post! Joe 0xc20802a1c0}]}
}