package logsrv

import (
	"io"
	"os"

	"github.com/reiver/go-log"
)

var writer io.Writer = os.Stdout

var prefixer log.Prefixer = log.NewLogger(writer)
