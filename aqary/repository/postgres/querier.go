// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package postgres

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetUser(ctx context.Context, phoneNumber string) (Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
