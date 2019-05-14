package main

import (
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string
}
// 结构,保存帖子

var PostById map[int]*Post
var PostByAuthor map[string][]*Post
// 通过 ID 或名字获取帖子,
// 两个map来创建两种不同的映射

func store(post Post)  {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}
// 存储帖子

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id:1,Content:"Hello World!",Author:"Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)
// 建立变量,调用函数

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["Pedro"] {
		fmt.Println(post)
	}
}