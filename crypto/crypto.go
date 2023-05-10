package crypto

import (
	"context"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CryptoClient struct {
	Client *ethclient.Client
}

type CryptoClientImpl interface {
	IsContractAddress(context.Context, string) bool
	IsEoaAddress(address string) bool
}

var re = regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // 40자리의 16진수 인지 검증

func NewEthClient(ctx context.Context, uri string) CryptoClientImpl {
	client, err := ethclient.DialContext(ctx, uri)
	if err != nil {
		panic("Eth Client Create Failed")
	}
	return &CryptoClient{Client: client}
}

func (c *CryptoClient) IsContractAddress(ctx context.Context, adr string) bool {
	address := common.HexToAddress(adr)

	bytecode, err := c.Client.CodeAt(ctx, address, nil) // nill is latest block
	if err != nil {
		return false
	}

	return len(bytecode) > 0
}

func (*CryptoClient) IsEoaAddress(address string) bool {
	return re.MatchString(address)
}
