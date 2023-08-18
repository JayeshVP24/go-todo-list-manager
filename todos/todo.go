package todos

import (
	"errors"
	"fmt"
)

type Todo struct {
	Task string
}

type Todos struct {
	todoList []Todo
}

func NewTodos() Todos {
	return Todos{todoList: []Todo{}}
}

func (t *Todos) Add(todo Todo) {
	t.todoList = append(t.todoList, todo)
}

func (t *Todos) Remove(id int8) error {
	if id < 0 || id > int8(len(t.todoList)) {
		return errors.New("ID doesn't exist in your todos")
	}
	t.todoList = append(t.todoList[:id], t.todoList[id+1:]...)
	return nil
}

func (t *Todos) Show() {
	for i, v := range t.todoList {
		fmt.Printf("%d. %s\n", i, v.Task)
	}
	fmt.Println()
}
