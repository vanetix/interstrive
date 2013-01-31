package main

import (
	"os"
	"fmt"
	"flag"
	"path"
	"container/heap"
	"github.com/vanetix/interstrive/interstrive"
)

var (
	// Flags
	list = flag.Bool("l", false, "list tasks, from highest to lowest priority")
	pop = flag.Bool("p", false, "highest priority task off the list")
	create = flag.String("c", "", "create a new task - add a priority with -n")
	priority = flag.Int("n", 0, "add a priority to the task being created")
	remove = flag.Bool("r", false, "remove all tasks")
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
	config := path.Join(os.Getenv("HOME"), ".interstrive.json")

	// TODO: Fix this error handling, basically ignoring a read error
	// Might check if the path exists first, then make a new tasks
    tasks.Load(config)

	if *list {
		fmt.Fprintln(os.Stdout, "\x1b[37mTasks:")
		for i := range tasks {
                        if i == 0 {
				fmt.Fprintf(os.Stdout, "\x1b[33;1m")
			} else {
				fmt.Fprintf(os.Stdout, "\x1b[0m\x1b[32m")
			}

			fmt.Fprintf(os.Stdout, "\t%d: %s\n", i + 1, tasks[i])
		}
	}

	if *pop {
		if tasks.Len() > 0 {
			task := heap.Pop(&tasks).(*interstrive.Task)
			fmt.Fprintf(os.Stdout, "\x1b[37;1mCompleted: \x1b[0m%s\n", task)
		} else {
			fmt.Fprintf(os.Stderr, "\x1b[31;1mYou have no tasks to pop.\x1b[0m\n\n")
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

	if *remove {
		tasks = make(interstrive.Tasks, 0)
	}

	// Save Tasks before exiting
	_, err := tasks.Save(config)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: \x1b[31m %s", err)
	}

	// Put a little padding on the bottom before the next term line
	fmt.Fprintf(os.Stdout, "\n")
	os.Exit(0)
}
