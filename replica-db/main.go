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

	values, err := rdb.LRange(ctx, "numbers", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Values in reverse order:")
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Printf("%s ", values[i])
		if (len(values)-i)%10 == 0 {
			fmt.Println()
		}
	}
}
