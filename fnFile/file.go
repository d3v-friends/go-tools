package fnFile

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnParams"
	"os"
	"path/filepath"
)

func New(fp string, force ...bool) (file *os.File, err error) {
	isForce := fnParams.Get(force)

	if err = os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
		return
	}

	if _, err = os.Stat(fp); err == nil {
		if !isForce {
			err = fmt.Errorf("already exist file: fp=%s", fp)
			return
		}
		if err = os.Remove(fp); err != nil {
			return
		}
	}

	file, err = os.Create(fp)

	return
}

func Open(fp string) (file *os.File, err error) {
	if !isExist(fp) {
		return nil, fmt.Errorf("not found file: fp=%s", fp)
	}
	return os.Open(fp)
}

func isExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}
