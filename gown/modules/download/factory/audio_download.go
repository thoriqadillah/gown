package factory

import (
	"changeme/gown/lib/factory"
	"changeme/gown/modules/download"
	"time"
)

type AudioFactory struct {
	res *download.Response
}

func audioFactory(res *download.Response) factory.Factory[download.Download] {
	return &AudioFactory{
		res: res,
	}
}

func (v *AudioFactory) Create() download.Download {
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
			Name:  download.TYPE_NAME_AUDIO,
			Icon:  download.TYPE_ICON_AUDIO,
			Color: download.TYPE_COLOR_AUDIO,
		},
		Metadata: download.Metadata{
			Url:       v.res.Url,
			Cansplit:  v.res.Cansplit,
			Totalpart: v.res.Totalpart,
		},
	}
}

func init() {
	register("audio", audioFactory)
}
