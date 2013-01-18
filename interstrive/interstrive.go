package interstrive

import "fmt"
import "encoding/json"
import "container/heap"

/**
 * Task type for storing associated
 * task data
 */

type Task struct {
	name string
	priority int
}

/**
 * Task string method for easy printing to stdout
 */
func (task *Task) String() string {
	return fmt.Sprintf("Task: %s, Priority: %d", task.title, task.priority)
}

/**
 * Task heap
 */

type Tasks []Task

func (tasks Tasks) Len() int {
	return len(tasks)
}

func (tasks Tasks) Less(i, j int) bool {
	return tasks[i].priority < tasks[j].priority
}

func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

func (tasks *Tasks) Push(t Task) {
	len := len(tasks)
	tasks = tasks[0 : len + 1]
	tasks[len] = t

	//next := *tasks
	//len := len(next)
	//next = next[0 : len + 1]
	//next[len] = t
	//tasks = next
}

func (tasks *Tasks) Pop() Task {
	len := len(tasks)
	task := tasks[len - 1]
	tasks = tasks[0 : len - 1]

	return task

	//next := *tasks
	//len := len(next)
	//task := next[len - 1]
	//*tasks = next[0 : len - 1]

	//return task
}

/**
 * Encode the `tasks` as json and save
 * to `~/.interstrive.json`.
 */

func (tasks Tasks) Save() (status bool, err error) {
	err, jsonStr = json.Marshal(tasks)
	return
}

/**
 * Load the json encoded tasks from
 * `~/.interstrive.json`
 */

func (tasks Tasks) Load() {

}
