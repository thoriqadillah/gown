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

func videoFactory(res *http.Response) factory.Factory {
	return &VideoFactory{
		res: res,
	}
}

func (v *VideoFactory) Create() interface{} {
	return Download{
		ID:          uuid.New().String(),
		Name:        v.res.Filename,
		TimeElapsed: 0,
		Size:        v.res.Size,
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
	}
}