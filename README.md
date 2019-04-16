**Go Web编程
[新加坡] 郑兆雄（Sau，Sheong，Chang)**

郑老师的 <Go Web编程> 是一本很好的web编程入门书籍,
该书讲解比较贴近生产环境,不像其他资料,只给出一个功能简单没什么实际用途的的helloworld 页面网站.

唯一的缺点就是代码实例不是很详细,比如第二章, 直接跳到章节末尾运行程序是,可以执行的, 但中间步骤没有git提交, 
对照文章学习很不方便, 在这里,**我会记录下自己阅读每一节的提交代码, ***,也方便给同好一个印证.

说明:
1. 对应**go web 书籍**请在京东自行购买,
2. 随书源码下载 https://github.com/sausheong/gwp,这里是阅读笔记

<br>

* **06-05 18:25** init 环境 
程序运行需要相应的**包环境**,比如我的:
GOPATH=C:\Users\yann\go
项目目录:
C:\Users\yann\go\src\yann\Go_Web_Programming_sg\chitchat

 如果不确认可也以用原文的地址
C:\Users\yann\go\src\github.com\sausheong\gwp

* **03-23 16:05**   [1.10 first_webapp](https://github.com/lluxury/Go_Web_Programming_sg/tree/d773bdfec4731991fdfdcb6fda12ce83f918074a)    
项目很简单, 进入目录 运行即可, 最后网页上可以看到 Hello World!

* **04-15 16:14**    [2.4.1 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/e3d7c7b74848f219469a6090f813ff1008578fc7)   
开始有代码, 尝试运行一下会报找不到 index.  正常报错 因为index函数还没写 <br>
.\main.go:14:22: undefined: index

* **04-15 16:35**    [2.4.3 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/5a309c58865ba96e380700cf8734492e32d488df)   
涉及的template,和data包现在还没有写,依然无法运行 
.\main.go:28:15: undefined: template
.\main.go:29:18: undefined: data

* **04-15 17:46**    [2.4.4 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/88d14ba2a50d3ff43a0d6e970d80dc3efb2f642f)  
新增2个文件,修改1个,报错更多,因为涉及的模块还没写,继续下去
.\main.go:35:22: cannot refer to unexported name data.threads<br>
.\main.go:35:22: undefined: data.threads<br>
.\main.go:36:20: undefined: session<br>
.\main.go:45:14: undefined: templates<br>
.\main.go:45:34: template.Must undefined (type *"html/template".Template has no field or method Must)<br>
.\main.go:47:14: undefined: templates<br>
.\main.go:47:34: template.Must undefined (type *"html/template".Template has no field or method Must)<br>
.\main.go:49:10: undefined: templates<br>

* **04-15 18:23**   [2.5.0 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/2e973e3ed446a73b779407f41eedf2506ad0c33b)  
新增3个模版文件,修改2个文件,仍然在报错,书上的代码和git上提供的,也有些不一样了<br>
.\main.go:31:25: cannot refer to unexported name data.threads<br>
.\main.go:31:25: undefined: data.threads<br>
.\main.go:32:23: undefined: session<br>
.\main.go:34:17: undefined: generateHTML<br>
.\main.go:36:17: undefined: generateHTML<br>

* **04-15 18:48**   [2.7.0 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/7ba3836bf43735736ce440f0ce0e418e4a909d51) 
好消息是pg数据库可以用了,命令行访问成功,对照源码做了一些修改, 仍然无法启动,不过报错在减少<br>
.\main.go:34:4: undefined: error_message<br>
.\main.go:36:23: undefined: session<br>
.\main.go:38:17: undefined: generateHTML<br>
.\main.go:40:17: undefined: generateHTML<br>

* **04-15 19:00**    [2.8.0 chitchat](https://github.com/lluxury/Go_Web_Programming_sg/tree/d14bd9c140242d4cd945ea00ba96f2f4ff9d5289) 
第二章的最后一节,我们自己敲的代码仍然没有运行起来,毕竟让其运行起来是全书的内容<br>
不过对一下午的劳动有个小奖励: 使用下载的源码,连上数据库是可以运行起来的,使用 go build,注意项目路径
C:\Users\yann\go\src\github.com\sausheong\gwp\Chapter_2_Go_ChitChat\chitchat<br>
http://localhost:8080



* **04-16 10:33**    [3.2.1 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/d14bd9c140242d4cd945ea00ba96f2f4ff9d5289) 
第三章的第一次代码,运行成功,最简单的Web服务器,默认端口80访问 <br>
404 page not found

04-16 11:01





Go_Web_Programming_sg go web 

