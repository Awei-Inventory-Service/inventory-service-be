# How to Run 

1. Run go mod tidy 
2. Copy .env.example and make new .env file
2. Build your docker image 
```
  docker-compose up -d --build
```

# How to connect to the AWS EC2 Instance

## How to connect using ssh
1. Download pem file yang namanya `admin-ssh-key.pem`
2. Connect ke dalem AWS instancenya
```sh
ssh -i "admin-ssh-keypair.pem" ec2-user@ec2-13-251-156-6.ap-southeast-1.compute.amazonaws.com
```

## How to copy file inside the instance
```sh
scp -i "admin-ssh-keypair.pem" filename ec2-user@ec2-13-251-156-6.ap-southeast-1.compute.amazonaws.com:~/
```
