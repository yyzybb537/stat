package displayer

import (
	"os"
	"fmt"
	"strings"
)

var displaying int

func Printf(format string, a ... interface{}) {
	info := fmt.Sprintf(format, a...) + " "
	fmt.Printf(strings.Repeat("\b", displaying) + info)
	os.Stdout.Sync()
	displaying = len(info)
}
