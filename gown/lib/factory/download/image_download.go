package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"time"
)

type ImageFactory struct {
	res *http.Response
}

func imageFactory(res *http.Response) factory.Factory[Download] {
	return &ImageFactory{
		res: res,
	}
}

func (v *ImageFactory) Create() Download {
	return Download{
		ID:          factory.ID(5),
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
			Name:  TYPE_NAME_IMAGE,
			Icon:  TYPE_ICON_IMAGE,
			Color: TYPE_COLOR_IMAGE,
		},
		Metadata: Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("image", imageFactory)
}
