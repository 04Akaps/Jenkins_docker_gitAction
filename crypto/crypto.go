package crypto

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CryptoClient struct {
	Client *ethclient.Client
}

func NewEthClient(ctx context.Context, uri string) *CryptoClient {
	client, err := ethclient.DialContext(ctx, uri)
	if err != nil {
		panic("Eth Client Create Failed")
	}
	return &CryptoClient{Client: client}
}

func (c *CryptoClient) IsAddress(ctx context.Context, adr string) bool {
	address := common.HexToAddress(adr)

	bytecode, err := c.Client.CodeAt(ctx, address, nil) // nill is latest block
	if err != nil {
		return false
	}

	isContract := len(bytecode) > 0

	return isContract
}
