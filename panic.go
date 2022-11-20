/*
Package panic implements a error handle
*/
package panic

import (
	"bytes"
	"fmt"
	"runtime"
	"sync/atomic"

	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

var (
	enabled   = true
	dumpStack = false
	id        = uint64(0)
)

//----------------------------------------------------------------------------------------------------------------------------//

func init() {
	enabled = !misc.IsDebug()
}

//----------------------------------------------------------------------------------------------------------------------------//

// Enable --
func Enable() {
	enabled = true
}

// Disable --
func Disable() {
	enabled = false
}

// SetDumpStack --
func SetDumpStack(enable bool) {
	dumpStack = enable
}

//----------------------------------------------------------------------------------------------------------------------------//

// ID --
func ID() uint64 {
	i := atomic.AddUint64(&id, 1)
	if dumpStack {
		misc.Logger("", "AL", "[panicID %d] Assigned to %s", i, GetStack())
	}
	return i
}

//----------------------------------------------------------------------------------------------------------------------------//

// SaveStackToLog - internal panic function
func SaveStackToLog() {
	SaveStackToLogEx(0)
}

// SaveStackToLogEx - internal panic function
func SaveStackToLogEx(id uint64, details ...any) {
	if enabled {
		r := recover()
		if r != nil {
			misc.Logger("", "EM", "[panicID %d] %v\n%s", id, r, GetStack())

			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			misc.Logger("", "IN", "AllocSys %d, HeapSys %d, HeapInuse: %d, HeapObjects %d, StackSys: %d, StackInuse: %d; NumCPU: %d; GoMaxProcs: %d; NumGoroutine: %d",
				mem.Sys, mem.HeapSys, mem.HeapInuse, mem.HeapObjects, mem.StackSys, mem.StackInuse, runtime.NumCPU(), runtime.GOMAXPROCS(-1), runtime.NumGoroutine())

			fmt := ""

			if len(details) > 0 {
				ok := false
				fmt, ok = details[0].(string)
				if !ok {
					fmt = ""
				}
			}

			if fmt != "" {
				p := []any{}
				if len(details) > 1 {
					p = details[1:]
				}
				misc.Logger("", "IN", fmt, p...)
			}

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

	return buf.String()
}

//----------------------------------------------------------------------------------------------------------------------------//
