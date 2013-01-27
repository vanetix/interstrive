package main

import (
	"fmt"
	"flag"
	"container/heap"
	"github.com/vanetix/interstrive/interstrive"
)

func main() {
	task := interstrive.Task{"Test Task", 2}
	fmt.Println(task)
}
