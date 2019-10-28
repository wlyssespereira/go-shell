package main

import (
	"fmt"
	"lexical"
	"os"
	"strings"
	"tokenizer"
)

func main() {
	fmt.Println(os.Args)
	str := ""
	str = "TESTE=/etc/hosts"
	str = "TESTE >> TESTE"
	str = "cat /etc/hosts | grep localhost | od \"teste teste\" "

	lst := strings.Fields(str)
	lst = tokenizer.Tokenize(str)

	analyzer := new(lexical.LexicalAnalyzer)
	analyzer.Tokens = lst
	analyzer.Analyze()

}
