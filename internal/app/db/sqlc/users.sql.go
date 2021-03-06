// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, is_mentor, status)
VALUES ($1, $2, $3)
RETURNING id, phone, email, created_at, status, full_name, is_mentor, user_name
`

type CreateUserParams struct {
	ID       string        `json:"id"`
	IsMentor sql.NullInt32 `json:"is_mentor"`
	Status   int32         `json:"status"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.IsMentor, arg.Status)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getMentorAvailable = `-- name: GetMentorAvailable :one
SELECT id, phone, email, created_at, status, full_name, is_mentor, user_name
FROM users u
WHERE u.is_mentor = 1
AND u.status = 1
AND not exists (SELECT 1 FROM conversation c WHERE c.mentor_id = u.id and c.status = 1 and c.valid_time > now())
LIMIT 1
`

func (q *Queries) GetMentorAvailable(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, getMentorAvailable)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, phone, email, created_at, status, full_name, is_mentor, user_name
FROM users
WHERE id = $1
    AND status = 1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}

const listAllUsers = `-- name: ListAllUsers :many
SELECT id, phone, email, created_at, status, full_name, is_mentor, user_name
FROM users
where status = 1
ORDER BY id
`

func (q *Queries) ListAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.Email,
			&i.CreatedAt,
			&i.Status,
			&i.FullName,
			&i.IsMentor,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, phone, email, created_at, status, full_name, is_mentor, user_name
FROM users
WHERE status = 1
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.Email,
			&i.CreatedAt,
			&i.Status,
			&i.FullName,
			&i.IsMentor,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserEmail = `-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
RETURNING id, phone, email, created_at, status, full_name, is_mentor, user_name
`

type UpdateUserEmailParams struct {
	ID    string         `json:"id"`
	Email sql.NullString `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}

const updateUserPhone = `-- name: UpdateUserPhone :one
UPDATE users
SET phone = $2
WHERE id = $1
RETURNING id, phone, email, created_at, status, full_name, is_mentor, user_name
`

type UpdateUserPhoneParams struct {
	ID    string         `json:"id"`
	Phone sql.NullString `json:"phone"`
}

func (q *Queries) UpdateUserPhone(ctx context.Context, arg UpdateUserPhoneParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPhone, arg.ID, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}

const updateUserStatus = `-- name: UpdateUserStatus :one
UPDATE users
SET STATUS = $2
WHERE id = $1
RETURNING id, phone, email, created_at, status, full_name, is_mentor, user_name
`

type UpdateUserStatusParams struct {
	ID     string `json:"id"`
	Status int32  `json:"status"`
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserStatus, arg.ID, arg.Status)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Email,
		&i.CreatedAt,
		&i.Status,
		&i.FullName,
		&i.IsMentor,
		&i.UserName,
	)
	return i, err
}
