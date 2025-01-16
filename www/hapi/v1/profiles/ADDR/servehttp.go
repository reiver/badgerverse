package verboten

import (
	"encoding/base64"
	"net/http"

	"github.com/reiver/go-jsonld"

	"github.com/reiver/badgerverse/srv/data"
	"github.com/reiver/badgerverse/srv/http"
	"github.com/reiver/badgerverse/srv/log"
)

const queryParam string = "ADDR"

const pathpattern string = "/hapi/v1/profiles/{"+queryParam+"}"

func init() {
	err := httpsrv.Mux.HandlePattern(httpsrv.PatternHandlerFunc(serveHTTP), pathpattern)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *httpsrv.ParameterizedRequest) {
	log := logsrv.Prefix("www("+pathpattern+")").Begin()
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

	var param0 string
	{
		var found bool
		param0, found = request.ParameterByIndex(0)
		if !found {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("parameter %q not found", queryParam)
			return
		}
	}
	log.Debugf("param0 = %q", param0)

	var addr string
	{
		bytes, err := base64.URLEncoding.DecodeString(param0)
		if nil != err {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("problem trying to decode param0 %q as base64url: %s", param0, err)
			return
		}

		addr = string(bytes)
	}
	log.Informf("addr = %q", addr)

	var attachments []any
	{
		var err error

		attachments, err = datasrv.Get(addr)
		if nil != err {
			switch err {
			case datasrv.ErrNotFound:
				const code int = http.StatusNotFound
				http.Error(responsewriter, http.StatusText(code), code)
				log.Debugf("data for \"addr\" %d not found", addr)
				return
			default:
				const code int = http.StatusInternalServerError
				http.Error(responsewriter, http.StatusText(code), code)
				log.Errorf("problem getting data for \"addr\" %d: %s", addr, err)
				return
			}
		}
	}
	log.Debugf("len(attachments) = %d", len(attachments))

	var bytes []byte
	{
		var response = struct{
			NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`

			Type         string `json:"type"`
			Describes    string `json:"describes"`
			AttributedTo string `json:"attributedTo"`
			Icon         any    `json:"icon"`
			Attachment []any `json:"attachment"`
		}{
			Type:"Profile",
			Describes:addr,
				
			AttributedTo:"https://example.com/apps/ratel", //@TODO
			Icon: struct{
				Type      string `json:"type"`
				MediaType string `json:"mediaType"`
				URL       string `json:"url"`
			}{
				Type:"Icon",
				MediaType:"image/png",
				URL:"https://example.com/apps/ratel/img/icon.png", //@TODO
			},
				
			Attachment: attachments,
		}

		var err error
		bytes, err = jsonld.Marshal(response)
		if nil != err {
			const code int = http.StatusInternalServerError
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("problem jsonld-marshaling response: %s", err)
			return
		}
	}

	{
		responsewriter.Header().Add("Content-Type", "application/activity+json")

		_, err := responsewriter.Write(bytes)
		if nil != err {
			log.Errorf("problem sending response to client: %s", err)
		}
	}
}
