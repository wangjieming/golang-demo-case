##### 连接MySQL单数据源

1. 假定已经在MySQL上建好了数据库，并且设置好账号密码

   数据库名：blog
   数据库表：blog_text

   blog_text表定义如下：

   CREATE TABLE `blog_text` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `blog_title` char(64) NOT NULL DEFAULT '' COMMENT '博文标题',
     `blog_content` text COMMENT '博文内容',
     `view_count` int(11) NOT NULL DEFAULT '0' COMMENT '阅读数',
     `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
     `remark` char(64) DEFAULT '' COMMENT '备注',
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='博客表'



2. 为blog_text表建model，文件名为blogtext.go，该文件内容如下，定义结构体，初始化函数，表名等。

   

   ```go
   package models
   
   // 如果库不存在，go get -t 去下载
   import (
      "github.com/astaxie/beego/orm"
      _ "github.com/go-sql-driver/mysql"
   )
   
   // mysql -ublog -p'blog_pass' -h127.0.0.1
   
   type BlogText struct {
   
      Id int `orm:"column(id)"`
   
      BlogTitle string `orm:"column(blog_title)"`
   
      BlogContent string `orm:"column(blog_content)"`
   
      ViewCount int `orm:"column(view_count)"`
   
      CommentCount int `orm:"column(comment_count)"`
   
      Remark string `orm:"column(remark)`
   
   }
   
   // 自定义表名（系统自动调用）
   func (u *BlogText) TableName() string {
      return "blog_text"
   }
   
   func init() {
      // 单一的数据源，第一个参数直接填default, 第二个参数是数据库的驱动名称，我们用的MySQL,所以这里是mysql，第三个参数中，第一个blog为用户名，blog_pass为密码，127.0.0.1为数据库主机名，3306为数据库端口，端口后面的blog为数据库名
      orm.RegisterDataBase("default", "mysql", "blog:blog_pass@tcp(127.0.0.1:3306)/blog?charset=utf8", 30)
   
      // register model
      orm.RegisterModel(new(BlogText))
   
   }
   
   //根据id获取一篇博文
   func GetBlogText(id int) (*BlogText,  error)  {
   
      o := orm.NewOrm()
   
      blog := &BlogText{Id : id}
   
      err := o.Read(blog)
   
      if err != nil{
         return nil, err
      }else{
         return  blog, err
      }
   ```



3. 在tests包下面建一个测试文件model_blogtext_test.go，该文件内容如下：

   ```go
   package test
   
   import (
      "encoding/json"
      "fmt"
      "github.com/astaxie/beego/orm"
      . "go-demo/models"
      "testing"
   )
   
   //测试数据的连通性
   func TestMysqlSingleSource(t *testing.T)  {
   
      o := orm.NewOrm()
   
      // 查询条件
      u := BlogText{Id: 1}
      // 根据查询条件进行查询
      err := o.Read(&u)
      //如果正常，则打印返回的数据
      if err == nil{
         bytes, _ := json.Marshal(u)
         fmt.Println("Query Result: ",string(bytes) )
      }else{
         //如果不正常，则打印返回的异常
         fmt.Printf("ERR: %v\n", err)
         fmt.Println()
      }
   
   }
   ```



4. 运行以上测试方法，如果匹配到数据，应该会有正常的查询结果，比如：

   === RUN   TestMysqlSingleSource
   Query Result:  {"Id":1,"BlogTitle":"测试标题","BlogContent":"测试内容","ViewCount":33,"CommentCount":133,"Remark":""}
   --- PASS: TestMysqlSingleSource (0.01s)
   PASS