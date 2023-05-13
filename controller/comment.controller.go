package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/04Akaps/Jenkins_docker_go.git/crypto"
	connection "github.com/04Akaps/Jenkins_docker_go.git/mysql"
	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	"github.com/04Akaps/Jenkins_docker_go.git/utils"
)

type CommentController struct {
	Ctx         context.Context
	MySQLClient *sqlc.Queries
	EthClient   crypto.CryptoClientImpl
}

type CommentImpl interface {
	CreateNewComment(http.ResponseWriter, *http.Request)
	// DeleteComment(http.ResponseWriter, *http.Request)
}

func NewCommentController() CommentImpl {
	context := context.Background()
	endPoint := "https://mainnet.infura.io/v3/299623e5cf3442c8bb2dbe870d8f7d88"

	client := crypto.NewEthClient(context, endPoint)

	return &CommentController{Ctx: context, MySQLClient: connection.NewMySQLClient("sns"), EthClient: client}
}

func (c *CommentController) CreateNewComment(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateNewCommentt")
	var req sqlc.CreateNewCommentParams

	decoder := utils.BodyDecoder(w, r)
	if err := decoder.Decode(&req); err != nil {
		log.Println("디코딩 실패")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 주소 검증
	if !c.EthClient.IsEoaAddress(req.CommentOwnerAccount) || c.EthClient.IsContractAddress(c.Ctx, req.CommentOwnerAccount) {
		log.Println("16진수가 아니고 40글자도 아닌경우 & Contract 주소 인 경우")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Text) == 0 {
		// 글 내용이 없는 경우
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.PostID <= 0 {
		// 외부키가 0인 경우
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := c.MySQLClient.GetPostId(c.Ctx, req.PostID)
	if err.Error() == "sql: no rows in result set" {
		// 해당하는 post_id가 없는 경우
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		// Get query에 실패한 경우
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = c.MySQLClient.CreateNewComment(c.Ctx, req)

	if err != nil {
		// Insert query에 실패한 경우
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
