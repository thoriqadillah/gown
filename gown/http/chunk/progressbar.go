package chunk

import (
	"context"
	"fmt"
	"io"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type progressbar struct {
	Reader     io.Reader
	ctx        context.Context
	id         string
	index      int
	partsize   int64
	totalsize  int64
	transfered int64
	tmp        int
	err        error
}

var errCanceled = fmt.Errorf("download canceled")

func (r *progressbar) Read(payload []byte) (n int, err error) {
	n, err = r.Reader.Read(payload)
	if err != nil {
		return n, err
	}

	runtime.EventsOn(r.ctx, "stop", func(optionalData ...interface{}) {
		r.err = errCanceled
	})

	if r.err != nil {
		return n, r.err
	}

	r.transfered += int64(n)
	r.tmp += n

	// emit event every 300kb downloaded because the default 32kb is too fast and the frontend cannot handle it
	if r.tmp > 300*1024 {
		runtime.EventsEmit(r.ctx, "transfered",
			r.id,
			r.index,
			float64(r.transfered)/float64(r.partsize)*100,
			float64(100*r.tmp)/float64(r.totalsize),
		)

		r.tmp = 0
	}

	return n, err
}
