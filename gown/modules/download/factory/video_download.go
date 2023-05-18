package factory

import (
	"changeme/gown/lib/factory"
	"changeme/gown/modules/download"
	"time"
)

type VideoFactory struct {
	res *download.Response
}

func videoFactory(res *download.Response) factory.Factory[download.Download] {
	return &VideoFactory{
		res: res,
	}
}

func (v *VideoFactory) Create() download.Download {
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
			Name:  download.TYPE_NAME_VIDEO,
			Icon:  download.TYPE_ICON_VIDEO,
			Color: download.TYPE_COLOR_VIDEO,
		},
		Metadata: download.Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("video", videoFactory)
}
