package goplayground

import (
	"testing"
	"goplayground/utils"
	"math"
)

func TestTimeConsuming(t *testing.T) {
	num := 187.0
	sqrt := math.Sqrt(num)
	if approximatedSqrt := utils.NeotainSqrt(num); math.Abs(approximatedSqrt - sqrt) > 0.000001 {
		t.Errorf("%v is not close to %v", approximatedSqrt, sqrt)
	}
}