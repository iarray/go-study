package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//注意字段首字母大写，不然不可见。通过bson:”name”这种方式可以定义MongoDB中集合的字段名,
//如果不定义，mgo自动把struct的字段名首字母小写并作为集合的字段名。如果不需要获得id_，Id_可以不定义，在插入的时候会自动生成。
type Cat struct {
	Id_   bson.ObjectId `bson:"_id"`
	Name  string        //mgo会把字段转换为name
	Color string        //mgo会把字段转换为color
}

func main() {
	//格式为[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	//mongodb先添加用户hph, 密码123456并分配角色权限, mongodb可视化管理工具可以用robo 3T
	session, err := mgo.Dial("mongodb://hph:123456@127.0.0.1:27017/test")
	handlerErr(&err)

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Cat{Name: "黑咪", Color: "Black"})
	handlerErr(&err)

	queryResult := Cat{}
	//注意这里的color是小写, 如果改成Color 会报错, 找不到
	err = c.Find(bson.M{"color": "Black"}).One(&queryResult)
	handlerErr(&err)

	fmt.Println(queryResult)
}

func handlerErr(err *error) {
	if *err != nil {
		fmt.Println(*err)
		panic(*err)
	}
}
