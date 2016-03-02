package simulation

import "math"
import "math/rand"

var generate bool = false

func GenerateGaussianNoise(mu, sigma float64) float64 {
	epsilon := math.SmallestNonzeroFloat64
	two_pi := 2.0 * math.Pi

	var z0, z1 float64

	generate = !generate;

	if (!generate) {
		return z1 * sigma + mu;
	}
	var u1, u2 float64
	for ok := true; ok; ok = (u1 <= epsilon) {
		u1 = rand.Float64()
		u2 = rand.Float64()

	}

	z0 = math.Sqrt(-2.0 * math.Log(u1)) * math.Cos(two_pi * u2);
	z1 = math.Sqrt(-2.0 * math.Log(u1)) * math.Sin(two_pi * u2);
	return z0 * sigma + mu;
}
