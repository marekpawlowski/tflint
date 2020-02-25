// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloud9EnvironmentEc2InvalidOwnerArnRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloud9_environment_ec2" "foo" {
	owner_arn = "arn:aws:elasticbeanstalk:us-east-1:123456789012:environment/My App/MyEnvironment"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloud9EnvironmentEc2InvalidOwnerArnRule(),
					Message: `"arn:aws:elasticbeanstalk:us-east-1:123456789012:environment/My App/MyEnvironment" does not match valid pattern ^arn:aws:(iam|sts)::\d+:(root|(user\/[\w+=/:,.@-]{1,64}|federated-user\/[\w+=/:,.@-]{2,32}|assumed-role\/[\w+=:,.@-]{1,64}\/[\w+=,.@-]{1,64}))$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloud9_environment_ec2" "foo" {
	owner_arn = "arn:aws:iam::123456789012:user/David"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloud9EnvironmentEc2InvalidOwnerArnRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
