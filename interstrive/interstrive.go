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
	return fmt.Sprintf("Name: %s, Priority: %d", task.Name, task.Priority)
}

/**
 * Tasks, implements `container/heap`
 */

type Tasks []*Task

func (tasks Tasks) Len() int {
	return len(tasks)
}

func (tasks Tasks) Less(i, j int) bool {
	return tasks[i].Priority > tasks[j].Priority
}

func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

func (tasks *Tasks) Push(i interface{}) {
	len := len(*tasks)
	*tasks = (*tasks)[0 : len + 1]
	(*tasks)[len] = i.(*Task)
}

func (tasks *Tasks) Pop() interface{} {
	len := len(*tasks)
	task := (*tasks)[len - 1]
	*tasks = (*tasks)[0 : len - 1]

	return task
}

/**
 * Encode the `tasks` as json and save
 * to `~/.interstrive.json`.
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
