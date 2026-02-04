package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis-10974.re-cluster1.ps-redislabs.org:10974",
	})
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to source-db")

	rdb.Del(ctx, "numbers")

	values := make([]interface{}, 100)
	for i := 1; i <= 100; i++ {
		values[i-1] = i
	}

	err = rdb.RPush(ctx, "numbers", values...).Err()
	if err != nil {
		log.Fatal(err)
	}

	count, _ := rdb.LLen(ctx, "numbers").Result()
	fmt.Printf("Inserted %d values into source-db\n", count)
}
