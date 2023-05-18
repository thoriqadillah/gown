package factory

import (
	"changeme/gown/lib/factory"
	"changeme/gown/modules/download"
	"time"
)

type OtherFactory struct {
	res *download.Response
}

func otherFactory(res *download.Response) factory.Factory[download.Download] {
	return &OtherFactory{
		res: res,
	}
}

func (v *OtherFactory) Create() download.Download {
	return download.Download{
		ID:          factory.ID(5),
		Name:        v.res.Filename,
		TimeElapsed: "",
		Size:        v.res.Size,
		Progres:     0,
		Chunks:      make([]download.Chunk, v.res.Totalpart),
		Date:        time.Now(),
		Status: download.DownloadStatus{
			Name:  download.STATUS_NAME_PROCESSING,
			Icon:  download.STATUS_ICON_PROCESSING,
			Color: download.STATUS_COLOR_PROCESSING,
		},
		Type: download.DownloadType{
			Name:  download.TYPE_NAME_OTHER,
			Icon:  download.TYPE_ICON_OTHER,
			Color: download.TYPE_COLOR_OTHER,
		},
		Metadata: download.Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("other", otherFactory)
}
