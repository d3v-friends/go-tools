package fnFilepath

import (
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/d3v-friends/go-pure/fnReflect"
	"os"
	"path/filepath"
	"strings"
)

type Filepath string

func NewFilepath(v string) *Filepath {
	return fnReflect.ToPointer(Filepath(v))
}

func (x *Filepath) String() string {
	var ls = filepath.SplitList(string(*x))
	return filepath.Join(ls...)
}
func (x *Filepath) AbsolutePath() (path string) {
	var str = x.String()
	var ls = filepath.SplitList(str)
	if !strings.HasPrefix(str, "/") {
		var pwd = filepath.SplitList(fnPanic.Get(os.Getwd()))
		ls = append(pwd, ls...)
	}
	return filepath.Join(ls...)
}

func (x *Filepath) AppendPrefix(prefix string) {
	var pathLs = filepath.SplitList(x.String())
	var fixLs = filepath.SplitList(prefix)
	*x = Filepath(filepath.Join(append(fixLs, pathLs...)...))
}

func (x *Filepath) AppendSuffix(suffix string) {
	var pathLs = filepath.SplitList(x.String())
	var fixLs = filepath.SplitList(suffix)
	*x = Filepath(filepath.Join(append(fixLs, pathLs...)...))
}

func (x *Filepath) Mkdir() (err error) {
	return os.MkdirAll(filepath.Dir(x.String()), os.ModePerm)
}
