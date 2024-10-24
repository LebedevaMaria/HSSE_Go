package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Input struct {
	Input string `json:"input"`
}

type Output struct {
	Output string `json:"output"`
}

type Client struct {
	Http string
}

func (client Client) GetVersion() (string, error) {
	resp, err := http.Get(client.Http + "/version")
	if err != nil {
		fmt.Printf("Failed to get version: %v", err)
	}
	defer resp.Body.Close()
	ans, _ := io.ReadAll(resp.Body)
	return string(ans), nil
}

func (client Client) PostDecode(str string) (string, error) {
	type Input struct {
		Input string `json:"input"`
	}

	type Output struct {
		Output string `json:"output"`
	}

	input := Input{base64.StdEncoding.EncodeToString([]byte(str))}
	inputJSON, _ := json.Marshal(input)
	resp, err := http.Post(client.Http+"/decode", "application/json", bytes.NewBuffer(inputJSON))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var ans Output
	err = json.NewDecoder(resp.Body).Decode(&ans)
	return ans.Output, nil
}

func (client Client) GetHardOp() (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", client.Http+"/hard-op", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return false, 0, nil
		} else {
			return false, 0, err
		}
	}
	defer resp.Body.Close()
	return true, resp.StatusCode, nil
}

func NewClient(Http string) Client {
	current := Client{Http}
	return current
}
