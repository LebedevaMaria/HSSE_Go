package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Input struct {
	Input string `json:"input"`
}

type Output struct {
	Output string `json:"output"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//GET version
	resp, err := http.Get("http://localhost:8080/version")
	if err != nil {
		fmt.Printf("Failed to get version: %v", err)
	}
	defer resp.Body.Close()
	ans, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", ans)

	//POST
	input := Input{base64.StdEncoding.EncodeToString([]byte("Roachhh"))}
	inputJSON, _ := json.Marshal(input)
	resp, err = http.Post("http://localhost:8080/decode", "application/json", bytes.NewBuffer(inputJSON))
	if err != nil {
		fmt.Printf("Failed: %v", err)
	}
	defer resp.Body.Close()
	ans, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", ans)

	//GET hard-op
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	if err != nil {
		fmt.Printf("Failed: %v", err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out after 15 seconds")
		} else {
			fmt.Println("Failed: %v", err)
		}
	} else {
		defer resp.Body.Close()
		ans, _ = ioutil.ReadAll(resp.Body)
		fmt.Printf("%d", resp.StatusCode)
	}
}
