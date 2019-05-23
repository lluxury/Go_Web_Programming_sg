package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id, attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id string 		`xml:"id,attr"`
	Content string 	`xml:"content`
	Author Author 	`xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for{
		t,err := decoder.Token()
		if err == io.EOF{
			break
		}

		if err != nil {
			fmt.Println("Error decoding XML into tokens:",err)
			return
		}

		// switch es := t.(type) {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment"{
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			} else if se.Name.Local == "author" {
				var author Author
				decoder.DecodeElement(&author, &se)
				fmt.Println(author,"\n")
			}
		// case xml.EndElement:  
		//     // 遇到结束节点，操作
        // case xml.CharData:  // 原始字符串，操作
        // case xml.Comment:  // 注释，操作

		}
	}

}
