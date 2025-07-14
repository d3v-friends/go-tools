package fnArchive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// OpenZip
// 나중에 Recursive 하게 로딩하는 것 추가 하기
func OpenZip(filename, outputPath string) (fileList []string, err error) {
	var zipFile *zip.ReadCloser
	if zipFile, err = zip.OpenReader(filename); err != nil {
		return
	}

	defer func() {
		_ = zipFile.Close()
	}()

	_ = os.RemoveAll(outputPath)

	if err = os.MkdirAll(outputPath, os.ModePerm); err != nil {
		return
	}

	for _, unzipFile := range zipFile.File {
		var unzipFilename = filepath.Join(outputPath, unzipFile.Name)

		if strings.HasPrefix(filepath.Base(unzipFilename), ".") {
			continue
		}

		if unzipFile.FileInfo().IsDir() {
			continue
		}

		_ = os.MkdirAll(filepath.Dir(unzipFilename), os.ModePerm)

		var dstFile *os.File
		if dstFile, err = os.Create(unzipFilename); err != nil {
			return
		}

		var archive io.ReadCloser
		if archive, err = unzipFile.Open(); err != nil {
			return
		}

		if _, err = io.Copy(dstFile, archive); err != nil {
			return
		}

		if filepath.Ext(unzipFilename) == ".zip" {
			var ls []string
			if ls, err = OpenZip(unzipFilename, filepath.Dir(unzipFilename)); err != nil {
				return
			}

			fileList = append(fileList, ls...)
		} else {
			fileList = append(fileList, unzipFilename)
		}

		_ = dstFile.Close()
		_ = archive.Close()
	}

	return
}
