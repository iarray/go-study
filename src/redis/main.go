package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	//检测是否在线
	fmt.Println(pong)

	client.Set("key1", 123, 0)

	val, _ := client.Get("key1").Result()
	fmt.Println("key1=", val)

	//取不存在的key
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		//return
	}
	fmt.Println("key2=", val2)

	//设置有效期
	if _, err := client.SetNX("key3", "value", 1*time.Second).Result(); err == nil {
		val3, _ := client.Get("key3").Result()
		fmt.Println("key3=", val3)

		time.Sleep(1 * time.Second)

		val3, _ = client.Get("key3").Result()
		fmt.Println("key3=", val3)
	}

}
