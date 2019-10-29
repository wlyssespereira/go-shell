package main

import (
	"fmt"
	"lexical"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	str := ""
	str = "TESTE=/etc/hosts"
	str = "TESTE >> TESTE"
	str = "cat /etc/hosts | grep localhost | od \"teste teste\" & "

	lst := strings.Fields(str)

	//analyzer := new(lexical.LexicalAnalyzer)
	analyzer := lexical.NewAnalyzer()
	analyzer.List = lst
	analyzer.Analyze()

	fmt.Printf("%v", analyzer.Tokens)

}
