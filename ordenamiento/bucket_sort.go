package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func BucketSort(arr *[]int) {
	n := len(*arr)
	if n <= 1 {
		return
	}

	// Find maximum element in array
	maxVal := (*arr)[0]
	for i := 1; i < n; i++ {
		if (*arr)[i] > maxVal {
			maxVal = (*arr)[i]
		}
	}

	// Create empty buckets
	bucketSize := 10
	buckets := make([][]int, bucketSize)

	// Add elements to the buckets
	for i := 0; i < n; i++ {
		bucketIndex := ((*arr)[i] * bucketSize) / (maxVal + 1)
		buckets[bucketIndex] = append(buckets[bucketIndex], (*arr)[i])
	}

	// Sort the elements within each bucket using insertion sort
	for i := 0; i < bucketSize; i++ {
		insertionSort(&buckets[i])
	}

	// Concatenate the sorted buckets into the original array
	k := 0
	for i := 0; i < bucketSize; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			(*arr)[k] = buckets[i][j]
			k++
		}
	}
}

func insertionSort(bucket *[]int) {
	n := len(*bucket)

	for i := 1; i < n; i++ {
		key := (*bucket)[i]
		j := i - 1

		for j >= 0 && (*bucket)[j] > key {
			(*bucket)[j+1] = (*bucket)[j]
			j--
		}

		(*bucket)[j+1] = key
	}
}
