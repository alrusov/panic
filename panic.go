/*
Package panic implements a error handle
*/
package panic

import (
	"bytes"
	"fmt"

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
	buf := new(bytes.Buffer)
	eos := ""
	n := len(stack)
	for i, df := range stack {
		buf.WriteString(fmt.Sprintf(`%s  at [%d] %s %s:%d`, eos, n-i-1, df.FuncName, df.FileName, df.Line))
		eos = misc.EOS
	}

	return string(buf.Bytes())
}

//----------------------------------------------------------------------------------------------------------------------------//
