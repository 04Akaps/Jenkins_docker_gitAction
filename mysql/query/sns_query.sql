-- name: CreateNewSnsPost :execresult
INSERT INTO post (
    post_owner_account,
    title,
    image_url,
    text
) VALUES (
   ?, ?, ?, ?
);

-- name: GetSnsPostAll :many
SELECT * FROM post
WHERE post_owner_account = ?;

-- name: GetSnsPost :one
SELECT * FROM post
WHERE post_id = ? LIMIT 1;

-- name: DeleteSnsPostByPostId :execresult
DELETE FROM post 
WHERE post_id = ? LIMIT 1;

-- name: GetPostId :one
SELECT post_id FROM post
WHERE post_id = ?;


-- update는 나중에 추가 예정