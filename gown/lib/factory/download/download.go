package download

import (
	"changeme/gown/lib/factory"
	"changeme/gown/modules/download"
	"sync"
)

type FactoryImpl func(res *download.Response) factory.Factory[download.Download]

var start sync.Once
var factories map[string]FactoryImpl

func NewFactory(res *download.Response) factory.Factory[download.Download] {
	factory := factories[res.ContentType]

	return factory(res)
}

func register(name string, factory FactoryImpl) {
	start.Do(func() {
		factories = make(map[string]FactoryImpl)
	})

	factories[name] = factory
}
