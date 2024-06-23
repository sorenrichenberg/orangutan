// tracing/tracing.go

package tracing

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func indentLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fullStr string) {
	fmt.Printf("%s%s\n", indentLevel(), fullStr)
}

func incIdent() { traceLevel += 1 }
func decIdent() { traceLevel -= 1 }

func Trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func UnTrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
