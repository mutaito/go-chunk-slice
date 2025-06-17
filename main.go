package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-chunk-slice/utils"
)

func main() {
	//return []string
	codes := GenerateRandomCodes(33)

	chunks := utils.ChunkSlice(codes, 5)

	for i, chunk := range chunks {
		fmt.Printf("Chunk %d: %v\n", i+1, chunk)
	}
}

func GenerateRandomCodes(count int) []string {
	rand.Seed(time.Now().UnixNano())

	codes := make([]string, 0, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(1000000)
		codes = append(codes, fmt.Sprintf("Item-%06d", n))
	}
	return codes
}
