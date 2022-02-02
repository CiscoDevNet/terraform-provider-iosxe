## Add an ACL using Terraform
Use the [simplified_acl.tf](simplified_acl.tf) file to configure an ACL on a Catalyst 9300 switch. 

1. Ensure Terraform is installed
1. Configure RESTCONF on the switch 
    ```
    configure terminal
    restconf 
     ```
1. Initialize Terraform from a host
   ```
   terraform init
   ``` 
1. Apply the JSON within the terraform file
   ```
   terraform apply -auto-approve
   ```

See an example terraform file to configure an ACL
![](acl.gif)
