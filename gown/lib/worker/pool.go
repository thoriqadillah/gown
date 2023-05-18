package worker

type Pool interface {
	Start()
	Stop()
	Add(Job)
}

type Job interface {
	Execute() error
	HandleError(error)
}
