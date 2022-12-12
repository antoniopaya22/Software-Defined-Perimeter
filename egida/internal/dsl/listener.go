package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener struct {
	Errors int
}

func (l *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool,
	ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
	conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex,
	prediction int, configs antlr.ATNConfigSet) {
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {
	fmt.Println("SyntaxError on line", line, ":", column, "-->", msg)
	l.Errors = l.Errors + 1
}
