package download

import (
	"time"
)

const (
	STATUS_NAME_SUCCESS    = "Success"
	STATUS_NAME_FAILED     = "Failed"
	STATUS_NAME_PAUSED     = "Paused"
	STATUS_NAME_QUEUED     = "Queued"
	STATUS_NAME_PROCESSING = "Processing"

	STATUS_ICON_SUCCESS    = "mdi-check-circle-outline"
	STATUS_ICON_FAILED     = "mdi-alert-outline"
	STATUS_ICON_PAUSED     = "mdi-pause-circle-outline"
	STATUS_ICON_QUEUED     = "mdi-tray-full"
	STATUS_ICON_PROCESSING = "mdi-progress-helper"

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
	Store map[string]Download

	Download struct {
		ID          string         `json:"id"`
		Name        string         `json:"name"`
		TimeElapsed string         `json:"timeElapsed"`
		Size        int64          `json:"size"`
		Date        time.Time      `json:"date"`
		Chunks      []Chunk        `json:"chunks"`
		Status      DownloadStatus `json:"status"`
		Progres     float64        `json:"progress"`
		Type        DownloadType   `json:"type"`
		Metadata    Metadata       `json:"metadata"`
	}

	Chunk struct {
		Downloaded  int64   `json:"downloaded"`
		Progressbar float64 `json:"progressbar"`
	}
	DownloadStatus struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	}

	Metadata struct {
		Url       string `json:"url"`
		Cansplit  bool   `json:"cansplit"`
		Totalpart int    `json:"totalpart"`
	}

	DownloadType struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	}
)
