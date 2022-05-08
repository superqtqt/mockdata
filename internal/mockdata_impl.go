package internal

type Imp struct {
	Random func() float64
}

func (i *Imp) Mock(v interface{}) interface{} {
	return nil
}
