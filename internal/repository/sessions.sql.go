// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: sessions.sql

package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSession, id)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, created_at, expires_at, persistent FROM sessions WHERE id = $1 AND expires_at > NOW()
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRow(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.Persistent,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT u.id, username, role, password, s.id, user_id, created_at, expires_at, persistent 
FROM users u
JOIN sessions s ON u.id = s.user_id 
WHERE s.id = $1
`

type GetUserBySessionRow struct {
	ID         int32
	Username   string
	Role       Role
	Password   []byte
	ID_2       uuid.UUID
	UserID     int32
	CreatedAt  time.Time
	ExpiresAt  time.Time
	Persistent bool
}

func (q *Queries) GetUserBySession(ctx context.Context, id uuid.UUID) (GetUserBySessionRow, error) {
	row := q.db.QueryRow(ctx, getUserBySession, id)
	var i GetUserBySessionRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.Password,
		&i.ID_2,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.Persistent,
	)
	return i, err
}

const insertSession = `-- name: InsertSession :one
INSERT INTO sessions (id, user_id, expires_at, persistent)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, created_at, expires_at, persistent
`

type InsertSessionParams struct {
	ID         uuid.UUID
	UserID     int32
	ExpiresAt  time.Time
	Persistent bool
}

func (q *Queries) InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, insertSession,
		arg.ID,
		arg.UserID,
		arg.ExpiresAt,
		arg.Persistent,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.Persistent,
	)
	return i, err
}

const updateSessionExpiry = `-- name: UpdateSessionExpiry :exec
UPDATE sessions SET expires_at = $2 WHERE id = $1
`

type UpdateSessionExpiryParams struct {
	ID        uuid.UUID
	ExpiresAt time.Time
}

func (q *Queries) UpdateSessionExpiry(ctx context.Context, arg UpdateSessionExpiryParams) error {
	_, err := q.db.Exec(ctx, updateSessionExpiry, arg.ID, arg.ExpiresAt)
	return err
}
