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