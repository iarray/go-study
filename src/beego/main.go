package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type MyUser struct {
	Id   int32  `orm:"pk;column(userid);"`
	Name string `orm:"column(name);"`
	Age  int32  `orm:"column(age);"`
}

func init() {
	//注册驱动
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//设置默认数据库
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=test host=127.0.0.1 port=5432 sslmode=disable", 30)
	//注册定义的model
	orm.RegisterModel(new(MyUser))

	// 创建table
	orm.RunSyncdb("default", false, true)

}

func main() {
	o := orm.NewOrm()
	user := &MyUser{Id: 2, Name: "小红", Age: 18}
	//插入数据
	id, err := o.Insert(user)
	if err != nil {
		fmt.Println(err)

		//删除数据
		id, err = o.Delete(user)
		if err == nil {
			fmt.Println("Model Exists do delete , return id:", id)
		}
		return
	}

	fmt.Println("Insert success , id:", id)

	//更新
	user.Age = 20
	id, err = o.Update(user)
	//只更新Name
	//o.Update(user, "Name")
	// 指定多个字段
	// o.Update(user, "Field1", "Field2", ...)

	if err != nil {
		fmt.Println("Update Fail ", err)
		return
	}
	fmt.Println("Update success , id:", id)

	//查询单个model
	user1 := &MyUser{Id: 2}
	err = o.Read(user1)
	if err != nil {
		fmt.Println("Read Fail ", err)
		return
	}
	fmt.Println(user1)

	//条件查询多个数据
	var users []*MyUser
	qs := o.QueryTable("my_user") // 返回 QuerySeter
	qs.Filter("userid", 2)        // WHERE userid = 1
	qs.Filter("age__gt", 18)      // WHERE age > 18

	_, err = qs.All(&users)
	if err != nil {
		fmt.Println("Query Fail ", err)
		return
	}
	for _, u := range users {
		fmt.Println(u)
	}

	//原声sql
	var rawUsers []*MyUser
	ret := o.Raw("SELECT * FROM my_user")
	_, err = ret.QueryRows(&rawUsers)

	for _, u := range rawUsers {
		fmt.Println(u)
	}
}
