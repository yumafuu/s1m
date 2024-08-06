package ssm

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type Parameter = types.Parameter

func (c Client) List(prefix string) ([]Parameter, error) {
	ctx := context.TODO()

	params, err := c.getParametersByPath(ctx, prefix)
	if err != nil {
		return []Parameter{}, err
	}

	return params, nil
}

func (c Client) getParametersByPath(
	ctx context.Context,
	path string,
) ([]Parameter, error) {
	var params []Parameter
	input := &ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	paginator := ssm.NewGetParametersByPathPaginator(c, input)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		params = append(params, output.Parameters...)
	}

	// sort params by name
	sort.Slice(params, func(i, j int) bool {
		return *params[i].Name < *params[j].Name
	})

	return params, nil
}
