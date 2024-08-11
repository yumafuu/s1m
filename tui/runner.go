package tui

import (
	"context"

	"github.com/YumaFuu/s1m/aws"
	"github.com/YumaFuu/s1m/aws/ssm"
)

func Run() {
	ctx := context.Background()

	if err := aws.ValidAccount(ctx); err != nil {
		panic(err)
	}

	client, err := ssm.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	a, err := NewTui(client)
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
