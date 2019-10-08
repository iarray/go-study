package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Cat struct {
	Name  string
	Color string
}

func main() {
	//格式为[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	session, err := mgo.Dial("mongodb://hph:123456@127.0.0.1:27017/test")
	handlerErr(&err)

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Cat{"黑咪", "Black"})
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
