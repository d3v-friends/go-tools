package fnPath

import (
	"os"
	"path/filepath"
)

type Path struct {
	ls []string
}

func NewPath(fp string) *Path {
	return &Path{
		ls: filepath.SplitList(fp),
	}
}

func (x *Path) WithWD() string {
	var cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	var ls = []string{cwd}
	ls = append(ls, x.ls...)

	return filepath.Join(ls...)
}

func (x *Path) RelativePath() string {
	return filepath.Join(x.ls...)
}
