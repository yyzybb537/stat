package displayer

import (
	"time"
	"testing"
	"fmt"
)

func TestDisplayer(t *testing.T) {
	Printf("QPS: %d", 1)
	time.Sleep(time.Second)
	Printf("QPS: %d", 23)
	time.Sleep(time.Second)
	Printf("QPS: %d", 1024)
	time.Sleep(time.Second)
	fmt.Println("Insert")
	time.Sleep(time.Second)
	Printf("QPS: %d", 56)
	time.Sleep(time.Second * 3)
}
