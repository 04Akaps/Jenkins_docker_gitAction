package controller

import (
	"context"

	"github.com/04Akaps/Jenkins_docker_go.git/crypto"
	connection "github.com/04Akaps/Jenkins_docker_go.git/mysql"
	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
)

type CommentController struct {
	Ctx         context.Context
	MySQLClient *sqlc.Queries
	EthClient   crypto.CryptoClientImpl
}

type CommentImpl interface{}

func NewCommentController() CommentImpl {
	context := context.Background()
	endPoint := "https://mainnet.infura.io/v3/299623e5cf3442c8bb2dbe870d8f7d88"

	client := crypto.NewEthClient(context, endPoint)

	return &CommentController{Ctx: context, MySQLClient: connection.NewMySQLClient("sns"), EthClient: client}
}
