package models

import ("github.com/astaxie/beego/orm")

// mysql -ublog -p'blog_pass' -h127.0.0.1
// 如果库不存在，go get -t 去下载
type Blog struct {

	Id int `orm:"column(id)"`

	BlogTitle string `orm:"column(blog_title)"`

	BlogContent string `orm:"column(blog_content)"`

	ViewCount int `orm:"column(view_count)"`

	CommentCount int `orm:"column(comment_count)"`

	Remark string `orm:"column(remark)`

}

func init() {
	orm.RegisterModel(new(Blog))
}