package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"

	"github.com/google/uuid"
)

type OtherFactory struct {
	res *http.Response
}

func otherFactory(res *http.Response) factory.Factory[Download] {
	return &OtherFactory{
		res: res,
	}
}

func (v *OtherFactory) Create() Download {
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
			Name:  TYPE_NAME_OTHER,
			Icon:  TYPE_ICON_OTHER,
			Color: TYPE_COLOR_OTHER,
		},
		Metadata: Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("other", otherFactory)
}
