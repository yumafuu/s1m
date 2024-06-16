package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func NewClient() (*ssm.Client, error) {

	// AWS コンフィグをロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// SSM クライアントを作成
	client := ssm.NewFromConfig(cfg)

	return client, nil
}
