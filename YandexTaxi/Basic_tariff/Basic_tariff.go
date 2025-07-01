package basicTariff

var (
	pricePerKm float64 = 10.0
	pricePerMinute float64 = 2.0
)


type TripParameters struct{
	Distance float64
	Duration float64
}

func CalculateBasePrice(Parametres TripParameters) float64 {
	price := Parametres.Distance * pricePerKm + Parametres.Duration * pricePerMinute
	return price
}