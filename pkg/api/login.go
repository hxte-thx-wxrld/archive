package api

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserLookupResult struct {
	UserId   string
	Username string
}

func LookupUser(db *pgxpool.Pool, userid string) (*UserLookupResult, error) {
	row := db.QueryRow(context.Background(), "select username from users where id = @userid", pgx.NamedArgs{
		"userid": userid,
	})

	var user UserLookupResult
	user.UserId = userid
	err := row.Scan(&user.Username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func CreateUser(db *pgxpool.Pool, username string, password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	row := db.QueryRow(context.Background(), "insert into users (username, password_hash) values (@username, @password) returning id::text", pgx.NamedArgs{
		"username": username,
		"password": bytes,
	})

	var id string
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func CheckPassword(db *pgxpool.Pool, username string, password string) (*string, error) {
	row := db.QueryRow(context.Background(), "select id, password_hash from users where username = @username", pgx.NamedArgs{
		"username": username,
	})

	var id string
	var hash []byte
	err := row.Scan(&id, &hash)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	} else {
		return &id, nil
	}

}
