package faas

import (
	"encoding/json"
	"errors"
	"github.com/synyi/faas.go/proto"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

type EventCtx struct {
	header http.Header
	send   func()
	*proto.Event
	resp        proto.Response
	respHeaders http.Header
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
		resp:        proto.Response{EventId: e.EventId, Status: 200},
		respHeaders: http.Header{},
	}
}

func ctxFromEvent(e *proto.Event) *EventCtx {
	return &EventCtx{
		Event:       e,
		resp:        proto.Response{EventId: e.EventId, Status: 200},
		respHeaders: http.Header{},
	}
}

// Send set body(see EventCtx.SetBody) and send response immediately
func (c *EventCtx) Send(body interface{}) error {
	err := c.SetBody(body)
	if err != nil {
		return err
	}
	c.send()
	return nil
}

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

// SetBody set response body. supported type: string, io.Reader, struct
func (c *EventCtx) SetBody(value interface{}) error {
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
	return nil
}

func (c *EventCtx) Error(status int32, error interface{}) {
	if status < 200 || status > 599 {
		status = 500
	}
	c.resp.Status = status
	_ = c.Send(error)
}

func (c *EventCtx) GetBlob(id string) (io.ReadCloser, error) {
	if gwUrl == "" {
		return nil, errors.New("gateway url not defined (use FAAS_GATEWAY env)")
	}
	resp, err := http.Get(path.Join(gwUrl, "/api/blob/", id))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *EventCtx) UploadBlob(r io.Reader) (string, error) {
	if gwUrl == "" {
		return "", errors.New("gateway url not defined (use FAAS_GATEWAY env)")
	}
	resp, err := http.Post(path.Join(gwUrl, "/api/blob/persist"), "application/zip", r)
	if err != nil {
		return "", err
	}
	d, _ := io.ReadAll(resp.Body)
	id := string(d)
	return id, nil
}
