package quicksort

func quicksort(arr []int) []int {
	return qsort(arr, 0, len(arr)-1)
}

func qsort(arr []int, low, high int) []int {
	if low >= 0 && high >= 0 && low < high {
		pi := partitionHoare(arr, low, high)
		qsort(arr, low, pi-1)
		qsort(arr, pi+1, high)
	}

	return arr
}

func partitionHoare(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low+1, high

	for true { // Basic do-while loop
		// Move i to the right until an element greater than the pivot is found
		for i < j && arr[i] <= pivot {
			i++
		}

		// Move j to the left until an element less than or equal to the pivot is found
		for j >= i && arr[j] > pivot {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}

	// Swap pivot into final position
	arr[low], arr[j] = arr[j], arr[low]

	return j
}
