package fnArchive

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestZip(test *testing.T) {
	var wd string
	var err error
	if wd, err = os.Getwd(); err != nil {
		test.Fatal(err)
	}

	test.Run("zip", func(t *testing.T) {
		fileList := []string{
			filepath.Join(wd, "temp/a1.txt"),
		}

		ctx := context.TODO()
		if err = Zip(ctx, filepath.Join(wd, "test.zip"), fileList); err != nil {
			test.Fatal(err)
		}
	})
}
