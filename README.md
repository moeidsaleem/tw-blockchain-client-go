#  Blockchain Go Client

This repository contains a simple Go-based blockchain client for the Polygon network. It provides an easy-to-use command-line interface to interact with the Polygon network using JSON-RPC API.



## Overview 

The Go source code is located in the root folder of the repository. The `Dockerfile` is used to create a Docker image of the application. The Terraform configuration in the `/terraform` folder can be used to deploy the application to AWS ECS Fargate.

## Requirements

Go 1.16 or later (1.20.2)
Docker
Terraform
An AWS account

## Usage

1. Clone the repository:
``` 
git clone https://github.com/moeidsaleem/tw-blockchain-client-go.git
cd tw-blockchain-client-go
```

2. Run the Go code:
```
go run main.go
```
This will print the current block number and the block with number 0x134e82a to the console.


## Quickstart

Bash scripts are available for setting up things quickly. 
- Run the code
```
sh run.sh
```

- Build an exectuable `client` 
```
sh build.sh
```
- Build image 
```
sh build_image.sh
```


3. Build the Docker image:
```
docker build -t blockchain_client .
```
This will build a Docker image of the application with the specified name.


4. Run the Docker container:

```
docker run --rm -it <image-name>
```
This will run a Docker container of the application.




## Installation
1. Clone the repository:
```
git clone https://github.com/moeidsaleem/tw-blockchain-client-python.git
cd tw-blockchain-client-python
```

2. Create a Python virtual environment and activate it:
``` 
python3 -m venv venv
source venv/bin/activate 
```

3. Install the required Python packages:
``` 
pip install -r app/requirements.txt
```

4. Run the application locally:
``` 
python app/main.py
```

## Building the Docker Image
To build a Docker image for the application, run the following command from the root of the repository:

``` 
docker build -t blockchain-client .
```

To run the Docker container locally:

```
docker run --rm -it blockchain-client
```

## Deploying to AWS ECS Fargate using Terraform

Please setup your AWS environment variables or update it in deploy.sh. 

I have setup `terraform/deploy.sh` file that will do complete terraform process i.e. from initilization to setting up example variables and deployment. 
``` bash
sh deploy.sh
```


### Deployment to AWS ECS Fargate using Terraform
To deploy the application to AWS ECS Fargate, follow these steps:

1. Navigate to the terraform directory:
`cd terraform`


2. Initialize the Terraform working directory:

```
terraform init
```

3. Plan the Terraform changes:
```
terraform plan
```

Apply the Terraform changes:
```
terraform apply
```

This will create the necessary AWS resources and deploy the application to AWS ECS Fargate.



##  Author

<img src="https://yt3.googleusercontent.com/LN0J3J7S-3QBM6LcjE6C43O7sG9UOW38srqPQAlovovNi_xBjqo4MqSmvlpCzffXbAUwZVR2c50=s900-c-k-c0x00ffffff-no-rj" width="120" height="120" style="border-radius:300px" />

Moeid Saleem Khan (Mo) 🚀