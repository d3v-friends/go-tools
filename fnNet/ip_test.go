package fnNet_test

import (
	"fmt"
	"github.com/d3v-friends/go-tools/fnNet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(test *testing.T) {
	test.Run("get_ip", func(t *testing.T) {
		var ip, err = fnNet.IP()
		assert.NoError(t, err)
		fmt.Println("ip", ip)
	})
}
