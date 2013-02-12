package interstrive

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/**
 * Task type for storing associated
 * task data
 */

type Task struct {
	Name string
	Priority int
}

/**
 * Task string method for easy printing to stdout
 */

func (task Task) String() string {
	return fmt.Sprintf("%s", task.Name)
}

/**
 * Tasks, implements `container/heap`
 */

type Tasks []*Task

/**
 * `Len()` to implement `container/heap`
 *
 * @return {int}
 */

func (tasks Tasks) Len() int {
	return len(tasks)
}

/**
 * `Less()` to implement `container/heap`
 *  Does the `Task` at index i have a greater `Priority` than index j?
 *
 * @param {int} i
 * @param {int} j
 * @return {bool}
 */

func (tasks Tasks) Less(i, j int) bool {
	return tasks[i].Priority > tasks[j].Priority
}

/**
 * `Swap()` to implement `container/heap`
 *  Swaps `Task` at i with `Task` at j
 *
 * @param {int} i
 * @param {int} j
 */
func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

/**
 * `Push()` to implement `container/heap`
 * Push `i.(*Task)` onto tasks
 *
 * @param {interface} i
 */

func (tasks *Tasks) Push(i interface{}) {
	*tasks = append(*tasks, i.(*Task))
}

/**
 * `Pop()` to implement `container/heap`
 *  Pops the highest priority task of the heap
 *
 * @return {interface}
 */

func (tasks *Tasks) Pop() interface{} {
	len := len(*tasks)
	task := (*tasks)[len - 1]
	*tasks = (*tasks)[0 : len - 1]

	return task
}

/**
 * Remove task at index `i` from the queue
 *
 * TODO: Might have to heapify after calling `Remove()`
 * TODO: Fix the silly loop over tail to append them
 *
 * @param {int} i
 */

func (tasks *Tasks) Remove(i int) Task {
	// Check for bounds
	if i > 0 && i < len(*tasks) {
		task := tasks[i]
		next := (*tasks)[:i]
		next = append(next, (*tasks)[i + 1:]...)
		*tasks = next
		return task
	}

	return nil
}

/**
 * Encode the `tasks` as json and save
 * to `~/.interstrive.json`.
 *
 * @param {string} path
 * @return {bool}
 * @return {error}
 */

func (tasks Tasks) Save(path string) (bool, error) {
	jsonStr, jsonErr := json.Marshal(tasks)

	if jsonErr != nil {
		return false, jsonErr
	}

	writeErr := ioutil.WriteFile(path, jsonStr, 0644)

	if writeErr != nil {
		return false, writeErr
	}

	return true, nil
}

/**
 * Load the json encoded tasks from
 * `~/.interstrive.json`
 *
 * @param {string} path
 * @return {bool}
 * @return {error}
 */

func (tasks *Tasks) Load(path string) (bool, error) {
	jsonStr, readErr := ioutil.ReadFile(path)

	if readErr != nil {
		return false, readErr
	}

	jsonErr := json.Unmarshal(jsonStr, tasks)

	if jsonErr != nil {
		return false, jsonErr
	}

	return true, nil
}
