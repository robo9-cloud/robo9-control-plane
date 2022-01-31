package api

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"strings"
)

// TODO: Integrate api layer with gorm database. Currently storing in memory.
type Task struct {
	ID        string `json:"id"`
	Task_id   string `json:"task_id"`
	Task_type string `json:"task_type"`
}

var (
	tasks []*Task
)

func GetTasks() []*Task {
	return tasks
}

func AddTask(t Task) (Task, error) {

	t.ID = snowflakeUUID()
	tasks = append(tasks, &t)
	return t, nil
}

func GetTaskByID(id string) (Task, error) {
	for _, u := range tasks {
		if u.ID == id {
			return *u, nil
		}
	}
	return Task{}, fmt.Errorf("Task with ID '%v' not found", id)
}

func UpdateTask(t Task) (Task, error) {
	for i, candidate := range tasks {
		if candidate.ID == t.ID {
			tasks[i] = &t
			return t, nil
		}
	}
	return Task{}, fmt.Errorf("Task with ID '%v' not found", t.ID)
}

func RemoveTaskById(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with ID '%v' not found", id)
}

func longUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func snowflakeUUID() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	id := node.Generate()
	return id.String()
}
