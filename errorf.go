package gcimporter

import (
	"fmt"
)

func errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
