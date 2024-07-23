package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type Parameter = types.Parameter

func List(client *ssm.Client, prefix string) ([]types.Parameter, error) {
	params, err := getParametersByPath(client, prefix)
	if err != nil {
		return []types.Parameter{}, err
	}

	return params, nil
}

func getParametersByPath(client *ssm.Client, path string) ([]types.Parameter, error) {
	var params []types.Parameter
	input := &ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	paginator := ssm.NewGetParametersByPathPaginator(client, input)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		params = append(params, output.Parameters...)
	}

	return params, nil
}
