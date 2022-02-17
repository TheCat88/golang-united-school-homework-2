package square

import (
	"math"
)

type sides string

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == "SidesTriangle" {
		return 0.5 * (sideLen * sideLen)
	}

	if sidesNum == "SidesSquare" {
		return (sideLen * sideLen)
	}

	if sidesNum == "SidesCircle" {
		return math.Pi * (sideLen * sideLen)
	}

	return 0
}
