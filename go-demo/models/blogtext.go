package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// mysql -ublog -p'blog_pass' -h127.0.0.1
// 如果库不存在，go get -t 去下载
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
	// set default database
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

}