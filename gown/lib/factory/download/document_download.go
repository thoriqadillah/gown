package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type DocumentFactory struct {
	res *http.Response
}

func documentFactory(res *http.Response) factory.Factory {
	return &DocumentFactory{
		res: res,
	}
}

func (v *DocumentFactory) Create() interface{} {
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
			Name:  TYPE_NAME_DOCUMENT,
			Icon:  TYPE_ICON_DOCUMENT,
			Color: TYPE_COLOR_DOCUMENT,
		},
	}
}