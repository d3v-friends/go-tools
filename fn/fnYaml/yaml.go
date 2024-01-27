package fnYaml

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Save[DOC any](f *os.File, doc *DOC) error {
	var encoder = yaml.NewEncoder(f)
	return encoder.Encode(doc)
}

func Open[DOC any](fp string, doc *DOC) (err error) {
	var file *os.File
	if file, err = os.Open(fp); err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()

	var decoder = yaml.NewDecoder(file)
	return decoder.Decode(doc)
}
