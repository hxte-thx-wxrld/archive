package htwarchive

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgxlisten"
)

func runScript(scripts *embed.FS, name string, id string) {
	lib, err := scripts.ReadFile("scripts/lib.py")
	if err != nil {
		fmt.Println(err)
		return
	}

	s, err := scripts.ReadFile("scripts/" + name + ".py")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(s))

	script := string(lib) + "\n\n" + string(s)

	cmd := exec.Command("python3", "-c", script, "--", id)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "AWS_ACCESS_KEY_ID="+os.Getenv("S3_ROOT_ACCESS_KEY"))
	cmd.Env = append(cmd.Env, "AWS_SECRET_ACCESS_KEY="+os.Getenv("S3_ROOT_SECRET_KEY"))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	/*out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
	*/
}

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

func postgresListener(scripts *embed.FS) {
	//testchannel := make(chan *pgconn.Notification)

	listener := &pgxlisten.Listener{
		Connect: func(ctx context.Context) (*pgx.Conn, error) {
			return pgx.Connect(ctx, GetDatabaseUri())
		},
	}

	listener.Handle("track_upload", pgxlisten.HandlerFunc(func(ctx context.Context, id *pgconn.Notification, conn *pgx.Conn) error {
		fmt.Println("new track uploaded", id)
		runScript(scripts, "track_upload", id.Payload)
		//testchannel <- notification
		return nil
	}))

	listener.Handle("cancel", pgxlisten.HandlerFunc(func(ctx context.Context, notification *pgconn.Notification, conn *pgx.Conn) error {
		fmt.Println("bye")
		time.Sleep(5 * time.Second)
		return nil
	}))

	go func() {
		err := listener.Listen(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func DefaultDaemon(scripts *embed.FS) chan bool {
	done := make(chan bool)
	postgresListener(scripts)
	go daemon(done)
	return done
}
