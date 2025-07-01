package roadCongestion

type WeatherCondition int

const (
    Clear WeatherCondition = iota
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64{
	count := 1.0
	if weather.Condition == 1{count += 0.125 }
	if weather.Condition == 2{count += 0.2 }
	if weather.Condition == 3{count += 0.15 }
	if weather.WindSpeed > 15{count += 0.1}
	return count
}

type TrafficClient interface {
    GetTrafficLevel(lat, lng float64) int 
}

func GetTrafficMultiplier(trafficLevel int) float64 {
    return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
    TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
    return 3
}
 