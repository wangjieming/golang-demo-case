package test


import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-demo/models"
	"testing"

)

func init() {
//	orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterDataBase("default", "mysql", "blog:blog_pass@tcp(127.0.0.1:3306)/blog?charset=utf8", 30)

//	orm.RegisterModel(new(models.Blog))

}

func TestMysql(t *testing.T)  {

	o := orm.NewOrm()
	// read one
	u := models.Blog{Id: 2}

	err := o.Read(&u)
	if err == nil{
		fmt.Println("-----u-----" + u.BlogContent)
	}else{
		fmt.Printf("ERR: %v\n", err)
		fmt.Println()
	}

	u1 := models.Blog{Id : 1}
	err1 := o.Read(&u1)

	if err1 == nil{
		fmt.Println("------u1------" + u1.BlogContent)
	}else{
		fmt.Printf("U1 ERR: %v\n", err1)
	}


}