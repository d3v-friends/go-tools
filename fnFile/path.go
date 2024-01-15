package fnFile

import (
	"github.com/d3v-friends/go-pure/fnPanic"
	"os"
	"path/filepath"
	"strings"
)

type Path string

func (x Path) Path() (path string, err error) {
	var pwd string
	var isAbsolutePath = strings.HasPrefix(x.String(), "/") || strings.HasPrefix(x.String(), "\\")
	if !isAbsolutePath {
		if pwd, err = os.Getwd(); err != nil {
			return
		}
	}

	var ls = filepath.SplitList(pwd)
	ls = append(ls, filepath.SplitList(x.String())...)
	path = filepath.Join(ls...)
	return
}

func (x Path) PathP() string {
	return fnPanic.OnValue(x.Path())
}

func (x Path) LinuxPath() (path string, err error) {
	if path, err = x.Path(); err != nil {
		return
	}
	path = strings.ReplaceAll(path, "\\", "/")
	return
}

func (x Path) LinuxPathP() string {
	return fnPanic.OnValue(x.LinuxPath())
}

func (x Path) WindowsPath() (path string, err error) {
	if path, err = x.Path(); err != nil {
		return
	}
	path = strings.ReplaceAll(path, "/", "\\")
	return
}

func (x Path) WindowsPathP() string {
	return fnPanic.OnValue(x.WindowsPath())
}

func (x Path) String() string {
	return string(x)
}
