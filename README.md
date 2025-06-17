# go-chunk-slice

A lightweight and generic Go utility that splits slices into fixed-size chunks.  
Useful for batching large requests, controlling API limits, or processing data in smaller batches.

---

## 📦 Installation

```bash
go get github.com/mutaito/go-chunk-slice
```

---

## 🔧 Usage

```go
package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/mutaito/go-chunk-slice/utils"
)



func main() {
    codes := GenerateRandomCodes(33)
    chunks := utils.ChunkSlice(codes, 5)

    for i, chunk := range chunks {
        fmt.Printf("Chunk %d: %v\n", i+1, chunk)
    }
}

// GenerateRandomCodes returns a slice of formatted random item codes
func GenerateRandomCodes(count int) []string {
    rand.Seed(time.Now().UnixNano())
    codes := make([]string, 0, count)

    for i := 0; i < count; i++ {
        n := rand.Intn(1000000)
        codes = append(codes, fmt.Sprintf("Item-%06d", n))
    }

    return codes
}
```

### 🔍 Example Output:
```
Chunk 1: [Item-834729 Item-000123 Item-972834 Item-112398 Item-348293]
Chunk 2: [Item-482930 Item-928374 Item-345678 Item-120938 Item-283746]
...
Chunk 7: [Item-123456 Item-654321 Item-908172]
```

---

## 💡 Why use this?

- ✅ **Generic:** works with any type (`[]string`, `[]PurchaseOrders`, `[]Users`, etc.)
- ✅ **Safe:** handles edge cases like `chunkSize <= 0`
- ✅ **Efficient:** preallocates memory to reduce allocations
- ✅ **Idiomatic:** pure Go, no third-party dependencies

---

## 🧪 Test

```bash
go test ./...
```

---

## 🪪 License

MIT — see the [LICENSE](./LICENSE) file for details.

---

## ✍️ Author

Carlos Morejón – [LinkedIn](https://www.linkedin.com/in/morejoncarlos/) · [Dev.to](https://dev.to/mutaito)

> Contributions, ideas and feedback are welcome! Open an issue or PR anytime.