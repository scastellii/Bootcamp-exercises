package products

type BigProduct struct {
	Cost            float64
	MaintenanceCost float64
	AdditionalCost  float64
}

func (b BigProduct) Price() float64 {
	return b.Cost + b.Cost*b.MaintenanceCost + b.AdditionalCost
}
