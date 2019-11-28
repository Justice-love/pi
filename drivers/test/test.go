package test

type Adaptor struct {
	N         string
	WriteChan chan<- CheckValue
}
type CheckValue struct {
	Pin string
	Val byte
}

func (a *Adaptor) DigitalWrite(pin string, val byte) (err error) {
	a.WriteChan <- CheckValue{pin, val}
	return nil
}

func (a *Adaptor) Name() string {
	return a.N
}

func (a *Adaptor) SetName(n string) {
	a.N = n
}

func (a *Adaptor) Connect() error {
	return nil
}

func (a *Adaptor) Finalize() error {
	return nil
}
