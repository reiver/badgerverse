package main

import (
	"net/http"

	"github.com/reiver/badgerverse/cfg"
	"github.com/reiver/badgerverse/srv/http"
	"github.com/reiver/badgerverse/srv/log"

        // This import enables all the HTTP handlers.
      _ "github.com/reiver/badgerverse/www"
)

func webserve() {
	log := logsrv.Prefix("webserve")

	var tcpaddr string = cfg.WebServerTCPAddress()
	log.Informf("serving HTTP on TCP address: %q", tcpaddr)

	err := http.ListenAndServe(tcpaddr, &httpsrv.Mux)
	if nil != err {
		log.Errorf("ERROR: problem with serving HTTP on TCP address %q: %s", tcpaddr, err)
		panic(err)
	}
}
