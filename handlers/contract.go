package handlers

import (
	"context"
	_ "log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetValue(client *ethclient.Client, contractAddress string) (*big.Int, error) {
	address := common.HexToAddress(contractAddress)
	instance, err := NewMyContract(address, client)
	if err != nil {
		return nil, err
	}

	value, err := instance.GetValue(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, err
	}

	return value, nil
}
