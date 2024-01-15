package fnEnv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile(test *testing.T) {

	test.Run("read", func(t *testing.T) {
		var err error
		if err = ReadFromFile("./.env"); err != nil {
			test.Fatal(err)
		}

		assert.Equal(t, "10000", Read("PORT"))
		assert.Equal(t, "www.mongodb.com", Read("MG_HOST"))
	})
}
