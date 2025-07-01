package pricelimit

const(
	minPrice float64 = 99.0
	maxPrice float64 = 20000  
)

func ApplyPriceLimits(price float64) float64{
	if price > maxPrice{return maxPrice}
	if price < minPrice{return minPrice}
	return price
}