package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get the latest block number: %v", err)
	}
	fmt.Printf("Latest block number: %d\n", blockNumber)

	address := common.HexToAddress(ADDRESS)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get the balance: %v", err)
	}
	fmt.Printf("Balance: %s\n", balance.String())

	etherValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	fmt.Printf("Balance in Ether: %f\n", etherValue)
}
