package products

type MediumProduct struct {
	Cost            float64
	MaintenanceCost float64
}

func (m MediumProduct) Price() float64 {
	return m.Cost + m.Cost*m.MaintenanceCost
}
