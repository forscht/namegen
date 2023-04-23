package dictionaries

import (
    "fmt"
    "math"
    "math/rand"
)

func Numbers(min int, max int, length *int) []string {

    if length != nil {
        length := math.Pow(10, float64(*length))
        min = int(length / 10)
        max = int(length - 1)
        return []string{fmt.Sprintf("%d", int(math.Floor(float64(rand.Intn(max-min)+min))))}
    }

    return []string{fmt.Sprintf("%d", int(math.Floor(float64(rand.Intn(max-min)+min))))}
}
