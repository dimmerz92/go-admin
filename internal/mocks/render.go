package mocks

import (
	"context"
	"io"
)

type Renderable struct {
	value string
}

func Template(value string) *Renderable {
	return &Renderable{value: value}
}

func (r *Renderable) Render(ctx context.Context, buf io.Writer) error {
	_ = ctx
	_, err := buf.Write([]byte(r.value))
	return err
}
