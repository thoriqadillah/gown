package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type VideoFactory struct {
	res *http.Response
}

func videoFactory(res *http.Response) factory.Factory[Download] {
	return &VideoFactory{
		res: res,
	}
}

func (v *VideoFactory) Create() Download {
	return Download{
		ID:          uuid.New().String(),
		Name:        v.res.Filename,
		TimeElapsed: "",
		Size:        v.res.Size,
		Progres:     0,
		Progressbar: make([]float64, v.res.Totalpart),
		Date:        time.Now(),
		Status: DownloadStatus{
			Name:  STATUS_NAME_PROCESSING,
			Icon:  STATUS_ICON_PROCESSING,
			Color: STATUS_COLOR_PROCESSING,
		},
		Type: DownloadType{
			Name:  TYPE_NAME_VIDEO,
			Icon:  TYPE_ICON_VIDEO,
			Color: TYPE_COLOR_VIDEO,
		},
		Metadata: Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("video", videoFactory)
}
