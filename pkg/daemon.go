package htwarchive

import (
	"fmt"
	"time"
)

func daemon(b chan bool) {
	i := 0
	for i < 100 {
		fmt.Println("Daemon Test")
		time.Sleep(5 * time.Second)
		i++
	}

	b <- true
}

func NewDefaultDaemon() chan bool {
	done := make(chan bool)
	go daemon(done)
	return done
}
