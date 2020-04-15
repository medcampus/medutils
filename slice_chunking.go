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

func IntSliceChunking(sliceString []int, chunkSize int) [][]int {
	var chunks [][]int

	for i := 0; i < len(sliceString); i += chunkSize {
		end := i + chunkSize

		if end > len(sliceString) {
			end = len(sliceString)
		}

		chunks = append(chunks, sliceString[i:end])
	}
	return chunks
}
