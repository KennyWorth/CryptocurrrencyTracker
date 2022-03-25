package redisdb

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func Ping() error {
	client := RClient()
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

func Set(key, value string) error {
	client := RClient()
	err := client.Set(key, value, 0).Err()
	return err
}

func Get(key string) (string, error) {
	client := RClient()
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("no value found in cache")
	}
	// fmt.Println(key, val)
	return val, err
}

func TestRedis(key string, value string) {
	err := Ping()
	if err != nil {
		fmt.Println(err)
	}

	err = Set(key, value)
	if err != nil {
		fmt.Println(err)
	}

	val, err := Get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(val, err)
}

/// GRPC