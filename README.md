# Simple base64 golang implementation

Usage:

```go
...
import "github.com/matmazurk/base64"
...

input := []byte("some input")
encoded := base64.Encode(input)

decoded, err := base64.Decode(encoded)
```

To run tests:

```bash
go test .
```

To run fuzz test:

```bash
go test -fuzz .
```
