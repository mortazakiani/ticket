// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	// users.sql
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
}

var _ Querier = (*Queries)(nil)
