package main

import (
	"context"
	"log"

	"github.com/YumaFuu/s1m/aws"
	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui"
)

func main() {
	ctx := context.Background()

	err := aws.ValidAccount(ctx)
	if err != nil {
		panic(err)
	}

	client, err := ssm.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	a, err := tui.NewTui(client)
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
