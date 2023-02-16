package faas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/synyi/faas.go/proto"
)

type EventCtx struct {
	header http.Header
	send   func()
	retry  func(string)
	*proto.Event
	resp        proto.Response
	respHeaders http.Header
	err         error
}

func ctxFromEventData(d []byte) *EventCtx {
	var e proto.Event
	err := e.Unmarshal(d)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &EventCtx{
		Event:       &e,
		resp:        proto.Response{EventId: e.EventId, Status: 200, RequestId: e.RequestId},
		respHeaders: http.Header{},
	}
}

func ctxFromEvent(e *proto.Event) *EventCtx {
	return &EventCtx{
		Event:       e,
		resp:        proto.Response{EventId: e.EventId, Status: 200, RequestId: e.RequestId},
		respHeaders: http.Header{},
	}
}

//// Send set body(see EventCtx.SetBody) and send response immediately
//func (c *EventCtx) Send(body interface{}) error {
//	err := c.Respinse(body)
//	if err != nil {
//		return err
//	}
//	c.send()
//	return nil
//}

// Header get event headers as http.Header
func (c *EventCtx) Header() http.Header {
	if c.header != nil {
		return c.header
	}
	c.header = map[string][]string{}
	for i := 0; i < len(c.Headers)/2; i++ {
		c.header.Add(c.Headers[2*i], c.Headers[2*i+1])
	}
	return c.header
}

// Response set response body. supported type: string, io.Reader, struct, error and send to client immediately
func (c *EventCtx) Response(status int32, value interface{}) error {
	if status == 0 {
		c.resp.Status = 200
	} else {
		c.resp.Status = status
	}
	switch val := value.(type) {
	case string:
		c.resp.Body = []byte(val)
		if c.respHeaders.Get("content-type") == "" {
			c.respHeaders.Set("content-type", "text/plain;charset=utf8")
		}
		c.respHeaders.Set("content-length", strconv.Itoa(len(c.resp.Body)))
	case io.Reader:
		d, _ := io.ReadAll(val)
		c.respHeaders.Set("content-length", strconv.Itoa(len(d)))
		c.resp.Body = d
	case error:
		c.resp.Body = []byte(val.Error())
		if c.respHeaders.Get("content-type") == "" {
			c.respHeaders.Set("content-type", "text/plain;charset=utf8")
		}
		c.respHeaders.Set("content-length", strconv.Itoa(len(c.resp.Body)))
	default:
		d, err := json.Marshal(val)
		if err != nil {
			return err
		}
		c.respHeaders.Set("content-length", strconv.Itoa(len(d)))
		c.respHeaders.Set("content-type", "application/json")
		c.resp.Body = d
	}
	c.send()
	return nil
}

//func (c *EventCtx) Error(status int32, error interface{}) {
//	if status < 200 || status > 599 {
//		status = 500
//	}
//	c.resp.Status = status
//	_ = c.Send(error)
//}

func (c *EventCtx) GetBlob(id string) (io.ReadCloser, error) {
	u := gwUrl
	u.Path = "/api/blob/" + id
	log.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, errors.New(string(data))
	}
	return resp.Body, nil
}

func (c *EventCtx) UploadBlob(r io.Reader) (string, error) {
	// u := gwUrl
	// u.Path = "/api/blob"
	// u.Query().Add("persist", "true")
	// resp, err := http.Post(u.String(), "application/zip", r)
	g := os.Getenv("FAAS_GATEWAY")
	resp, err := http.Post(fmt.Sprintf("%s/api/blob?persist=true", g), "application/zip", r)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return "", errors.New(string(data))
	}
	d, _ := io.ReadAll(resp.Body)
	id := string(d)
	return id, nil
}
func (c *EventCtx) NeedRetry(reason string) {
	c.retry(reason)
}
