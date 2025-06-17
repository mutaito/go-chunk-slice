package utils

func ChunkSlice[T any](data []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}

	n := len(data)
	chunks := make([][]T, 0, (n+chunkSize-1)/chunkSize)

	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		chunks = append(chunks, data[i:end])
	}

	return chunks
}
