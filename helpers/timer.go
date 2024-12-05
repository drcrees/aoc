package helpers

import (
	"fmt"
	"time"
)

func Measure(exec func()) {
	now := time.Now()
	defer func() {
		fmt.Printf("in %s\n", time.Now().Sub(now))
	}()

	exec()
}
