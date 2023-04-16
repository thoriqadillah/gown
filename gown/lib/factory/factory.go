package factory

type Factory interface {
	Create() interface{}
}
