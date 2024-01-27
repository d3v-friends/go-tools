package fnEnv

import (
	"bufio"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)

func GetString(key string) (res string) {
	res = os.Getenv(key)
	if res == "" {
		panic(fmt.Errorf("not found env: key=%s", key))
	}
	return
}

func GetBool(key string) (res bool) {
	return GetString(key) == "true"
}

func GetInt(key string) (res int) {
	var err error
	if res, err = strconv.Atoi(GetString(key)); err != nil {
		panic(err)
	}
	return
}

func GetDecimal(key string) decimal.Decimal {
	return fnPanic.Get(decimal.NewFromString(GetString(key)))
}

func Load(fp string) (err error) {
	var file *os.File
	if file, err = os.Open(fp); err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()
	var reader = bufio.NewReader(file)

	for {
		var body []byte
		if body, _, err = reader.ReadLine(); err != nil {
			if err.Error() == "EOF" {
				err = nil
			}
			break
		}

		var strBody = string(body)

		if strBody == "" {
			continue
		}

		if strings.HasPrefix(strBody, "#") {
			continue
		}

		var strLs = strings.Split(strBody, "=")
		if len(strLs) != 2 {
			continue
		}

		if err = os.Setenv(strLs[0], strLs[1]); err != nil {
			return
		}
	}

	return
}
