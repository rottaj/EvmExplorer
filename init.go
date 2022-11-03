package main

import (
  "fmt"
  "log"
  "os"
  //"context"
  //"time"
  //"math/big"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env") 
  if err != nil {
    log.Fatal(err)
  }
  RPC_URL := os.Getenv("RPC_URL")
  client, err := ethclient.Dial(RPC_URL)
  if err != nil {
    log.Fatal(err) 
  }
  fmt.Println("Client Connected")
  //latestBlock := GetLatestBlock(client)
  _ = client
  //time.Sleep(12 * time.Second)

}
