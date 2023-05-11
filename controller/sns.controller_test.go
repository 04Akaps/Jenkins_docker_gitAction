package controller

import (
	"testing"

	sns_mock "github.com/04Akaps/Jenkins_docker_go.git/mock/sns_mock"
	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	"github.com/golang/mock/gomock"
)

// mockgen -source=./mock/sns/sns_doer.go -destination=./mock/sns/sns_gomock.go -package=mock
func TestSnsController(t *testing.T) {
	// log.Println(" ----------- Test Sns Controller ----------- ")
	// log.Println(" ----------- Post Make New Sns Controller ----------- ")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockInterface := sns_mock.NewMockISnsMockInterface(mockCtrl)
	testSns := &sns_mock.SnsMock{IMock: mockInterface}

	newSnsPost := &sqlc.CreateNewSnsPostParams{
		PostOwnerAccount: "0x365a00A8211b7a21aD7f74c2705092c50f0adfe1",
		Title:            "New Test Post Title",
		ImageUrl:         "New Test Image",
		Text:             "New Test Text",
	}
	mockInterface.EXPECT().NewPost(newSnsPost).Return(nil).Times(1)
	testSns.Use(newSnsPost)
}
