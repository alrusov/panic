/*
Package panic implements a error handle
*/
package panic

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

var (
	enabled = true
)

//----------------------------------------------------------------------------------------------------------------------------//

// Enable --
func Enable() {
	enabled = true
}

// Disable --
func Disable() {
	enabled = false
}

func init() {
	enabled = !misc.IsDebug()
}

//----------------------------------------------------------------------------------------------------------------------------//

// SaveStackToLog - internal panic function
func SaveStackToLog() {
	if enabled {
		r := recover()
		if r != nil {
			misc.Logger("", "EM", "%v\n%s", r, GetStack())
			misc.StopApp(misc.ExPanic)
			misc.Exit()
		}
	}
}

//----------------------------------------------------------------------------------------------------------------------------//

// GetStack -
func GetStack() string {
	stack := misc.GetCallStack(1)
	n := len(stack)
	buf := new(bytes.Buffer)
	show := false
	lf := ""

	for i := 0; i < n; i++ {
		df := stack[i]

		if !show && !strings.HasPrefix(df.FuncName, "runtime.") {
			show = true
		}

		if show {
			buf.WriteString(fmt.Sprintf(`%s  at [%d] %s %s:%d`, lf, n-i-1, df.FuncName, df.FileName, df.Line))
			lf = "\n"
		}
	}

	return string(buf.Bytes())
}

//----------------------------------------------------------------------------------------------------------------------------//
