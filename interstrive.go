package main

import (
	"os"
	"fmt"
	"path"
	"strconv"
	"container/heap"
	"github.com/vanetix/commander.go"
	"github.com/vanetix/interstrive/interstrive"
)

/**
 * Global variables
 */

var (
	config string
	task *interstrive.Task
	tasks interstrive.Tasks
	program *commander.Commander
)

/**
 * Initialize globals
 */

func init() {
	// Initialize a new commander instance
	program = commander.Init("interstrive", "1.0.0")

	// Initialize and load tasks from ~/.interstrive.json
	tasks = make(interstrive.Tasks, 0)

	// Save config path
	config = path.Join(os.Getenv("HOME"), ".interstrive.json")

	// TODO: Fix this error handling, basically ignoring a read error
	tasks.Load(config)
}

/**
 * List tasks
 */

func listTasks(args ...string) {
	if tasks.Len() == 0 {
		fmt.Fprintf(os.Stdout, "\x1b[31;1m  You have no tasks.\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "\x1b[37m  Tasks:\n")

		for i := range tasks {
			if i == 0 {
				fmt.Fprintf(os.Stdout, "\x1b[33;1m")
			} else {
				fmt.Fprintf(os.Stdout, "\x1b[0m\x1b[32m")
			}

			fmt.Fprintf(os.Stdout, "    %d: %s\n", i + 1, tasks[i])
		}

		fmt.Fprintf(os.Stdout, "\n")
	}
}

/**
 * Pop highest priority task off the list
 */

func popTask(args ...string) {
	if tasks.Len() > 0 {
		task := heap.Pop(&tasks).(*interstrive.Task)
		fmt.Fprintf(os.Stdout, "\x1b[37;1m  Completed: \x1b[0m%s\n\n", task)
	} else {
		fmt.Fprintf(os.Stderr, "\x1b[31;1m  You have no tasks to pop.\x1b[0m\n\n")
		program.Usage()
	}
}

/**
 * Create a task
 */

func createTask(args ...string) {
	if len(args) == 0 {
		program.Usage()
	} else {
		task = &interstrive.Task{
			Name: args[0],
			Priority: 0, // Default priority == 0
		}
	}
}

/**
 * Set the priority if a task is being created
 */

func setPriority(args ...string) {
	if len(args) == 0 || task == nil {
		program.Usage()
	} else {
		n, err := strconv.Atoi(args[0])

		if err != nil {
			program.Usage()
		} else {
			task.Priority = n
		}
	}
}

/**
 * Remove task `args[0]` if present, else remove all tasks
 */

func removeTask(args ...string) {
	if len(args) == 0 {
		tasks = make(interstrive.Tasks, 0)
	}
}

func main() {
	list := &commander.Option{
		Name: "list",
		Tiny: "-l",
		Verbose: "--list",
		Description: "list tasks from highest to lowest priority",
		Required: false,
		Callback: listTasks,
	}

	pop := &commander.Option{
		Name: "pop",
		Tiny: "-p",
		Verbose: "--pop",
		Description: "pop the highest priority task off the list",
		Required: false,
		Callback: popTask,
	}

	create := &commander.Option{
		Name: "create",
		Tiny: "-c",
		Verbose: "--create",
		Description: "create a new task - add a priority with -n",
		Required: false,
		Callback: createTask,
	}

	priority := &commander.Option{
		Name: "priority",
		Tiny: "-n",
		Verbose: "--priority",
		Description: "create a new task with priority",
		Required: false,
		Callback: setPriority,
	}

	remove := &commander.Option{
		Name: "remove",
		Tiny: "-r",
		Verbose: "--remove",
		Description: "remove all tasks",
		Required: false,
		Callback: removeTask,
	}

	program.Add(list, pop, create, priority, remove)
	program.Parse()

	if task != nil {
		heap.Push(&tasks, task)
		fmt.Fprintf(os.Stdout, "\x1b[371m  Added: \x1b[0m\x1b[32m%s\n\n", task.Name)
	}

	// Save Tasks before exiting
	_, err := tasks.Save(config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "  Error: \x1b[31m %s\n", err)
	}

	os.Exit(0)
}
