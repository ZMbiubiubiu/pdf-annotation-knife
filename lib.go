package annotation

import (
	"math"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

var (
	epsilon float32 = 1e-6
)

func IsZeroEpsilon(f float32) bool {
	return float32(math.Abs(float64(f))) < epsilon
}
