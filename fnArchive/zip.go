package fnArchive

import (
	"archive/zip"
	"context"
	"github.com/d3v-friends/go-pure/fnLogger"
	"io"
	"os"
	"path/filepath"
)

func Zip(ctx context.Context, zipPath string, fileList []string) (err error) {
	logger := fnLogger.Get(ctx)
	logger = logger.WithFields(fnLogger.Fields{
		"fn":      "fnArchive.Zip()",
		"zipFile": zipPath,
	})

	if err = os.MkdirAll(filepath.Dir(zipPath), os.ModePerm); err != nil {
		return
	}

	var file *os.File
	if file, err = os.Create(zipPath); err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()

	writer := zip.NewWriter(file)
	defer func() {
		_ = writer.Close()
	}()

	for _, fp := range fileList {
		fileLogger := logger.WithFields(fnLogger.Fields{
			"file": fp,
		})

		var sv io.Writer
		if sv, err = writer.Create(filepath.Base(fp)); err != nil {
			return
		}

		var f *os.File
		if f, err = os.Open(fp); err != nil {
			return
		}

		if _, err = io.Copy(sv, f); err != nil {
			return
		}

		_ = f.Close()

		fileLogger.Trace("zipped")
	}

	logger.Trace("zip done")

	return
}
