package chunk

import (
	"changeme/gown/modules/download"
	"context"
	"fmt"
	"io"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type progressbar struct {
	Reader     io.Reader
	ctx        context.Context
	index      int
	partsize   int64
	toDownload *download.Download
	tmp        int
	err        error
}

var errCanceled = fmt.Errorf("download canceled")

func (r *progressbar) Read(payload []byte) (n int, err error) {
	n, err = r.Reader.Read(payload)
	if err != nil {
		return n, err
	}

	r.toDownload.Chunks[r.index].Downloaded += int64(n)
	r.toDownload.Chunks[r.index].Progressbar = float64(r.toDownload.Chunks[r.index].Downloaded) / float64(r.partsize) * 100
	r.tmp += n

	// emit event every 300kb downloaded because the default 32kb is too fast and the frontend cannot handle it
	if r.tmp > 300*1024 {
		runtime.EventsEmit(r.ctx, "transfered",
			r.toDownload.ID,
			r.index,
			r.toDownload.Chunks[r.index].Downloaded,
			r.toDownload.Chunks[r.index].Progressbar,      // actual progress bar
			float64(100*r.tmp)/float64(r.toDownload.Size), // progress in percentage
		)

		r.tmp = 0
	}

	runtime.EventsOn(r.ctx, "stop", func(optionalData ...interface{}) {
		r.err = errCanceled
	})

	if r.err != nil {
		runtime.EventsEmit(r.ctx, "total-bytes",
			r.toDownload.ID,
			r.index,
			r.toDownload.Chunks[r.index].Downloaded,  //update the latest downloaded
			r.toDownload.Chunks[r.index].Progressbar, // update the progress bar
		)
		return n, r.err
	}

	return n, err
}
