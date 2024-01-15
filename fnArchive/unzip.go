package fnArchive

import (
	"archive/zip"
	"context"
	"github.com/d3v-friends/go-pure/fnLogger"
	"io"
	"os"
	"path"
)

func Unzip(
	ctx context.Context,
	zipPath string,
	unzipPath string,
) (fileList []string, err error) {
	defer func() {
		if err != nil {
			fileList = nil
		}
	}()

	logger := fnLogger.Get(ctx, fnLogger.NewDefaultLogger())
	logger = logger.WithFields(fnLogger.Fields{
		"zipPath":   zipPath,
		"unzipPath": unzipPath,
	})

	fileList = make([]string, 0)
	var reader *zip.ReadCloser
	if reader, err = zip.OpenReader(zipPath); err != nil {
		return
	}

	defer func() {
		if closeErr := reader.Close(); err != nil {
			logger.Fatal(closeErr.Error())
		}

		if err != nil {
			logger.Warn(err.Error())
		}
	}()

	for _, file := range reader.File {

		unzipFilePath := path.Join(unzipPath, file.Name)

		if err = os.MkdirAll(path.Dir(unzipFilePath), os.ModePerm); err != nil {
			return
		}

		if file.FileInfo().IsDir() {
			continue
		}

		fileLogger := logger.WithFields(fnLogger.Fields{
			"unzipFile": unzipFilePath,
		})

		var destFile *os.File
		if destFile, err = os.Create(unzipFilePath); err != nil {
			return
		}

		var origin io.ReadCloser
		if origin, err = file.Open(); err != nil {
			return
		}

		if _, err = io.Copy(destFile, origin); err != nil {
			return
		}

		fileList = append(fileList, unzipFilePath)

		fileLogger.Trace("")
	}

	return
}
