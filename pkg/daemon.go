package htwarchive

func NewDefaultDaemon() chan bool {
	done := make(chan bool)
	return done
}
