package main

import (
	"context"
	"log"

	"github.com/YumaFuu/ssm-tui/app"
	"github.com/YumaFuu/ssm-tui/aws"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
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

	a, err := app.NewApp(client)
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
