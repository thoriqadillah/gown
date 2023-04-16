package filetype

import (
	"changeme/gown/lib/factory"
	"sync"
)

type filetypeFactory struct {
	factory.Factory
}

var start sync.Once

func NewFactory(filename string) filetypeFactory {
	var factories map[string]factory.Factory

	start.Do(func() {
		//TODO: implement factory
	})

	return filetypeFactory{
		Factory: factories[filename],
	}
}
