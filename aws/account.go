package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

func ValidAccount(ctx context.Context) error {
	_, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	return nil
}
