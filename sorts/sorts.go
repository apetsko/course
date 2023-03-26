package sorts

func BubbleSort(sl []int) []int {
	for i := 1; i < len(sl); i++ {
		ok := true
		for j := 0; j < len(sl)-i; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
				ok = false
			}
		}
		if ok {
			break
		}
	}
	return sl
}

func MergeSort(sl []int) []int {
	size := len(sl)
	if size <= 1 {
		return sl
	}
	L := MergeSort(sl[:size/2])
	R := MergeSort(sl[size/2:])

	n, m, k := 0, 0, 0
	C := make([]int, len(L)+len(R))

	for n < len(L) && m < len(R) {
		if L[n] <= R[m] {
			C[k] = L[n]
			n += 1

		} else {
			C[k] = R[m]
			m += 1
		}
		k += 1
	}

	for n < len(L) {
		C[k] = L[n]
		n += 1
		k += 1
	}
	for m < len(R) {
		C[k] = R[m]
		m += 1
		k += 1
	}
	for i := range sl {
		sl[i] = C[i]
	}
	return sl
}

func InsertionSort(sl []int) []int {
	for j := 1; j < len(sl); j++ {
		key := sl[j]
		i := j - 1
		for i >= 0 && sl[i] > key {
			sl[i+1] = sl[i]
			i--
		}
		sl[i+1] = key
	}
	return sl
}
