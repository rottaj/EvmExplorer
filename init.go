package main

import (
  "fmt"
  "log"
  "os"
  //"time"
  //"math/big"
  //"context"
  //"github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/joho/godotenv"

  "github.com/rottaj/GoEVMExplorer/eth"
  "github.com/rottaj/GoEVMExplorer/ui"
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

  block := eth.GetLatestBlock(client)
  fmt.Println("Block", block)
  //time.Sleep(12 * time.Second)
  viewer := new(ui.Viewer)
  viewer.Init()

}
