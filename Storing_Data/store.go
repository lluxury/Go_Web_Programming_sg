package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id int
	Content string
	Author  string `sql:"not null"`
	// AuthorName string `db:author`
	Comments []Comment
	CreatedAT time.Time
}

type Comment struct {
	Id int
	Content string
	Author  string `sql:"not null"`
	PostId int  `sql:"index"`
    CreatedAT time.Time
}

// var Db gorm.DB
var Db *gorm.DB

func init()  {
	var err error
	// Db, err = sqlx.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	Db, err = gorm.Open("postgres", "postgres://gwp:gwp@192.168.142.9/gwp?sslmode=disable")
	if err != nil {
		panic(err)
	}	// 连接数据库
	// Db.AuthoMigrate(&Post{}, &Comment{})
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	// post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	// {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}

	Db.Create(&post)  // 创建帖子
	fmt.Println(post)
	// {5 Hello World! Sau Sheong [] 2019-05-15 17:41:50.1946001 +0800 CST m=+0.059000001}
	
	comment := Comment{Content: "Good post!", Author: "Joe"}	// 添加评论
	Db.Model(&post).Association("Comments").Append(comment)
	
	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)  // 通过帖子获取评论
	  var comments []Comment
	  Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])  
	// {1 Good post! Joe 1 0001-01-01 00:00:00 +0000 UTC}
}