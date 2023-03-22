

- Using Go, implement a simple blockchain client: 
"
HOST: https://polygon-rpc.com/
  
>>> Get block number

POST
{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "id": 2
}

>>> Get block by number

POST
{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": [
    "0x134e82a",
    true
  ],
  "id": 2
}
"
- Add Dockerfile to pack the client into the deployable application
- Create a /terraform folder and write HCL to deploy the application to AWS ECS Fargate (terraform local state, no need to deploy)  
- Add README with documentation, specify what could be added to ensure the application is production ready


