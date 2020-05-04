package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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

// 自定义表名（系统自动调用）
func (u *Blog) TableName() string {
	return "blog_text"
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "blog:blog_pass@tcp(127.0.0.1:3306)/blog?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Blog))

	// create table
	//orm.RunSyncdb("default", false, true)
}

func GetBlog(id int) (*Blog,  error)  {

	//blogId, err1 := strconv.Atoi(id)

	//blogId := id
	//if err1 != nil{
	//	return nil,err1
	//}else{
	//	o := orm.NewOrm()
	//
	//	blog := Blog{Id : id}
	//
	//	err := o.Read(&blog)
	//
	//	if err != nil{
	//		return nil, err1
	//	}else{
	//		return  &blog, err1
	//	}
	//}
	//
	//return nil, err1

	o := orm.NewOrm()

	blog := &Blog{Id : id}

	err := o.Read(blog)

	if err != nil{
		return nil, err
	}else{
		return  blog, err
	}

}