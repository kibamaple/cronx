package common

type CJob struct {
	Name string
	Spec string
	Task string
	Params map[string]string
	Status uint8
}

