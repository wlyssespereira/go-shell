package lexical

import (
	"fmt"
)

type TokenType string

const (
	UNDEFINED = "UNDEFINED"

	WORD            = "WORD"
	ASSIGNMENT_WORD = "ASSIGNMENT_WORD"
	NAME            = "NAME"
	NEWLINE         = "NEWLINE"
	IO_NUMBER       = "IO_NUMBER"

	AND_IF = "AND_IF" // '&&'
	OR_IF  = "OR_IF"  // '||'
	DSEMI  = "DSEMI"  // ';;'

	DLESS     = "DLESS"     // '<<'
	DGREAT    = "DGREAT"    // '>>'
	LESSAND   = "LESSAND"   // '<&'
	GREATAND  = "GREATAND"  // '>&'
	LESSGREAT = "LESSGREAT" // '<>'
	DLESSDASH = "DLESSDASH" // '<<-'

	CLOBBER = "CLOBBER" // '>|'

	IF   = "IF"   // if
	THEN = "THEN" // then
	ELSE = "ELSE" // else
	ELIF = "ELIF" // elif
	FI   = "FI"   // fi
	DO   = "DO"   // do
	DONE = "DONE" // done

	CASE  = "CASE"  // case
	ESAC  = "ESAC"  // esac
	WHILE = "WHILE" // while
	UNTIL = "UNTIL" // until
	FOR   = "FOR"   // for

	LBRACE = "LBRACE" // '{'
	RBRACE = "RBRACE" // '}'
	BANG   = "BANG"   // '!'

	IN = "IN" // in
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
