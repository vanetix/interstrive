package interstrive

/**
 * Test dependencies
 * - container/heap is implemented by Tasks
 */

import (
	"testing"
	"strconv"
	"container/heap"
)

/**
 * Basic Tests for the Tasks type implementing 
 * container/heap.
 */

// Test that `Add()` properly adds the element to 
// the heap.
func TestTaskAdd (t *testing.T) {
	tasks := make(Tasks, 0, 2)
	task := &Task{"Add Test", 1}

	heap.Push(&tasks, task)

	if len(tasks) != 1 {
		t.Error("Tasks should have a length of 1 after adding task")
	}
}

// Test that `Pop()` properly removed the task with the 
// lowest priority (most important)
func TestTaskPop (t *testing.T) {
	tasks := make(Tasks, 0, 2)

	for i := cap(tasks); i > 0; i-- {
		task := &Task{
			Name: "Task",
			Priority: i,
		}
		heap.Push(&tasks, task)
	}

	pop := heap.Pop(&tasks).(*Task)

	if pop.Priority != 2 {
		t.Error("The wrong priority task was popped. Got:", pop.Priority)
	}
}

// Test that remove functions correctly
func TestTaskRemove (t *testing.T) {
	tasks := make(Tasks, 0)

	for i := 1; i <= 10; i++ {
		task := &Task{
			Name: "Task " + strconv.Itoa(i),
			Priority: i,
		}
		heap.Push(&tasks, task)
	}

	if task := tasks.Remove(0); task == nil || task.Name != "Task 10" {
		t.Error("The wrong task was removed. Got: ", task)
	}
}
