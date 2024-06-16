package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

func AwsAccountName() string {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return ""
	}

	return cfg.AppID
}
