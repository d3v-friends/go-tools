package fnCrypt

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAes(test *testing.T) {
	const str = "hello_world_golang"
	var sec1 = NewSecret(32)
	var sec2 = NewSecret(32)

	test.Run("secret", func(t *testing.T) {
		var try = 20

		for i := 0; i < try; i++ {
			var secret = NewSecret(32)
			fmt.Printf("secret: %s\n", secret)
			assert.Equal(t, 32, len(secret))
		}

	})

	test.Run("simple", func(t *testing.T) {
		var enc = fnPanic.Get(EncryptAES256(sec1, str))

		fmt.Printf("env: %x\n", enc)

		var dec = fnPanic.Get(DecryptAES256(sec1, enc))
		fmt.Printf("dec: %s\n", str)
		assert.Equal(t, str, dec)
	})

	test.Run("invalid secret", func(t *testing.T) {
		var enc1 = fnPanic.Get(EncryptAES256(sec1, str))
		var enc2 = fnPanic.Get(EncryptAES256(sec2, str))

		assert.NotEqual(t, enc1, enc2)
	})
}
