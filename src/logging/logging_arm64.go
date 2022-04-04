package logging

import (
	"os"
)

func StderrToLogfile(logfile *os.File) {
	//syscall.Dup3(int(logfile.Fd()), 2, 0)
}
