package util

import (
	"log"
	"os"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)


func InitializeClient() *ethclient.Client {
	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatal(err)
	}
	RPC_URL := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err) 
	}
	return client
}