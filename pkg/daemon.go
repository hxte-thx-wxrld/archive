package htwarchive

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgxlisten"
)

func daemon(b chan bool) {
	i := 0
	for {
		if i > 100 {
			break
		}
		//fmt.Println("Daemon Test")
		time.Sleep(5 * time.Second)
		i++
	}

	b <- true
}

func postgresListener() {
	//testchannel := make(chan *pgconn.Notification)

	listener := &pgxlisten.Listener{
		Connect: func(ctx context.Context) (*pgx.Conn, error) {
			return pgx.Connect(ctx, os.Getenv("PG_URL"))
		},
	}

	listener.Handle("testchannel", pgxlisten.HandlerFunc(func(ctx context.Context, notification *pgconn.Notification, conn *pgx.Conn) error {
		fmt.Println("message on testchannel", notification)
		//testchannel <- notification
		return nil
	}))

	listener.Handle("cancel", pgxlisten.HandlerFunc(func(ctx context.Context, notification *pgconn.Notification, conn *pgx.Conn) error {
		fmt.Println("bye")
		time.Sleep(5 * time.Second)
		return nil
	}))

	go func() {
		listener.Listen(context.Background())
	}()
}

func DefaultDaemon() chan bool {
	done := make(chan bool)
	postgresListener()
	go daemon(done)
	return done
}
