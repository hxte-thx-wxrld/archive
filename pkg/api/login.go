package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/htw-archive/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserLookupResult struct {
	UserId          string
	Username        string
	Admin           bool
	AssignedArtists []model.Artist
}

type CheckPasswordResponse struct {
	UserId   string
	Username string
	Admin    bool
}

func AuthenticatedMiddleware(ctx *gin.Context) {

	session := sessions.Default(ctx)

	id := session.Get("userid")
	fmt.Println("Middleware", id)
	if id == nil {
		ctx.Status(http.StatusUnauthorized)
	} else {
		ctx.Keys["userid"] = id.(string)
		ctx.Next()
	}

}

func AdminMiddleware(ctx *gin.Context) {
	session := sessions.Default(ctx)
	admin := session.Get("admin").(bool)
	fmt.Println("Is Admin:", admin)
	if admin {
		ctx.Next()
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

func UpdateAdminAccess(db *pgxpool.Pool, password string) error {
	/*INSERT INTO inventory (id, name, price, quantity) VALUES (1, 'A', 16.99, 120)
	ON CONFLICT(id)
	DO UPDATE SET
	  price = EXCLUDED.price,
	  quantity = EXCLUDED.quantity;
		_, err := db.Exec(context.Background(), "")*/

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), "insert into users (username, password_hash, admin) values ('admin', @password, true) on conflict(username) do update set password_hash=@password", pgx.NamedArgs{
		"password": bytes,
	})
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetAssignedArtists(db *pgxpool.Pool, userid string) ([]model.Artist, error) {
	rows, err := db.Query(context.Background(), "select i.id as ArtistId, i.name from artists_of_user aou join interpret i on i.id = aou.artist_id where aou.user_id::text = @userId", pgx.NamedArgs{
		"userId": userid,
	})

	if err != nil {
		return nil, err
	}

	a := []model.Artist{}
	for rows.Next() {
		var artist model.Artist

		err := rows.Scan(&artist.ArtistId, &artist.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println(artist)

		a = append(a, artist)
	}
	return a, nil
}

func LookupUser(db *pgxpool.Pool, userid string) (*UserLookupResult, error) {
	row := db.QueryRow(context.Background(), "select username, admin from users where id = @userid", pgx.NamedArgs{
		"userid": userid,
	})

	var user UserLookupResult
	user.UserId = userid
	err := row.Scan(&user.Username, &user.Admin)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	aa, err := GetAssignedArtists(db, userid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.AssignedArtists = aa

	//fmt.Println(user)

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
