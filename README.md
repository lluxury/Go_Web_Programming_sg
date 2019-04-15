Go Web编程
[新加坡] 郑兆雄（Sau，Sheong，Chang)

郑老师的 <Go Web编程> 是一本很好的web编程入门书籍,
该书讲解比较贴近生产环境,不像其他资料,只给出一个功能简单没什么实际用途的的helloworld 页面网站.

唯一的缺点就是代码实例不是很详细,比如第二章, 直接跳到章节末尾运行程序是,可以执行的, 但中间步骤没有git提交, 
对照文章学习很不方便, 在这里,我会记录下自己阅读每一节的提交代码, 也方便给同好一个印证.

说明:
1. 这里只会有一个日常操作的记录,详细的笔记在子目录内,
2. 程序运行需要相应的包环境,比如我的:

GOPATH=C:\Users\yann\go
项目目录:
C:\Users\yann\go\src\yann\Go_Web_Programming_sg\chitchat

 如果不确认可也以用原文的地址
C:\Users\yann\go\src\github.com\sausheong\gwp



03-23 16:05
1.10  first_webapp  
项目很简单, 进入目录 运行即可, 最后网页上可以看到 Hello World!

04-15 16:14
2.4.1 chitchat   开始有代码, 尝试运行一下会报找不到 index.  正常报错 因为index函数还没写
.\main.go:14:22: undefined: index

04-15 16:35
2.4.3 chitchat   涉及的template,和data包现在还没有写,依然无法运行 
.\main.go:28:15: undefined: template
.\main.go:29:18: undefined: data

04-15 17:46
2.4.4 chitchat  新增2个文件,修改1个,报错更多,因为涉及的模块还没写,继续下去
.\main.go:35:22: cannot refer to unexported name data.threads
.\main.go:35:22: undefined: data.threads
.\main.go:36:20: undefined: session
.\main.go:45:14: undefined: templates
.\main.go:45:34: template.Must undefined (type *"html/template".Template has no field or method Must)
.\main.go:47:14: undefined: templates
.\main.go:47:34: template.Must undefined (type *"html/template".Template has no field or method Must)
.\main.go:49:10: undefined: templates







# Go_Web_Programming_sg
go web 
