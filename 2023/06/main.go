package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	records := map[float64]float64{}
	records[38] = 241
	records[94] = 1549
	records[79] = 1074
	records[70] = 1091

	calculateMargin(records)

	records2 := map[float64]float64{}
	records2[38947970] = 241154910741091

	calculateMargin(records2)
}

func calculateMargin(records map[float64]float64) {
	margin := 1
	for time, distance := range records {
		x1, x2, err := quadratic_formula(1, -(time), distance)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("time:", time, "distance:", distance, "x1:", x1, "x2:", x2)

		// make sure to round up
		minX := max(int(math.Round(x1+0.5)), 0)
		maxX := min(int(math.Round(x2+0.5)), int(time))

		solutions := maxX - minX
		fmt.Println("min:", minX, "max:", maxX, "num:", solutions)

		margin = margin * solutions
	}

	fmt.Println("margin:", margin)
}

func quadratic_formula(a float64, b float64, c float64) (float64, float64, error) {
	var x1, x2, discriminant float64

	discriminant = math.Pow(b, 2) - (4 * a * c)

	if discriminant > 0 {
		x1 = (-b - math.Sqrt(discriminant)) / (2 * a)
		x2 = (-b + math.Sqrt(discriminant)) / (2 * a)
	} else if discriminant == 0 {
		x1 = -b / (2 * a)
		x2 = -b / (2 * a)
	} else if discriminant < 0 {
		return 0, 0, errors.New("cannot solve")
	}

	return x1, x2, nil
}
