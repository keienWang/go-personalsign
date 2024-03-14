package main

// /login
// https://kaidro.com/api/login
// {
// "walletAddress": "0x8803eD3df8Aa36493f43292729F5325dbc629C1f"
// }

// 签名
// {
// "walletAddress": "0x8803eD3df8Aa36493f43292729F5325dbc629C1f",
// "signature": "0x32429c4e2c04a3f96253763524f4d3ddaae5ea130591ee446ad503e7ee682e0367336e6450649718c2ef7592e9be171e1190a5bfa7d54a5eb5728a452ceae9bd1c"
//
//	}

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Define a struct that matches the JSON payload structure
type Payload struct {
	WalletAddress string `json:"walletAddress"`
}

func main() {
	login()
	// Create the payload
	getSign()
}

func getSign() {

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	message := "hello"

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		fmt.Println("sign error!")
	}
	signatureBytes[64] += 27

	fmt.Println(hexutil.Encode(signatureBytes))
}

func login() {

	url := "https://kaidro.com/api/login"
	// 请求体（根据你的需要修改这个 JSON 字符串）
	jsonStr := []byte(`{"walletAddress": "0x8803eD3df8Aa36493f43292729F5325dbc629C1f"}`)

	// 创建一个新的请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 添加请求头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,ja;q=0.8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "_vcrcs=1.1710402739.3600.ZWJiMjBjN2Y5ZTdmNzhlNzExNWEzZGFhNjFhZGIyYzc=.5539eace41e768b7cf30b88333f2764c; accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYyMDMwLCJ3YWxsZXRBZGRyZXNzIjoiMHg4ODAzZUQzZGY4QWEzNjQ5M2Y0MzI5MjcyOUY1MzI1ZGJjNjI5QzFmIiwiaWF0IjoxNzEwNDAyODE5LCJleHAiOjE3MTA2NjIwMTl9.HtDE-qWbgEpOjpVA1ATmCn_tLZXnEO5xmofmS5NnH4g")
	req.Header.Set("Origin", "https://kaidro.com")
	req.Header.Set("Referer", "https://kaidro.com/")
	req.Header.Set("Sec-Ch-Ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}
