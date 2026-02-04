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
		Addr: "redis-11941.re-cluster1.ps-redislabs.org:11941",
	})
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to replica-db")

	fmt.Println("Values in reverse order:")
	for i := 100; i >= 1; i-- {
		key := fmt.Sprintf("%d", i)
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s ", val)
		if (100-i+1)%10 == 0 {
			fmt.Println()
		}
	}
}
