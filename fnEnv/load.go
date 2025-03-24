package fnEnv

import (
	"bufio"
	"os"
	"strings"
)

func Load(fp string) (err error) {
	var file *os.File
	if file, err = os.Open(fp); err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()
	var reader = bufio.NewReader(file)

	for {
		var body []byte
		if body, _, err = reader.ReadLine(); err != nil {
			if err.Error() == "EOF" {
				err = nil
			}
			break
		}

		var strBody = string(body)

		if strBody == "" {
			continue
		}

		if strings.HasPrefix(strBody, "#") {
			continue
		}

		var strLs = strings.Split(strBody, "=")
		if len(strLs) != 2 {
			continue
		}

		if err = os.Setenv(strLs[0], strLs[1]); err != nil {
			return
		}
	}

	return
}

func Map() (env map[string]string) {
	env = make(map[string]string)
	var ls = os.Environ()
	for _, value := range ls {
		var str = strings.Split(value, "=")
		if len(str) != 2 {
			continue
		}
		env[str[0]] = str[1]
	}
	return
}
