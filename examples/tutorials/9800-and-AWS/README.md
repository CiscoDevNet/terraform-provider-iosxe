## Create a 9800 instance in AWS using Terraform
Use the [create-9800-in-aws.tf](create-9800-in-aws.tf) file to create a Catalyst 9800 instance on an existing AWS account. 

1. Replace the 'XXXXXXXXXX' and 'YYYYYYYYYY' in the [.aws/credentials](./.aws/credentials.txt) file to align with your AWS account
1. Navigate to the 9800-and-AWS folder
1. Ensure Terraform is installed
1. Initialize Terraform from a host to use the AWS Terraform provder
   ```
   terraform init
   ``` 
1. Apply the JSON within the terraform file
   ```
   terraform apply -auto-approve
   ```
1. Now the device will be created in your AWS EC2 instances. This will take a few minutes
1. Configure RESTCONF on the 9800 
    ```
    configure terminal
    restconf 
     ```
1. Add a username and password to the 9800 to be able to reach it using the REST API. For example, the following will create a user with username = admin and password = pass123
    ```
    configure terminal
    username admin privilege 15 password 0 pass123
     ```
1. Navigate to the `configure-WLC` folder
1. Modify the "X.X.X.X" values in [configure-WLC/9800.tfvars](configure-WLC/9800.tfvars) to be the IP address of the new 9800.
1. Initialize Terraform from a host to use the IOS XE Terraform Provider
   ```
   terraform init
   ``` 
1. Apply the JSON within the terraform file
   ```
    terraform apply -auto-approve -var-file="9800.tfvars"
   ```