package faas

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	"github.com/synyi/faas.go/proto"
	"log"
	"os"
	"sync"
)

var functionName string
var waiterMap map[string]chan *proto.Response
var wl = sync.Mutex{}

func initCallback() {
	functionName = "function:" + os.Getenv("FAAS_NAME")
	if functionName == "" {
		functionName = "other caller"
	}
	c := make(chan *nats.Msg)
	_, err := nc.ChanSubscribe("faas.response."+clientId, c)
	if err != nil {
		panic(err)
	}
	go func() {
		for m := range c {
			var r *proto.Response
			err = r.Unmarshal(m.Data)
			if err != nil {
				log.Println(err)
				continue
			}
			wl.Lock()
			rc, ok := waiterMap[r.EventId]
			delete(waiterMap, r.EventId)
			wl.Unlock()
			if ok {
				rc <- r
			}
		}
	}()
}

func wait(id string) chan *proto.Response {
	c := make(chan *proto.Response)
	wl.Lock()
	waiterMap[id] = c
	wl.Unlock()
	return c
}

var callback = sync.Once{}

// Call another faas function
func call(ctx context.Context, target string, req *proto.Event, ttl int32) (*proto.Response, error) {
	callback.Do(initCallback)
	req.EventId = nuid.Next()
	req.Ttl = ttl
	req.Source = functionName
	req.SenderId = clientId
	d, err := req.Marshal()
	if err != nil {
		return nil, err
	}
	c := wait(req.EventId)
	err = nc.Publish("faas.event."+target, d)
	if err != nil {
		return nil, err
	}
	select {
	case r := <-c:
		return r, nil
	case <-ctx.Done():
		wl.Lock()
		delete(waiterMap, req.EventId)
		close(c)
		wl.Unlock()
		return nil, ctx.Err()
	}
}
