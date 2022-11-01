// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id,refresh_token, user_agent, ip, is_valid, expires_at) 
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, "user", refresh_token, user_agent, ip, is_valid, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID      `json:"id"`
	RefreshToken string         `json:"refresh_token"`
	UserAgent    sql.NullString `json:"user_agent"`
	Ip           sql.NullString `json:"ip"`
	IsValid      bool           `json:"is_valid"`
	ExpiresAt    time.Time      `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.queryRow(ctx, q.createSessionStmt, createSession,
		arg.ID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.Ip,
		arg.IsValid,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.User,
		&i.RefreshToken,
		&i.UserAgent,
		&i.Ip,
		&i.IsValid,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, "user", refresh_token, user_agent, ip, is_valid, expires_at, created_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.queryRow(ctx, q.getSessionStmt, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.User,
		&i.RefreshToken,
		&i.UserAgent,
		&i.Ip,
		&i.IsValid,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
