package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type AudioFactory struct {
	res *http.Response
}

func audioFactory(res *http.Response) factory.Factory[Download] {
	return &AudioFactory{
		res: res,
	}
}

func (v *AudioFactory) Create() Download {
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
			Name:  TYPE_NAME_AUDIO,
			Icon:  TYPE_ICON_AUDIO,
			Color: TYPE_COLOR_AUDIO,
		},
	}
}

func init() {
	register("audio", audioFactory)
}
