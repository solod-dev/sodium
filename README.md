# sodium

Solod bindings for [Sodium](https://libsodium.gitbook.io/doc), a modern, easy-to-use library for encryption, decryption, signatures and password hashing.

## Usage

1. Install libsodium for your operating system.

2. Install the Solod bindings.

```
go get solod.dev/sodium@latest
```

3. Use it in your code.

```go
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/os"
	"solod.dev/sodium/libsodium"
)

func main() {
	if libsodium.Init() < 0 {
		fmt.Println("libsodium could not be initialized")
		os.Exit(1)
	}
	fmt.Println("libsodium is ready")
}
```

## Examples

[String hash](example/hash/main.go)
