package model

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId          string
	Username        string
	Admin           bool
	AssignedArtists []Artist
}

type CheckPasswordResponse struct {
	UserId   string
	Username string
	Admin    bool
}

func (u *User) GetAssignedArtists(db *pgxpool.Pool, userid string) error {
	rows, err := db.Query(context.Background(), "select i.id as ArtistId, i.name, i.artist_picture from artists_of_user aou join interpret i on i.id = aou.artist_id where aou.user_id::text = @userId", pgx.NamedArgs{
		"userId": userid,
	})

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var artist Artist
		artist.fromRow(rows)

		err := rows.Scan(&artist.ArtistId, &artist.Name)
		if err != nil {
			log.Println(err)
			return err
		}
		u.AssignedArtists = append(u.AssignedArtists, artist)
	}

	return nil
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

func CheckPassword(db *pgxpool.Pool, username string, password string) (*CheckPasswordResponse, error) {
	row := db.QueryRow(context.Background(), "select id, admin, password_hash from users where username = @username", pgx.NamedArgs{
		"username": username,
	})

	var u CheckPasswordResponse
	var hash []byte
	err := row.Scan(&u.UserId, &u.Admin, &hash)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	} else {
		u.Username = username
		return &u, nil
	}

}
