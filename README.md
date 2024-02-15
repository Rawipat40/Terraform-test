# Terraform Testing Example with Terratest

This repository contains an example of testing Terraform configurations using Terratest, a Go library for automated infrastructure testing.

## Overview

The code in this repository demonstrates how to write automated tests for Terraform configurations using Terratest. It includes a simple Terraform configuration that deploys a basic "Hello, World!" web server on AWS, and a corresponding test written in Go using Terratest.

## Prerequisites

To run the tests in this repository, you'll need the following installed on your machine:

- Go (version 1.14 or higher)
- Terraform (version 0.12.x)
- AWS CLI configured with necessary permissions

## Install Terraform

Official installation instructions from HashiCorp: https://learn.hashicorp.com/tutorials/terraform/install-cli

## AWS Account Setup

AWS Terraform provider documentation: https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication

1) create non-root AWS user
2) Add the necessary IAM roles (e.g. AmazonEC2FullAccess)
3) Save Access key + secret key (or use AWS CLI `aws configure` -- https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)

## Hello World

`./main.tf` contains minimal configuration to provision an EC2 instance.

1. `aws configure`
   - **Purpose**: This command is used to configure your AWS CLI (Command Line Interface) with your AWS access credentials (Access Key ID and Secret Access Key), default region, and output format.
   - **Usage**: After installing the AWS CLI, you can run this command in your terminal or command prompt. It will prompt you to enter your AWS access key ID, secret access key, default region, and output format.
   - **Example**: 
     ```bash
     aws configure
     ```

2. `terraform init`
   - **Purpose**: This command initializes a Terraform working directory by downloading modules and providers specified in your Terraform configuration files (e.g., `main.tf`).
   - **Usage**: Run this command in the directory containing your Terraform configuration files.
   - **Example**:
     ```bash
     terraform init
     ```

3. `terraform plan`
   - **Purpose**: This command generates an execution plan showing what actions Terraform will take to change the infrastructure as defined in your Terraform configuration files.
   - **Usage**: After initializing your Terraform workspace with `terraform init`, you can run this command to see the planned changes before applying them.
   - **Example**:
     ```bash
     terraform plan
     ```

4. `terraform apply`
   - **Purpose**: This command applies the changes required to reach the desired state of the configuration as defined in your Terraform configuration files.
   - **Usage**: After running `terraform plan` to review the planned changes, you can apply those changes by running this command.
   - **Example**:
     ```bash
     terraform apply
     ```

These commands are essential for working with Terraform to manage infrastructure as code (IaC) on AWS. They allow you to configure your AWS credentials, initialize your Terraform workspace, plan infrastructure changes, and apply those changes to your AWS account.


## Understanding the Code

The main components of this repository are:

- `main.tf`: Terraform configuration file defining the infrastructure to deploy.
- `test/main_test.go`: Go test file containing the Terratest test function `TestTerraformHelloWorldExample`.
- `README.md`: This README file providing an overview of the repository and instructions for usage.

## Go Packages Used

- `github.com/gruntwork-io/terratest/modules/http-helper`: Terratest module for HTTP testing.
- `github.com/gruntwork-io/terratest/modules/terraform`: Terratest module for testing Terraform configurations.

## Test Function Explanation

The `TestTerraformHelloWorldExample` test function does the following:

1. Sets up Terraform options with default retryable errors.
2. Defers the destruction of provisioned resources until the test function exits.
3. Initializes and applies the Terraform configuration.
4. Retrieves the URL output variable from Terraform.
5. Makes an HTTP GET request to the provisioned instance with retries and custom validation.

## Credits
This code was part of [Complete Terraform Course](https://www.youtube.com/watch?v=7xngnjfIlK4)
