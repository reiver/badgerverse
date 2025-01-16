package verboten

import (
	"encoding/base64"
	"net/http"
	gourl "net/url"

	libpath "github.com/reiver/go-path"

	"github.com/reiver/badgerverse/srv/http"
	"github.com/reiver/badgerverse/srv/log"
)

const path string = "/hapi/v1/profiles"

const queryName string = "addr"

func init() {
	err := httpsrv.Mux.HandlePath(http.HandlerFunc(serveHTTP), path)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	log := logsrv.Prefix("www("+path+")").Begin()
	defer log.End()

	if nil == responsewriter {
		log.Error("nil response-writer")
		return
	}
	if nil == request {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil request")
		return
	}
	if nil == request.URL {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil request-url")
		return
	}

	var addr string
	{
		var query gourl.Values = request.URL.Query()
		if len(query) <= 0 {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("empty url-query")
			return
		}

		addr = query.Get(queryName)
		if "" == addr {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("empty \"address\" query parameter")
			return
		}
	}
	log.Informf("addr = %q", addr)

	var redirectURL string
	{
		var codename string = base64.URLEncoding.EncodeToString([]byte(addr))
		log.Informf("md5(addr) = %q", codename)

		redirectURL = libpath.Join(path, codename)
	}
	log.Informf("redirect-url = %q", redirectURL)

	http.Redirect(responsewriter, request, redirectURL, http.StatusTemporaryRedirect)
}
