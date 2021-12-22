package faas

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	"github.com/synyi/faas.go/proto"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var nc *nats.Conn

func prodInit(handler func(ctx context.Context, eventCtx *EventCtx) error) {
	target := os.Getenv("FAAS_TARGET")
	natsUrl := os.Getenv("NATS_URL")
	timeout := os.Getenv("FAAS_TIMEOUT")
	concurrent := os.Getenv("FAAS_CONCURRENT")
	if target == "" {
		log.Fatalln("env FAAS_TARGET missing")
	}
	if natsUrl == "" {
		log.Fatalln("env NATS_URL missing")
	}
	timeoutDuration := time.Second * 120
	if timeout != "" {
		t, _ := strconv.Atoi(timeout)
		if t > 0 {
			timeoutDuration = time.Duration(t) * time.Second
		}
	}
	var concurrenti int
	if concurrent != "" {
		concurrenti, _ = strconv.Atoi(concurrent)
		if concurrenti < 1 {
			concurrenti = 1
		}
	}
	var err error
	nc, err = nats.Connect(natsUrl)
	if err != nil {
		log.Panicln("cannot connect to nats, ", err)
	}
	js, _ := nc.JetStream()
	stream := "faas.event." + target
	msgCh := make(chan *nats.Msg, concurrenti)
	sub, err := js.ChanQueueSubscribe(stream, "queue."+stream, msgCh, nats.Durable(strings.ReplaceAll(stream, ".", "_")))
	if err != nil {
		log.Panicln(err)
	}
	stopCtx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM)
		<-c
		cancel()
		_ = sub.Drain()
		os.Exit(0)
	}()
	wg := sync.WaitGroup{}
	wg.Add(concurrenti)
	for i := 0; i < concurrenti; i++ {
		go func() {
			for {
				var msg *nats.Msg
				select {
				case <-stopCtx.Done():
					wg.Done()
					return
				case msg = <-msgCh:
				}
				err = func() (err error) {
					var senderr error
					var sent bool
					ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
					defer func() {
						cancel()
						e := recover()
						if e != nil {
							err = fmt.Errorf("%s", e)
						}
					}()
					c := ctxFromEventData(msg.Data)
					c.send = func() {
						c.resp.Headers = buildHeaders(c.respHeaders)
						c.resp.Time = time.Now().UnixMilli()
						d, _ := c.resp.Marshal()
						senderr = nc.Publish("faas.response."+c.SenderId, d)
						sent = true
					}
					go func() {
						<-ctx.Done()
						if ctx.Err() == context.DeadlineExceeded {
							log.Println("handler not finish in timeout, exit")
							os.Exit(1)
						}
					}()
					err = handler(context.Background(), c)
					if !sent {
						if err == nil {
							c.send()
							err = senderr
						} else {
							c.Error(500, err)
						}
						if err == nil {
							_ = msg.Ack()
						} else {
							_ = msg.Nak()
						}
					}
					return
				}()
				if err != nil {
					log.Println("panic:", err)
					_ = msg.Nak()
				}
			}
		}()
	}
	wg.Wait()
}

func testInit(handler func(ctx context.Context, eventCtx *EventCtx) error) {

}

func buildHeaders(h http.Header) []string {
	var ret []string
	for k, v := range h {
		for _, vv := range v {
			ret = append(ret, k, vv)
		}
	}
	return ret
}

func localInit(handler func(ctx context.Context, eventCtx *EventCtx) error) {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":5001"
	}
	h := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		event := &proto.Event{
			EventId:     nuid.Next(),
			SenderId:    "LOCAL",
			Type:        proto.HTTP,
			Method:      r.Method,
			Url:         r.RequestURI,
			Body:        body,
			BodyType:    proto.Raw,
			ContentType: "",
			Headers:     buildHeaders(r.Header),
			Source:      r.RemoteAddr,
		}
		c := ctxFromEvent(event)
		sended := false
		c.send = func() {
			h := wr.Header()
			for k, v := range c.respHeaders {
				for _, vv := range v {
					h.Add(k, vv)
				}
			}
			wr.WriteHeader(int(c.resp.Status))
			wr.Write(c.resp.Body)
			sended = true
		}
		err := handler(context.Background(), c)
		if sended {
			if err != nil {
				log.Printf("err after Send(), %v", err)
			}
			return
		}
		if err != nil {
			http.Error(wr, err.Error(), 500)
		} else {
			c.send()
		}
	})
	log.Println("test handler in " + listen)
	log.Fatalln(http.ListenAndServe(listen, h))
}

func Init(handler func(ctx context.Context, eventCtx *EventCtx) error) {
	mode := os.Getenv("MODE")
	switch mode {
	case "prod":
		prodInit(handler)
	case "test":
		testInit(handler)
	default:
		localInit(handler)
	}
}
