package utils

import (
	"context"
	"fmt"

	"github.com/dogslee/genid"
	"github.com/go-redis/redis/v8"
)

func Generate() {
	cli := redis.NewClient(
		&redis.Options{
			Addr: "127.0.0.1:16379",
		},
	)
	g, err := genid.New(
		genid.DB(15),
		genid.Cli(cli),
	)
	if err != nil {
		panic(err)
	}
	code, _ := g.Create(context.Background(), "1000")
	fmt.Println(code)
}
