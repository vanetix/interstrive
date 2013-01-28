package main

import (
	"os"
	"fmt"
	"flag"
	"container/heap"
	"github.com/vanetix/interstrive/interstrive"
)

var (
	// Flags
	list = flag.Bool("l", false, "list tasks, from highest to lowest priority")
	pop = flag.Bool("p", false, "highest priority task off the list")
	create = flag.String("a", "", "create a new task - add a priority with -n")
	priority = flag.Int("n", 0, "add a priority to the task being created")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: strive [flag] [value]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

    // If no arguments passed print usage and exit
	if flag.NFlag() == 0 {
		usage()
	}

	// Initialize and load tasks from ~/.interstrive.json
    tasks := make(interstrive.Tasks, 0)

	// TODO: Fix this error handling, basically ignoring a read error
	// Might check if the path exists first, then make a new tasks
    tasks.Load("~/.interstrive.json")

	if *list {
		for i := range tasks {
			fmt.Println(tasks[i])
		}
	}

	if *pop {
		if tasks.Len() > 0 {
			task := heap.Pop(&tasks).(*interstrive.Task)
			fmt.Println(task)
		} else {
			fmt.Fprintf(os.Stderr, "You have no tasks to pop")
			usage()
		}
	}

	if *create != "" {
		task := &interstrive.Task{
			Name: *create,
			Priority: *priority,
		}

		heap.Push(&tasks, task)
	}


	// Save Tasks before exiting
	tasks.Save("~/.interstrive.json")
	os.Exit(0)
}
