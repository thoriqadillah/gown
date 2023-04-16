package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type ImageFactory struct {
	res *http.Response
}

func imageFactory(res *http.Response) factory.Factory {
	return &ImageFactory{
		res: res,
	}
}

func (v *ImageFactory) Create() interface{} {
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
			Name:  TYPE_NAME_IMAGE,
			Icon:  TYPE_ICON_IMAGE,
			Color: TYPE_COLOR_IMAGE,
		},
	}
}
