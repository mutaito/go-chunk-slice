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
    "github.com/mutaito/go-chunk-slice"
)

func main() {
    //TODO Add two example implements
}
```

Output A:
```
Chunk 0: []
Chunk 1: []
Chunk 2: []
```
Output B:
```
Chunk 0: []
Chunk 1: []
Chunk 2: []
```

---

## 💡 Why use this?

- ✅ Generic: works with any type (`[]string`, `[]PurchaseOrders`, `[]Users`, etc.)
- ✅ Safe: handles edge cases like `chunkSize <= 0`
- ✅ Efficient: preallocates memory for better performance
- ✅ Idiomatic: pure Go, no dependencies

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
