

# [创建一个beego项目](/usr/local/go/bin/src/go-demo/doc/创建一个beego项目)

1. 安装beego（此过程略，注意设置好GOPATH）



2. 新建一个beego api项目，命名为 go-demo，命令如

   bee api go-demo

   新建成功之后，回在$GOPATH/src下面生成项目目录



3. 建议用GOLAND(此IDE比较成熟，建议使用，有30天试用期,如果不放弃，30天已经过足够你学成了，后面自己想办法)打开此项目，点击运行按钮，项目就运行起来了，运行成功之后，可以看到下面的信息输出：

   GOROOT=/usr/local/go #gosetup
   GOPATH=/usr/local/go/bin #gosetup
   /usr/local/go/bin/go build -o /private/var/folders/5m/hrs2j5g16h32vw0q4kyxdcz00000gn/T/___1go_build_main_go /usr/local/go/bin/src/go-demo/main.go #gosetup
   /private/var/folders/5m/hrs2j5g16h32vw0q4kyxdcz00000gn/T/___1go_build_main_go #gosetup
   2020/05/03 20:24:37.801 [I] [router.go:270]  /usr/local/go/bin/src/go-demo/controllers no changed
   2020/05/03 20:24:37.802 [I] [router.go:270]  /usr/local/go/bin/src/go-demo/controllers no changed
   2020/05/03 20:24:37.808 [I] [asm_amd64.s:1373]  http server Running on http://:8080



4. 项目默认生成了两个restful接口，在浏览器访问链接（http://127.0.0.1:8080/v1/user/user_11111），即可返回一些demo信息

   ```
   {
     "Id": "user_11111",
     "Username": "astaxie",
     "Password": "11111",
     "Profile": {
       "Gender": "male",
       "Age": 20,
       "Address": "Singapore",
       "Email": "astaxie@gmail.com"
     }
   }
   ```