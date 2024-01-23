package fnEnv

import (
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
)

func Read(key string) (res string) {
	res = os.Getenv(key)
	if res == "" {
		panic(fmt.Errorf("not found env: key=%s", key))
	}
	return
}

func ReadBool(key string) (res bool) {
	return Read(key) == "true"
}

func ReadInt(key string) (res int) {
	var err error
	if res, err = strconv.Atoi(Read(key)); err != nil {
		panic(err)
	}
	return
}

func ReadDecimal(key string) decimal.Decimal {
	return fnPanic.Get(decimal.NewFromString(Read(key)))
}
