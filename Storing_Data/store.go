package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct {
	Id int
	Content string
	Author string
}
// 结构,保存帖子


func main() {
	csvFile, err := os.Create("posts.csv")	// 创建CSV
	if err !=nil{
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id:1,Content:"Hello World!",Author:"Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)	// 写入器
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author} // 每个帖子一个切片
		err := writer.Write(line)
		if err != nil{
			panic(err)
		}
	}
	writer.Flush()	// 清缓存
	

	file, err := os.Open("posts.csv")	// 打开CSV
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)	// 读取器
	reader.FieldsPerRecord = -1		// 缺失字段不中断,正数为保证字段值,0为自适应
	record, err := reader.ReadAll()	// 返回切片组成的结果
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0],0,0)
		post := Post{Id:int(id), Content:item[1], Author:item[2]}
		posts = append(posts,post) // 组成结果
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}

