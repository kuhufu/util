package daemon

import (
	"fmt"
	"log"
	"os"
)

func Log(args ...interface{}) {
	cp := make([]interface{}, 0, len(args)+1)
	cp = append(cp, fmt.Sprintf("[%v]", os.Getpid()))
	cp = append(cp, args...)
	log.Println(cp...)
}
