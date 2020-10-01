package todo_test

import (
	"testing"

	"github.com/mrinjamul/gtodo/todo"
)

// TestGetVersion test GetVersion()
func TestGetVersion(t *testing.T) {
	v := todo.GetVersion()
	if len(v) == 0 {
		t.Errorf("Wants strings but got %v", v)
	}
}
