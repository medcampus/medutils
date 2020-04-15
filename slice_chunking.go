package medutils

func StringSliceChunking(sliceString []string, chunkSize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(sliceString); i += chunkSize {
		end := i + chunkSize

		if end > len(sliceString) {
			end = len(sliceString)
		}

		chunks = append(chunks, sliceString[i:end])
	}
	return chunks
}

func IntSliceChunking(sliceInt []int, chunkSize int) [][]int {
	var chunks [][]int

	for i := 0; i < len(sliceInt); i += chunkSize {
		end := i + chunkSize

		if end > len(sliceInt) {
			end = len(sliceInt)
		}

		chunks = append(chunks, sliceInt[i:end])
	}
	return chunks
}

func SliceChunking(slice []interface{}, chunkSize int) [][]interface{} {
	var chunks [][]interface{}

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
