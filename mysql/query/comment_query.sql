-- name: CreateNewComment :execresult
INSERT INTO comment (
    post_id,
    comment_owner_account,
    text
) VALUES (
   ?, ?, ?
);

-- name: DeleteComment :execresult
DELETE FROM comment 
WHERE comment_id = ? LIMIT 1;





--  update는 나중에 추가 예정