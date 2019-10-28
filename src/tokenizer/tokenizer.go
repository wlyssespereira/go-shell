package tokenizer

import (
	"scanner"
	"strings"
)

func Tokenize(str string) []string {
	var lst []string

	codeReader := strings.NewReader(str)
	var scn scanner.Scanner
	scn.Init(codeReader)

	tok := scn.Scan()
	for tok != scanner.EOF {
		lst = append(lst, scn.TokenText())
		tok = scn.Scan()
	}

	return lst
}
