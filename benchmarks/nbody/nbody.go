package main

import "fmt"

func main() {
	N := 5
	steps := 500000

	x := []float64{0.0, 4.841431, 8.343367, 12.894369, 15.379697}
	y := []float64{0.0, -1.160320, 4.124799, -15.201544, -25.632172}
	z := []float64{0.0, -0.103860, -0.403093, -0.223581, 0.179258}

	vx := []float64{0.0, 0.001660, 0.007699, 0.002965, 0.002681}
	vy := []float64{0.0, 0.007699, 0.006286, 0.002378, 0.001628}
	vz := []float64{0.0, -0.000069, -0.000169, -0.000030, -0.000015}

	mass := []float64{39.478418, 0.037694, 0.011287, 0.001724, 0.000203}

	dt := 0.01

	for step := 0; step < steps; step++ {
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				dx := x[i] - x[j]
				dy := y[i] - y[j]
				dz := z[i] - z[j]
				dist2 := dx*dx + dy*dy + dz*dz
				mag := dt / (dist2*100.0 + 0.001)

				fx := dx * mag
				fy := dy * mag
				fz := dz * mag

				vx[i] -= fx * mass[j]
				vy[i] -= fy * mass[j]
				vz[i] -= fz * mass[j]
				vx[j] += fx * mass[i]
				vy[j] += fy * mass[i]
				vz[j] += fz * mass[i]
			}
		}

		for i := 0; i < N; i++ {
			x[i] += dt * vx[i]
			y[i] += dt * vy[i]
			z[i] += dt * vz[i]
		}
	}

	energy := 0.0
	for i := 0; i < N; i++ {
		energy += 0.5 * mass[i] * (vx[i]*vx[i] + vy[i]*vy[i] + vz[i]*vz[i])
	}
	fmt.Println(energy)
}