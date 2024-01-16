package fnEnv

import (
	"fmt"
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
