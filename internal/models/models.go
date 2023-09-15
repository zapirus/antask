package models

import (
	"encoding/json"
	"net/http"
)

type ReqData struct {
	Headers http.Header
	Body    json.RawMessage
}
