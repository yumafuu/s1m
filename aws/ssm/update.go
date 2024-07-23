package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// Update an SSM parameter
func Update(
	ctx context.Context,
	client *ssm.Client,
	param types.Parameter,
	newValue string,
) error {
	_, err := client.PutParameter(ctx, &ssm.PutParameterInput{
		Name:      param.Name,
		Value:     aws.String(newValue),
		Type:      param.Type,
		Overwrite: aws.Bool(true),
	})
	return err
}
