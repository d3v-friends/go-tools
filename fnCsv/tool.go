package fnCsv

import "fmt"

var Bom = string([]byte{0xEF, 0xBB, 0xBF})

func QuoteString(str string) string {
	return fmt.Sprintf("=\"%s\"", str)
}
