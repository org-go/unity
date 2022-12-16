package help

import (
	"fmt"
	"math"
	"strconv"
)

func FloatAddition(x, y float64) float64 {
	return action2(x + y)
}

func FloatSubtraction(x, y float64) float64 {
	return action2(x - y)
}

func FloatMultiplication(x, y float64) float64 {
	return action2(x * y)
}

func FloatDivision(x, y float64) float64 {
	return action2(x / y)
}

func action2(val float64) float64 {

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return value

}


func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000 // 6378137
	rad := math.Pi/180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1) * math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2) * math.Cos(theta))

	return dist * float64(radius)

}
