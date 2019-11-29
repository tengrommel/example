# Main APIs

User Profile
 
- By using this application, user will track all this expanses, income etc
- User can create different accounts and transaction. Can create budget and schedule income/payments.

## API
- Create user -Email, Social Login(Facebook, Google)
- Login
- Get User
- List Users(Admin)
- Delete user(Target user can delete himself/Admin can delete any user) -Do we save user data?

## Technologies
This application will have 2 client apps Web and Android.For storing and data management, we will create backend using Golang and Postgresql.

Backend will be deployed to AWS or Heroku(depend on final cost, need additional research)

Deployment will be done by Docker and if we will choose AWS,environment setup will be done by Terraform.

All code will be on Bitbucket.

## Features
- User can create profile and login to the app, edit user info
    - Profile can be created using user email,facebook or google
- User can create financial accounts inside application. Accounts can be different types.(Ex: cheque/saving, credit card, etc)
- User can create transaction in each account, transaction can be income, expanses or transfer to other account.
- User can see total amount of money on each account and trends