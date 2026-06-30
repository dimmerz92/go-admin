package common

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
)

const maxBufSize = 4096

type Renderable interface {
	Render(ctx context.Context, buf io.Writer) error
}

var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

func acquireBuffer() *bytes.Buffer {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func returnBuffer(buf *bytes.Buffer) {
	if buf.Cap() > maxBufSize {
		return
	}
	bufPool.Put(buf)
}

func Render(w http.ResponseWriter, r *http.Request, status int, tpls ...Renderable) error {
	buf := acquireBuffer()
	defer returnBuffer(buf)

	for _, tpl := range tpls {
		if err := tpl.Render(r.Context(), buf); err != nil {
			return err
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	_, err := buf.WriteTo(w)
	return err
}
