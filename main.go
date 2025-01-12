package main

import (
	"github.com/reiver/badgerverse/srv/log"
)

func main() {
	log := logsrv.Prefix("main")

	log.Inform("badgerverse ⚡")
	yell()

	log.Inform("Here we go…")
	webserve()
}
