package text

import (
	"fmt"
	"io"
	"os"
)

func PrintStart(s string, args ...any) {
	doPrint(os.Stdout, "🚀", s, args...)
}

func PrintSuccess(s string, args ...any) {
	doPrint(os.Stdout, "✅ ", s, args...)
}

func PrintError(w io.Writer, s string, args ...any) {
	doPrint(w, "❌ ", s, args...)
}

func doPrint(w io.Writer, prefix, s string, args ...any) {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	_, _ = fmt.Fprintf(w, "%s %s\n", prefix, s)
}
