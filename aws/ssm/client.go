package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type (
	Client struct {
		*ssm.Client
	}
	ParameterType = types.ParameterType
)

const (
	ParameterTypeSecureString = ParameterType(types.ParameterTypeSecureString)
	ParameterTypeString       = ParameterType(types.ParameterTypeString)
	ParameterTypeStringList   = ParameterType(types.ParameterTypeStringList)
)

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := ssm.NewFromConfig(cfg)

	return &Client{client}, nil
}
