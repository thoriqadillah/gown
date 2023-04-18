package download

import (
	"changeme/gown/http"
	"changeme/gown/lib/factory"
	"sync"
	"time"
)

const (
	STATUS_NAME_SUCCESS    = "success"
	STATUS_NAME_FAILED     = "failed"
	STATUS_NAME_PAUSED     = "paused"
	STATUS_NAME_QUEUED     = "queued"
	STATUS_NAME_PROCESSING = "processing"

	STATUS_ICON_SUCCESS    = "mdi-check-circle-outline"
	STATUS_ICON_FAILED     = "mdi-alert-outline"
	STATUS_ICON_PAUSED     = "mdi-pause-circle-outline"
	STATUS_ICON_QUEUED     = "mdi-tray-full"
	STATUS_ICON_PROCESSING = ""

	STATUS_COLOR_SUCCESS    = "success"
	STATUS_COLOR_FAILED     = "warning"
	STATUS_COLOR_PAUSED     = ""
	STATUS_COLOR_QUEUED     = "info"
	STATUS_COLOR_PROCESSING = ""

	TYPE_NAME_VIDEO      = "video"
	TYPE_NAME_IMAGE      = "image"
	TYPE_NAME_AUDIO      = "audio"
	TYPE_NAME_COMPRESSED = "compressed"
	TYPE_NAME_DOCUMENT   = "document"
	TYPE_NAME_OTHER      = "other"

	TYPE_ICON_VIDEO      = "mdi-video"
	TYPE_ICON_AUDIO      = "mdi-music-box"
	TYPE_ICON_DOCUMENT   = "mdi-file-document"
	TYPE_ICON_IMAGE      = "mdi-image"
	TYPE_ICON_COMPRESSED = "mdi-zip-box"
	TYPE_ICON_OTHER      = "mdi-file-question"

	TYPE_COLOR_DOCUMENT   = "blue-accent-2"
	TYPE_COLOR_VIDEO      = "deep-orange-accent-2"
	TYPE_COLOR_AUDIO      = "purple-accent-2"
	TYPE_COLOR_COMPRESSED = "yellow-accent-4"
	TYPE_COLOR_IMAGE      = "red-accent-2"
	TYPE_COLOR_OTHER      = ""
)

type (
	Download struct {
		ID          string         `json:"id"`
		Name        string         `json:"name"`
		TimeElapsed int            `json:"timeElapsed"`
		Size        int64          `json:"size"`
		Date        time.Time      `json:"date"`
		Status      DownloadStatus `json:"status"`
		Type        DownloadType   `json:"type"`
	}

	DownloadStatus struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	}

	DownloadType struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	}

	DownloadData struct {
		Response *http.Response `json:"response"`
		Data     Download       `json:"data"`
	}
)

type FactoryImpl func(res *http.Response) factory.Factory[Download]

var start sync.Once
var factories map[string]FactoryImpl

func NewFactory(res *http.Response) factory.Factory[Download] {
	factory := factories[res.ContentType]

	return factory(res)
}

func register(name string, factory FactoryImpl) {
	start.Do(func() {
		factories = make(map[string]FactoryImpl)
	})

	factories[name] = factory
}
