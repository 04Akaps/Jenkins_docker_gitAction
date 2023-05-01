-- name: CreateNewSnsPost :one
INSERT INTO post (
    eoa_address,
    contract_address,
    network_chain_id,
    price,
    meta_data_uri
) VALUES (
   ?, ?, ?, ?, ?
);

-- name: GetMyAllLaunchpad :many
SELECT * FROM evmLaunchpad
WHERE eoa_address = ?;

-- name: GetLaunchpad :one
SELECT * FROM evmLaunchpad
WHERE contract_address = ? LIMIT 1;

-- name: DeleteAllLaunchpad :execresult
DELETE FROM evmLaunchpad
