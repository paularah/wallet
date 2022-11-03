// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id, user_id, refresh_token, user_agent, ip_address, is_valid, expires_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, user_id, refresh_token, user_agent, ip_address, location, is_valid, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	IpAddress    string    `json:"ip_address"`
	IsValid      bool      `json:"is_valid"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.queryRow(ctx, q.createSessionStmt, createSession,
		arg.ID,
		arg.UserID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.IpAddress,
		arg.IsValid,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.IpAddress,
		&i.Location,
		&i.IsValid,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, refresh_token, user_agent, ip_address, location, is_valid, expires_at, created_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.queryRow(ctx, q.getSessionStmt, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.IpAddress,
		&i.Location,
		&i.IsValid,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
