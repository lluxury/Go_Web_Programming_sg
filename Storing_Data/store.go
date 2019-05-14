package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// 结构,保存帖子

func store(data interface{}, filename string) {		// 存储数据
	// buffer := new(bytes.buffer)
	buffer := new(bytes.Buffer)		// 字节缓存区
	encoder := gob.NewEncoder(buffer)  // 编码器
	err := encoder.Encode(data)    // 编码
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {				// 写入文件
		panic(err)
	}
}

func load(data interface{}, filename string) {		// 载入数据
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
