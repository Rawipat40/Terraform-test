package test

import (
	"crypto/tls" // Importing TLS package for secure communication
	"fmt"
	"testing" // Importing testing package for writing tests
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	// Setting up Terraform options with default retryable errors
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/hello-world", // Specifying the directory containing Terraform configuration files
	})

	defer terraform.Destroy(t, terraformOptions) // Deferring the destruction of provisioned resources until the test function exits

	terraform.InitAndApply(t, terraformOptions) // Initializing and applying Terraform configuration

	instanceURL := terraform.Output(t, terraformOptions, "url") // Retrieving the value of the 'url' output variable from Terraform

	tlsConfig := tls.Config{} // Creating an empty TLS configuration struct for HTTP requests

	maxRetries := 30                       // Maximum number of retries for HTTP requests
	timeBetweenRetries := 10 * time.Second // Time interval between retries for HTTP requests

	// Making an HTTP GET request with retries and custom validation
	http_helper.HttpGetWithRetryWithCustomValidation(
		t, instanceURL, &tlsConfig, maxRetries, timeBetweenRetries, validate,
	)
}

func validate(status int, body string) bool {
	fmt.Println(body)
	return status == 200 // Returning true if the HTTP status code is 200 (OK), otherwise false
}
