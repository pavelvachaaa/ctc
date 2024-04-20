package domain

type Register struct {
	ID            int
	Queue         chan *Car
	HandleTimeMin int
	HandleTimeMax int
}
