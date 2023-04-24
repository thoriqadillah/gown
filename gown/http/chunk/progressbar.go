package chunk

import (
	"context"
	"io"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type progressbar struct {
	Reader     io.Reader
	ctx        context.Context
	id         string
	index      int
	size       int64
	transfered int64
}

func (r *progressbar) Read(payload []byte) (n int, err error) {
	n, err = r.Reader.Read(payload)
	if err != nil {
		return n, err
	}

	r.transfered += int64(n)

	runtime.EventsEmit(r.ctx, "transfered",
		r.id,
		r.index,
		float64(r.transfered)/float64(r.size)*100, // for the actual progress bar
		n, // for percentage
	)

	return n, err
}
