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

		for employee := 0; employee < k; employee++ {
			var sign string
			var temperature int

			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				fmt.Println("Error reading temperature requirement")
				return
			}
		}
	}
}
