package fnFile

import (
	"os"
	"path/filepath"
)

type PathBuilder struct {
	elems []string
}

func NewPathBuilder(root ...string) (res *PathBuilder, err error) {
	res = &PathBuilder{
		elems: make([]string, 0),
	}

	if len(root) == 0 {
		return
	}

	for _, p := range root {
		res.elems = append(res.elems, filepath.SplitList(p)...)
	}

	return
}

func NewPathBuilderWithWD(root ...string) (res *PathBuilder, err error) {
	var wd string
	if wd, err = os.Getwd(); err != nil {
		return
	}

	var ls = make([]string, 0)
	ls = append(ls, wd)
	ls = append(ls, root...)

	return NewPathBuilder(ls...)
}

func (x *PathBuilder) Join(path ...string) {
	for _, elem := range path {
		x.elems = append(x.elems, filepath.SplitList(elem)...)
	}
}

func (x *PathBuilder) String() string {
	return filepath.Join(x.elems...)
}
