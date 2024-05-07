package utils

import "fmt"

func Avg(arr []float64) string {
	if len(arr) == 0 {
		return "NaN"
	}

	sum := 0.0
	for _, v := range arr {
		sum += v
	}

	average := sum / float64(len(arr))
	return fmt.Sprintf("%.2f", average)
}

func Max(arr []float64) string {
	m := arr[0]

	for _, v := range arr {
		if m < v {
			m = v
		}
	}

	return fmt.Sprintf("%d", m)
}

func Min(arr []float64) string {
	m := arr[0]

	for _, v := range arr {
		if m < v {
			m = v
		}
	}

	return fmt.Sprintf("%d", m)
}
