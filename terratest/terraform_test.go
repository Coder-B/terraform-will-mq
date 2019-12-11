package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformBasicExampleNew(t *testing.T) {
	terraformOptions := &terraform.Options {
		// The path to where your Terraform code is located
		TerraformDir: "../",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": "tf-example-ons-instance",
			"remark" = "tf-example-ons-instance-remark",
		},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}
	  
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)
	  
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
	  
	// Validate your code works as expected
	validateServerIsWorking(t, terraformOptions)

	mqInstanceId := terraform.Output(t, terraformOptions, "id")
	assert.NotNil(t, mqInstanceId)
	
}