package common

const Decimal = 10

const (
	Failed uint8 = iota
)

type IHandle interface {
	GetConcurrent() int
	Process(<-chan struct{})
}

type CConcurrent struct {
	Concurrent int	
}

func (this CConcurrent) GetConcurrent() int{
	return this.Concurrent 
}

type IService interface {
	Init()
	Start()
	Stop()
}