package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type CompressedFactory struct {
	res *http.Response
}

func compressedFactory(res *http.Response) factory.Factory[Download] {
	return &CompressedFactory{
		res: res,
	}
}

func (v *CompressedFactory) Create() Download {
	return Download{
		ID:          uuid.New().String(),
		Name:        v.res.Filename,
		TimeElapsed: "",
		Size:        v.res.Size,
		Progres:     0,
		Chunks:      make([]Chunk, v.res.Totalpart),
		Date:        time.Now(),
		Status: DownloadStatus{
			Name:  STATUS_NAME_PROCESSING,
			Icon:  STATUS_ICON_PROCESSING,
			Color: STATUS_COLOR_PROCESSING,
		},
		Type: DownloadType{
			Name:  TYPE_NAME_COMPRESSED,
			Icon:  TYPE_ICON_COMPRESSED,
			Color: TYPE_COLOR_COMPRESSED,
		},
		Metadata: Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("compressed", compressedFactory)
}
