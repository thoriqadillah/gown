package store

import "changeme/gown/modules/setting"

type Store interface {
	GetSetting() *setting.Settings
	UpdateSetting(data *setting.Settings) error
}
