package main

import (
    "encoding/json"
    "fmt"
    "io"
    "math/big"
    "net/http"
    "os"
    "strconv"
    "time"
)

type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Result  string `json:"result"`
}

func main() {
    if len(os.Args) < 3 || len(os.Args) > 4 {
        fmt.Println("Usage: go run main.go <address> <apikey> [<chainid>]")
        os.Exit(1)
    }

    address := os.Args[1]
    apikey := os.Args[2]
    chainid := 1
    if len(os.Args) == 4 {
        var err error
        chainid, err = strconv.Atoi(os.Args[3])
        if err != nil {
            fmt.Printf("Invalid chainid: %v\n", err)
            os.Exit(1)
        }
    }

    for {
        url := fmt.Sprintf("https://api.etherscan.io/v2/api?chainid=%d&module=account&action=balance&address=%s&tag=latest&apikey=%s", chainid, address, apikey)

        resp, err := http.Get(url)
        if err != nil {
            fmt.Printf("Error fetching API: %v\n", err)
            time.Sleep(120 * time.Second)
            continue
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Printf("Error reading response: %v\n", err)
            time.Sleep(120 * time.Second)
            continue
        }

        var data Response
        err = json.Unmarshal(body, &data)
        if err != nil {
            fmt.Printf("Error parsing JSON: %v\n", err)
            time.Sleep(120 * time.Second)
            continue
        }

        if data.Status == "1" {
            wei := new(big.Int)
            _, ok := wei.SetString(data.Result, 10)
            if !ok {
                fmt.Printf("Error parsing balance: %s\n", data.Result)
                time.Sleep(120 * time.Second)
                continue
            }

            ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
            fmt.Printf("Balance: %s ETH\n", ether.Text('f', 18))
        } else {
            fmt.Printf("API error: %s\n", data.Message)
        }

        time.Sleep(120 * time.Second)
    }
}
