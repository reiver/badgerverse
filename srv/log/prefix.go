package logsrv

import (
	"github.com/reiver/go-log"
)

func Prefix(name ...string) log.Logger {
	if nil == prefixer {
		return nil
	}

	return prefixer.Prefix(name...)
}
