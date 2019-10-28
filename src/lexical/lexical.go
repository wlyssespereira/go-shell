package lexical

import (
	"fmt"
)

type TokenType int

const (
	UNDEFINED TokenType = iota

	WORD
	ASSIGNMENT_WORD
	NAME
	NEWLINE
	IO_NUMBER

	AND_IF // '&&'
	OR_IF  // '||'
	DSEMI  // ';;'

	DLESS     // '<<'
	DGREAT    // '>>'
	LESSAND   // '<&'
	GREATAND  // '>&'
	LESSGREAT // '<>'
	DLESSDASH // '<<-'

	CLOBBER // '>|'

	IF   // if
	THEN // then
	ELSE // else
	ELIF // elif
	FI   // fi
	DO   // do
	DONE // done

	CASE  // case
	ESAC  // esac
	WHILE // while
	UNTIL // until
	FOR   // for

	LBRACE // '{'
	RBRACE // '}'
	BANG   // '!'

	IN // in
)

type Token struct {
	Value string
	Type  TokenType
}

type LexicalAnalyzer struct {
	Tokens []string
}

func (analyzer LexicalAnalyzer) Analyze() {

	for _, token := range analyzer.Tokens {
		x := new(Token)
		value := token
		x.Value = value
		x.Type = DetermineTypeToken(value)
		fmt.Println(x)
	}

}

func DetermineTypeToken(value string) TokenType {
	switch value {
	case "&&":
		return AND_IF
	case "||":
		return OR_IF
	case ";;":
		return DSEMI
	case "<<":
		return DLESS
	case ">>":
		return DGREAT
	case "<&":
		return LESSAND
	case ">&":
		return GREATAND
	case "<>":
		return LESSGREAT
	case "<<-":
		return DLESSDASH
	case ">|":
		return CLOBBER
	case "if":
		return IF
	case "then":
		return THEN
	case "else":
		return ELSE
	case "elif":
		return ELIF
	case "fi":
		return FI
	case "do":
		return DO
	case "done":
		return DONE
	case "case":
		return CASE
	case "esac":
		return ESAC
	case "while":
		return WHILE
	case "until":
		return UNTIL
	case "for":
		return FOR
	case "{":
		return LBRACE
	case "}":
		return RBRACE
	case "!":
		return BANG
	case "in":
		return IN
	default:
		return UNDEFINED
	}
}
