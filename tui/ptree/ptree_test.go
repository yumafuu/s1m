package ptree_test

import (
	"reflect"
	"testing"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/ptree"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/davecgh/go-spew/spew"
)

func Test_buildMapFromPaths(t *testing.T) {
	t.Parallel()

	test := struct {
		name     string
		params   []ssm.Parameter
		expected ptree.Node
	}{
		name: "success",
		params: []ssm.Parameter{
			{
				Name:  aws.String("/path1/key1"),
				Value: aws.String("value1-1"),
			},
			{
				Name:  aws.String("/path1/key2"),
				Value: aws.String("value2-2"),
			},
			{
				Name:  aws.String("/path2/key1"),
				Value: aws.String("value2-1"),
			},
		},
		expected: ptree.Node{
			"path1": ptree.Node{
				"key1": ssm.Parameter{
					Name:  aws.String("/path1/key1"),
					Value: aws.String("value1-1"),
				},
				"key2": ssm.Parameter{
					Name:  aws.String("/path1/key2"),
					Value: aws.String("value2-2"),
				},
			},
			"path2": ptree.Node{
				"key1": ssm.Parameter{
					Name:  aws.String("/path2/key1"),
					Value: aws.String("value2-1"),
				},
			},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		want := test.expected
		actual := ptree.BuildMapFromPaths(test.params)
		spew.Dump(actual)

		t.Logf("[INFO] want: %v, actual: %v", want, actual)

		if !reflect.DeepEqual(want, actual) {
			t.Errorf("[FAILED] want: %v, actual: %v", want, actual)
		}
	})
}
