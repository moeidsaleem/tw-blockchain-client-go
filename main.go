package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

const rpcEndpoint = "https://polygon-rpc.com/"

type rpcRequest struct {
    Jsonrpc string        `json:"jsonrpc"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
    Id      int           `json:"id"`
}

type rpcResponse struct {
    Jsonrpc string      `json:"jsonrpc"`
    Id      int         `json:"id"`
    Result  interface{} `json:"result"`
}

func makeRpcRequest(method string, params []interface{}) (*rpcResponse, error) {
    request := &rpcRequest{
        Jsonrpc: "2.0",
        Method:  method,
        Params:  params,
        Id:      1,
    }
    requestBody, err := json.Marshal(request)
    if err != nil {
        return nil, err
    }
    response, err := http.Post(rpcEndpoint, "application/json", bytes.NewReader(requestBody))
    if err != nil {
        return nil, err
    }
    var rpcResp rpcResponse
    err = json.NewDecoder(response.Body).Decode(&rpcResp)
    if err != nil {
        return nil, err
    }
    return &rpcResp, nil
}

func getBlockNumber() (string, error) {
    resp, err := makeRpcRequest("eth_blockNumber", nil)
    if err != nil {
        return "", err
    }
    blockNumber, ok := resp.Result.(string)
    if !ok {
        return "", fmt.Errorf("unexpected result format: %#v", resp.Result)
    }
    return blockNumber, nil
}

func getBlockByNumber(blockNumber string, withTx bool) (interface{}, error) {
    params := []interface{}{blockNumber, withTx}
    resp, err := makeRpcRequest("eth_getBlockByNumber", params)
    if err != nil {
        return nil, err
    }
    return resp.Result, nil
}

func main() {
    blockNumber, err := getBlockNumber()
    if err != nil {
        panic(err)
    }
    fmt.Println("Current block number:", blockNumber)

    block, err := getBlockByNumber("0x134e82a", true)
    if err != nil {
        panic(err)
    }
    fmt.Println("Block 0x134e82a:", block)
}
