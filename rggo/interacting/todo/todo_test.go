package todo_test

import (
	"os"
	"pragprog.com/rggo/interacting/todo"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Done {
		t.Errorf("Task should not be completed already.")
	}
	l[0].Done = true

	if !l[0].Done {
		t.Errorf("Unable to mark the task complete.")
	}
}

func TestDelete(t *testing.T) {
	tasks := []string{
		"task One",
		"task two",
		"task threee",
	}

	l := todo.List{}

	for _, v := range tasks {
		l.Add(v)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Length should be equal to 2 after deletion.")
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Elements do not match after deletion")
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "task1"
	l1.Add(taskName)

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Errorf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting data from file %s", err)
	}
	if l2[0].Task != taskName {
		t.Errorf("Expected task does not match from %q, %q", l2[0].Task, taskName)
	}
}
