package square

import (
	"math"
)

type sides int32

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {

	s := int(sidesNum)
	if s == 3 {
		return math.Sqrt(3) / 4 * sideLen
	}

	if s == 4 {
		return (sideLen * sideLen)
	}

	if s == 0 {
		return math.Pi * (sideLen * sideLen) / 2
	}

	return 0
}
