package products

type LittleProduct struct {
	Cost float64
}

func (l LittleProduct) Price() float64 {
	return l.Cost
}
