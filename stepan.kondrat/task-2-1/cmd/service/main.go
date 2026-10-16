package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading department count")
		return
	}

	for department := 0; department < n; department++ {
		var k int

		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println("Error reading employee count")
			return
		}

		lower, upper := 15, 30

		for employee := 0; employee < k; employee++ {
			var sign string
			var temperature int

			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				fmt.Println("Error reading temperature requirement")
				return
			}

			switch sign {
			case ">=":
				if temperature > lower {
					lower = temperature
				}
			case "<=":
				if temperature < upper {
					upper = temperature
				}
			default:
				fmt.Println("Incorrect input data format")
				return
			}

			if lower <= upper {
				fmt.Println(lower)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
