**Go Web编程
[新加坡] 郑兆雄（Sau，Sheong，Chang)**

郑老师的 <Go Web编程> 是一本很好的web编程入门书籍,
该书讲解比较贴近生产环境,不像其他资料,只给出一个功能简单没什么实际用途的的helloworld 页面网站.

唯一的缺点就是代码实例不是很详细,比如第二章, 直接跳到章节末尾运行程序是,可以执行的, 但中间步骤没有git提交, 
对照文章学习很不方便, 在这里,**我会记录下自己阅读每一节的提交代码, **,也方便给同好一个印证.

说明:
1. 对应**go web 书籍**请在京东自行购买,
2. 随书源码下载 https://github.com/sausheong/gwp,这里是阅读笔记

<br>

##### 第一章

* **06-05 18:25** init 环境 
程序运行需要相应的**包环境**,比如我的:
GOPATH=C:\Users\yann\go
项目目录:
C:\Users\yann\go\src\yann\Go_Web_Programming_sg\chitchat

 如果不确认可也以用原文的地址
C:\Users\yann\go\src\github.com\sausheong\gwp

* **03-23 16:05**   [1.10 first_webapp](https://github.com/lluxury/Go_Web_Programming_sg/tree/d773bdfec4731991fdfdcb6fda12ce83f918074a)    
项目很简单, 进入目录 运行即可, 最后网页上可以看到 Hello World!

##### 第二章

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
不过对一下午的劳动有个小奖励: 使用下载的源码,连上数据库是可以运行起来的,使用 go build<br>
注意项目路径
C:\Users\yann\go\src\github.com\sausheong\gwp\Chapter_2_Go_ChitChat\chitchat<br>
访问 http://localhost:8080

##### 第三章

* **04-16 10:33**    [3.2.1 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/092d31cca82cdf816a5eb9befe00ec583afe5855) 
第三章的第一次代码,运行成功,最简单的Web服务器,默认端口80访问 <br>
404 page not found

* **04-16 11:01**    [3.2.2 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/0eb335d15b216ea304c6de0379b13b9984376411) 
配置带证书的web服务器,及自己生成证书,运行成功 <br>
go run gencert.go
go run server.go
注意,这次要用 https 访问, 运行成功,出现404页面
404 page not found

* **04-16 11:53**    [3.3.1 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/d2ec468b23cf3af107f2d335d454e2215d46293a) 
修改了默认的多路复用器,访问任何地址都显示hello world<br>
访问 localhost:8080 <br>
访问 localhost:8080/abc<br>

* **04-16 13:29**    [3.3.2 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/5d060a9398b3a45a5b1d3358fb7033852630a0cf) 
使用多个处理器,处理请求<br>
访问 localhost:8080/hello<br>
访问 localhost:8080/world<br>

* **04-16 13:38**    [3.3.3 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/68b24a2c1a4181982eec728c35dc72740a3c7ba9) 
使用处理器函数,处理请求<br>
访问 localhost:8080/hello<br>
访问 localhost:8080/world<br>

* **04-16 14:30**    [3.3.4 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/1933e0b139cacd8b58217ea8be63fb35d51c86c0) 
串联多个处理器或处理器函数 <br>
访问 localhost:8080/hello
注意看控制台输出

* **04-16 15:23**    [3.3.6 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/8d5cf7ca1c57f944d51c17ca7935e38250973c82) 
换用第三方多路复用器 <br>
访问 localhost:8080/hello/yann

* **04-16 15:42**    [3.4.0 Handling_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/ff309baa47063a1db1a1422e550ec7ee2b7df0da) 
导入包,启用http2,花费了很多时间,最终测试成功 <br> 
curl.exe -I --http2 --insecure https://localhost:8080/   <br>
HTTP/2 200 <br>
content-type: text/plain; charset=utf-8 <br>
content-length: 17 <br>
date: Tue, 16 Apr 2019 08:48:15 GMT  <br>

##### 第四章

* **04-17 15:18**    [4.1.3 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/fad48c90cff435687b07584e31d8b2c7ca7c2cdb) 
处理请求,首部 header <br>
访问 localhost:8080/headers  显示出首部信息

* **04-17 15:47**    [4.1.4 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/b3b3c005b1f3704ec98ae9be9c7b1a7d22b4c3c8) 
处理请求,主体 body<br>
curl -id "first_name=yann&last_name=cao" 127.0.0.1:8080/body

* **04-17 16:18**    [4.2.1 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/a24cfc34b7e774da4f0c6fa1596ec54a5865b55c) 
处理请求,表单 Form <br>
双击 client.html 文件,打开返回如下: <br>
map[hello:[yann world] post:[456] thread:[123]]

* **04-17 17:28**    [4.2.2 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/de4fef50f1018225f7e5d1815341c1b7a4ef4b3a) 
处理请求,表单方法 PostForm <br>
双击 client.html 文件,打开返回如下: <br>
map[hello:[yann] post:[456]]

* **04-17 18:17**    [4.2.3 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/10306b54baa32386774affc1b083d57b8e21efad) 
处理请求,表单方法 ParseMultipartForm <br>
双击 client.html 文件,打开返回如下: <br>
&{map[hello:[sau sheong] post:[456]] map[]}  <br>

* **04-18 14:21**    [4.2.4 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/79be6c022ad01dabfb0a0d70a03bbce3bd894c50) 
测试文件上传 fileupload <br>
双击 client.html 文件,选择hello文本文件上传,会返回文件内容: <br>

* **04-18 14:44**    [4.3.0 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/743bec984504d192bdbc61e9523283212c0c7f1b) 
处理请求,返回 json数据, 访问对应域名返回 json格式数据(1行) <br>
curl -i 127.0.0.1:8080/json <br>

* **04-18 16:06**    [4.4.2 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/bab255dbc3e6b7db97a65dee835e2088c0f6f07b) 
设置cookie <br>
注意cookie不能直接看到,需要点浏览器地址栏左边的小锁,或者使用postman <br>
second_cookie "yann's test" 127.0.0.1

* **04-18 16:59**    [4.4.3 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/47f952586b36a494536f17e70585532cd110695e) 
处理请求,获取 cookie, 并返回  <br>
curl -i 127.0.0.1:8080/get_cookie <br>
返回数据如下 <br>
[first_cookie="Go Web Programming"; second_cookie="yann's test"]

* **04-18 17:59**    [4.4.4 Processing_Requests](https://github.com/lluxury/Go_Web_Programming_sg/tree/84d0b1929aa2fa3eefdda66a5d93c23dde6b3463) 
处理请求, flash 短信, 做一个看完就消失的短信,需要访问三次 <br>
curl -i 127.0.0.1:8080/set_message
    // 检查 cookie,确认信息
curl -i 127.0.0.1:8080/show_message
返回 Hello World!
curl -i 127.0.0.1:8080/show_message
返回 No message found

##### 第五章

* **04-23 18:55**    [5.21 base template](https://github.com/lluxury/Go_Web_Programming_sg/tree/0efa72d62b3752589d145020ad6961a652141ec1) 
初步了解下模板的使用方法 <br>
验证方法
浏览器访问 127.0.0.1:8080/process <br>

* **04-24 16:32**    [5.3.1 random_number](https://github.com/lluxury/Go_Web_Programming_sg/tree/07a00a22eb583eed485c7381261f3d2cf8ca98b6) 
动作-条件动作 if <br>
生成一个随机数,根据数字显示不同的文字
浏览器访问 127.0.0.1:8080/process <br>

* **04-24 18:35**    [5.3.2 iterator](https://github.com/lluxury/Go_Web_Programming_sg/tree/675a0d0a52f897e97dbe47821a4cd2cd1c146fc9) 
动作-迭代动作 range <br>
迭代显示出一周星期数
浏览器访问 127.0.0.1:8080/process <br>

* **04-24 18:41**    [5.3.3 set_dot](https://github.com/lluxury/Go_Web_Programming_sg/tree/9d2455292ee8bb1966f063e1e50a9d229ad1d525) 
动作-设置动作 set <br>
设置字段
浏览器访问 127.0.0.1:8080/process <br>

* **04-24 19:05**    [5.3.4 include](https://github.com/lluxury/Go_Web_Programming_sg/tree/4ca806291d860e905c404555c28f9d7344363290) 
动作-包含动作 include <br>
设置字段
浏览器访问 127.0.0.1:8080/process <br>

* **04-25 11:26**    [5.5.0 custom_function](https://github.com/lluxury/Go_Web_Programming_sg/tree/0e6af82dc54e3d37e0ad30e8f7c0e83f9f2f6ae5) 
自定义函数 <br>
日期转换函数, 数字显示有些异常
浏览器访问 127.0.0.1:8080/process <br>

* **04-25 13:24**    [5.6.0 context_aware](https://github.com/lluxury/Go_Web_Programming_sg/tree/6aeb448bf129af8c707286a935d50280a1c3b5b4) 
上下文感知 <br>
如果模板显示的是 html 格式的内容,那么模板将对其实施 html 转义<br>
如果模板显示的是 JavaScript 格式的内容,那么模板将对其实施 JavaScript 转义<br>
浏览器访问 127.0.0.1:8080/process <br>

* **04-25 13:41**    [5.6.1 XSS](https://github.com/lluxury/Go_Web_Programming_sg/tree/7419a6f7c239c24ffcbeaf3451b35b06b454ca56) 
防御XSS攻击 <br>
浏览器访问 127.0.0.1:8080/form <br>
攻击脚本:
<script>alert('Pwnd!');</script>

* **04-25 14:59**    [5.6.2 no escape](https://github.com/lluxury/Go_Web_Programming_sg/tree/eee872edcb96aa75b7e34970591c779616519c3b) 
取消转 <br>
这个测试失败了, 无论 chrome还是火狐全没有显示出弹窗
浏览器访问 127.0.0.1:8080/form <br>
攻击脚本:
<script>alert('Pwnd!');</script>

* **04-25 15:26**    [5.7.0 content](https://github.com/lluxury/Go_Web_Programming_sg/tree/e4d721c019fcc37b888194d3a5fe1ccbe0671f0f) 
模板嵌套 <br>
多次刷新浏览器,有时显示红色字符,有时蓝色<br>
浏览器访问 127.0.0.1:8080/process <br>

* **04-25 16:03**    [5.8.0 block action](https://github.com/lluxury/Go_Web_Programming_sg/tree/a0d5fff02f2cba021e335b332b97e7b09436f2ef) 
块动作定义默认模板 <br>
多次刷新浏览器,有时显示红色字符,有时蓝色<br>
浏览器访问 127.0.0.1:8080/process <br>

![http://hero.iamyann.com/blog_webchat.jpg](http://hero.iamyann.com/blog_webchat.jpg)<br>
Go_Web_Programming_sg go web 



05-14 15:55
