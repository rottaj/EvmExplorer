package eth

import (
  _"fmt"
  "log"
  "context"
  "math/big"
  "github.com/ethereum/go-ethereum/core/types"
)

type Client interface {
  BlockNumber(ctx context.Context) (uint64, error)
  BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

func GetLatestBlock(cl Client) *types.Block {
  blockNumber, err := cl.BlockNumber(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  block, err := cl.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
  if err != nil {
    log.Fatal(err)
  }
  _ = block
  //fmt.Println("BLOCK", block.Time())
  return block 
}

