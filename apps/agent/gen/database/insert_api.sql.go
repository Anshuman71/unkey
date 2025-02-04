// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: insert_api.sql

package database

import (
	"context"
	"database/sql"
)

const insertApi = `-- name: InsertApi :exec
INSERT INTO
    ` + "`" + `apis` + "`" + ` (
        id,
        workspace_id,
        name,
        ip_whitelist,
        auth_type,
        key_auth_id
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
`

type InsertApiParams struct {
	ID          string
	WorkspaceID string
	Name        string
	IpWhitelist sql.NullString
	AuthType    NullApisAuthType
	KeyAuthID   sql.NullString
}

func (q *Queries) InsertApi(ctx context.Context, arg InsertApiParams) error {
	_, err := q.db.ExecContext(ctx, insertApi,
		arg.ID,
		arg.WorkspaceID,
		arg.Name,
		arg.IpWhitelist,
		arg.AuthType,
		arg.KeyAuthID,
	)
	return err
}
