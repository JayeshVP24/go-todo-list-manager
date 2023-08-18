package main

import (
	"fmt"

	"jayeshvp24.dev/todo-list-manager/todos"
)

type operation struct {
	operand   Operand
	operation string
}

func main() {
	todoList := todos.NewTodos()

	ops := []Operand{Add, Remove, Show}

	fmt.Println("--- Welcome to Todo List Manager ---")
	for {
		fmt.Println("Choose one option")
		for _, v := range ops {
			fmt.Printf("%d: %s \n", v, v.String())
		}
		var inp Operand
		fmt.Scanf("%d", &inp)
		switch inp {
		case Add:
			fmt.Println("Enter your new todo")
			var newtask todos.Todo
			fmt.Scanf("%s", &newtask.Task)
			todoList.Add(newtask)
		case Remove:
			fmt.Println("Enter the ID of todo to delete")
			var id int8
			fmt.Scanf("%d", &id)
			todoList.Remove(id)
		case Show:
			todoList.Show()
		}
	}
}
