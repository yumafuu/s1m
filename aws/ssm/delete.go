package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func (c Client) Delete(
	name *string,
) error {
	ctx := context.TODO()

	_, err := c.DeleteParameter(ctx, &ssm.DeleteParameterInput{
		Name: name,
	})
	return err
}
