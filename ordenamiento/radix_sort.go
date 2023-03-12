package ordenamiento

import (
	"math"
)

// Creado por Bing AI
// RadixSort sorts an array of integers using radix sort
// It assumes that the numbers are in base 10
// It returns a pointer to the sorted array
// It is safe to be called using goroutines
func RadixSort(arr *[]int) {
	// Find the maximum number in the array
	max := 0
	for _, num := range *arr {
		if num > max {
			max = num
		}
	}

	// Find the number of digits in the maximum number
	digits := int(math.Log10(float64(max))) + 1

	// Create a slice of 10 buckets, each bucket is a slice of integers
	buckets := make([][]int, 10)

	// Loop from the least significant digit to the most significant digit
	for i := 0; i < digits; i++ {
		// Loop through the array and put each number into the corresponding bucket
		// based on the current digit
		for _, num := range *arr {
			// Get the current digit using modulo and division
			digit := (num / int(math.Pow10(i))) % 10
			// Append the number to the bucket
			buckets[digit] = append(buckets[digit], num)
		}

		// Reset the array index
		index := 0

		// Loop through the buckets and put the numbers back into the array
		// in the order of the buckets
		for j := 0; j < 10; j++ {
			// Loop through the bucket
			for _, num := range buckets[j] {
				// Put the number into the array
				(*arr)[index] = num
				// Increment the array index
				index++
			}
			// Clear the bucket
			buckets[j] = nil
		}
	}

}
