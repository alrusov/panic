/*
Package panic implements a error handle
*/
package panic

import (
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
			misc.Logger("", "EM", "%v", r)

			stack := misc.GetCallStack(1)
			n := len(stack)
			show := false
			for i := 0; i < n; i++ {
				df := stack[i]

				if !show && !strings.HasPrefix(df.FuncName, "runtime.") {
					show = true
				}

				if show {
					misc.Logger("", "EM", ` at [%d] %s %s:%d`, n-i-1, df.FuncName, df.FileName, df.Line)
				}
			}

			misc.StopApp(misc.ExPanic)
			misc.Exit()
		}
	}
}

//----------------------------------------------------------------------------------------------------------------------------//
