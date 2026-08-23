package libsodium_test

import (
	"solod.dev/so/c"
	"solod.dev/sodium/libsodium"
)

func ExampleGenerichash() {
	libsodium.Init()

	msg := "Hello, World!"
	hash := make([]byte, libsodium.Generichash_BYTES)
	libsodium.Generichash(
		c.SliceData[c.UChar](hash), c.Size(len(hash)),
		c.StringData[c.UChar](msg), c.ULongLong(len(msg)),
		nil, 0)
}
