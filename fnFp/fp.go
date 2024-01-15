package fnFp

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnEnv"
	"os"
	"path/filepath"
	"strings"
)

type Fp string

var GOOS = fnEnv.Read("GOOS")

func NewFp(v string) Fp {
	return Fp(v)
}

func (x Fp) Mkdir() error {
	return os.MkdirAll(filepath.Dir(x.Path()), os.ModePerm)
}

func (x Fp) Path() string {
	switch GOOS {
	case "windows":
		return strings.ReplaceAll(x.String(), "/", "\\")
	case "darwin", "linux":
		return strings.ReplaceAll(x.String(), "\\", "/")
	default:
		var err = fmt.Errorf("unsupported os: os=%s", GOOS)
		panic(err)
	}
}

func (x Fp) PathWithWD() string {
	var err error
	if strings.HasPrefix(x.String(), "/") {
		err = fmt.Errorf("fp is absolute path: fp=%s", x)
		panic(err)
	}

	var wd string
	if wd, err = os.Getwd(); err != nil {
		panic(err)
	}

	return Fp(filepath.Join(wd, x.String())).Path()
}

func (x Fp) String() string {
	return string(x)
}

func (x Fp) Dirname() string {
	return filepath.Dir(x.String())
}

func (x Fp) Filename() string {
	var str = x.String()
	return fmt.Sprintf(
		"%s%s",
		filepath.Base(str),
		filepath.Ext(str),
	)
}
