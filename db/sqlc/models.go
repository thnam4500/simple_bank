// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be positive or negative
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username         string       `json:"username"`
	HashPassword     string       `json:"hash_password"`
	FullName         string       `json:"full_name"`
	Email            string       `json:"email"`
	PasswordChangeAt time.Time    `json:"password_change_at"`
	CreatedAt        sql.NullTime `json:"created_at"`
}
