package verboten

import (
	"io"
	"net/http"

	"github.com/reiver/badgerverse/srv/http"
)

const path string = "/"

func init() {
	err := httpsrv.Mux.HandlePath(http.HandlerFunc(serveHTTP), path)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	if nil == responsewriter {
		return
	}
	if nil == request {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		return
	}

	io.WriteString(responsewriter, "<h1>badgerverse</h1><p>mushroom, mushroom.</p>")
}
