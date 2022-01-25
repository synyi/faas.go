package faas

import "net/http"

type Event interface {
	Header() http.Header
	Response(status int32, v interface{}) error
	NeedRetry()
}
