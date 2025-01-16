package datasrv

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotFound = erorr.Error("not found")
)

func Get(name string) ([]any, error) {
//@TODO
	return attachments, nil
}

//@TODO
var attachments = []any{
	struct{
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type:"Label",
		Name:"Programmer",
	},
	struct{
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type:"Label",
		Name:"Canadian",
	},
	struct{
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		URL       string `json:"url"`
	}{
		Type:"Icon",
		MediaType:"image/png",
		URL:"https://example.com/pfp/751f76ed540a40a3b4caae50e50cc867.png",
	},
	struct{
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		URL       string `json:"url"`
	}{
		Type:"Icon",
		MediaType:"image/png",
		URL:"https://example.com/pfp/45f7b459257940f490133070a975924b.png",
	},
}
