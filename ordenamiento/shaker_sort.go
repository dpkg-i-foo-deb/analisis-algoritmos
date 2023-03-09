package ordenamiento

//Generado por ChatGPT
func ShakerSort(arr []int) {
    left := 0
    right := len(arr) - 1
    for left < right {
        // Move the largest element to the right
        for i := left; i < right; i++ {
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
            }
        }
        right--

        // Move the smallest element to the left
        for i := right; i > left; i-- {
            if arr[i] < arr[i-1] {
                arr[i], arr[i-1] = arr[i-1], arr[i]
            }
        }
        left++
    }
}
