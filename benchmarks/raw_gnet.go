package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
)

// Raw gnet+ants benchmark server — equivalent to the Chuks HTTP engine architecture.

var (
	httpSep  = []byte("\r\n\r\n")
	httpCL   = []byte("Content-Length: ")
	resp200  = []byte("HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Length: 13\r\n\r\nHello, World!")
	resp404  = []byte("HTTP/1.1 404 Not Found\r\nConnection: keep-alive\r\nContent-Length: 19\r\n\r\n404 page not found\n")
	respJSON = []byte("HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Type: application/json\r\nContent-Length: 27\r\n\r\n{\"message\":\"Hello, World!\"}")
)

type silentLogger struct{}

func (l *silentLogger) Debugf(format string, args ...any) {}
func (l *silentLogger) Infof(format string, args ...any)  {}
func (l *silentLogger) Warnf(format string, args ...any)  {}
func (l *silentLogger) Errorf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "[error] "+format+"\n", args...)
}
func (l *silentLogger) Fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "[fatal] "+format+"\n", args...)
	os.Exit(1)
}

type reqCtx struct {
	conn gnet.Conn
	path []byte
}

type benchEngine struct {
	gnet.BuiltinEventEngine
	eng  gnet.Engine
	pool *ants.PoolWithFunc
}

func (e *benchEngine) OnBoot(eng gnet.Engine) gnet.Action {
	e.eng = eng
	poolSize := runtime.NumCPU() * 8192
	var err error
	e.pool, err = ants.NewPoolWithFunc(poolSize, func(v interface{}) {
		ctx := v.(*reqCtx)
		path := ctx.path
		if len(path) == 1 && path[0] == '/' {
			ctx.conn.AsyncWrite(resp200, nil)
		} else if bytes.Equal(path, []byte("/json")) {
			ctx.conn.AsyncWrite(respJSON, nil)
		} else {
			ctx.conn.AsyncWrite(resp404, nil)
		}
	}, ants.WithPreAlloc(false), ants.WithNonblocking(false))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: failed to create pool: %v\n", err)
		os.Exit(1)
	}
	return gnet.None
}

func (e *benchEngine) OnShutdown(eng gnet.Engine) {
	if e.pool != nil {
		e.pool.Release()
	}
}

func (e *benchEngine) OnTraffic(c gnet.Conn) (action gnet.Action) {
	buf, _ := c.Peek(-1)
	n := len(buf)
	if n == 0 {
		return
	}

	processed := 0
	for processed < n {
		data := buf[processed:]
		hdrEnd := bytes.Index(data, httpSep)
		if hdrEnd < 0 {
			break
		}

		sp1 := bytes.IndexByte(data, ' ')
		if sp1 <= 0 || sp1 > 7 {
			c.Close()
			return
		}
		sp2 := bytes.IndexByte(data[sp1+1:hdrEnd], ' ')
		if sp2 <= 0 {
			c.Close()
			return
		}

		reqLen := hdrEnd + 4
		isGet := data[0] == 'G' || data[0] == 'H'
		if !isGet {
			if clIdx := bytes.Index(data[:hdrEnd], httpCL); clIdx >= 0 {
				clStart := clIdx + 16
				clEnd := bytes.IndexByte(data[clStart:hdrEnd], '\r')
				if clEnd > 0 {
					cl, _ := strconv.Atoi(string(data[clStart : clStart+clEnd]))
					reqLen += cl
				}
			}
		}
		if processed+reqLen > n {
			break
		}

		rawPath := data[sp1+1 : sp1+1+sp2]
		pathCopy := make([]byte, len(rawPath))
		copy(pathCopy, rawPath)

		e.pool.Invoke(&reqCtx{conn: c, path: pathCopy})
		processed += reqLen
	}
	c.Discard(processed)
	return
}

func main() {
	logging.SetDefaultLoggerAndFlusher(&silentLogger{}, nil)
	fmt.Println("Raw gnet+ants benchmark server on :9090")
	err := gnet.Run(&benchEngine{}, "tcp://0.0.0.0:9090",
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
