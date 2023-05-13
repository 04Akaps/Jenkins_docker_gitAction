// Code generated by sqlc. DO NOT EDIT.

package sns_mysql

import (
	"database/sql"
)

type Comment struct {
	CommentID           int64        `json:"comment_id"`
	PostID              int64        `json:"post_id"`
	CommentOwnerAccount string       `json:"comment_owner_account"`
	Text                string       `json:"text"`
	CreatedAt           sql.NullTime `json:"created_at"`
}

type Post struct {
	PostID           int64         `json:"post_id"`
	PostOwnerAccount string        `json:"post_owner_account"`
	Title            string        `json:"title"`
	ImageUrl         string        `json:"image_url"`
	Text             string        `json:"text"`
	LikePoint        sql.NullInt64 `json:"like_point"`
	CreatedAt        sql.NullTime  `json:"created_at"`
}
