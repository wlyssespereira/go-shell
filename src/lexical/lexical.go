package lexical

import (
	"regexp"
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
	List   []string
	Tokens []Token
	Errors []Token
}

func NewAnalyzer() *LexicalAnalyzer {
	analyzer := new(LexicalAnalyzer)
	return analyzer
}

func (analyzer *LexicalAnalyzer) Analyze() bool {
	for _, text := range analyzer.List {
		var token = new(Token)
		token.Value = text
		token.Type = analyzer.DetermineTypeToken(text)
		if token.Type == UNDEFINED {
			analyzer.Errors = append(analyzer.Errors, *token)
		} else {
			analyzer.Tokens = append(analyzer.Tokens, *token)
		}
	}
	return len(analyzer.Errors) == 0
}

func (analyzer LexicalAnalyzer) DetermineTypeToken(value string) TokenType {
	var word = regexp.MustCompile(`^\S+$`)

	switch {
	case value == "&&":
		return AND_IF
	case value == "||":
		return OR_IF
	case value == ";;":
		return DSEMI
	case value == "<<":
		return DLESS
	case value == ">>":
		return DGREAT
	case value == "<&":
		return LESSAND
	case value == ">&":
		return GREATAND
	case value == "<>":
		return LESSGREAT
	case value == "<<-":
		return DLESSDASH
	case value == ">|":
		return CLOBBER
	case value == "if":
		return IF
	case value == "then":
		return THEN
	case value == "else":
		return ELSE
	case value == "elif":
		return ELIF
	case value == "fi":
		return FI
	case value == "do":
		return DO
	case value == "done":
		return DONE
	case value == "case":
		return CASE
	case value == "esac":
		return ESAC
	case value == "while":
		return WHILE
	case value == "until":
		return UNTIL
	case value == "for":
		return FOR
	case value == "{":
		return LBRACE
	case value == "}":
		return RBRACE
	case value == "!":
		return BANG
	case value == "in":
		return IN
	case word.MatchString(value):
		return WORD
	default:
		return UNDEFINED
	}
}
