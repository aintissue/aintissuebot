package bot

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

func loge(err error) {
	errMsg := getCallerInfo() + err.Error()
	log.Println(errMsg)
}

func logs(errMsg string) {
	errMsg = getCallerInfo() + errMsg
	log.Println(errMsg)
}

func getCallerInfo() (info string) {

	// pc, file, lineNo, ok := runtime.Caller(2)
	_, file, lineNo, ok := runtime.Caller(2)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	// funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // The Base function returns the last element of the path
	return fmt.Sprintf("%s:%d: ", fileName, lineNo)
}
