package main

import (
	"fmt"

	"jayeshvp24.dev/todo-list-manager/todos"
)

func main() {
	todoList := todos.NewTodos()

	fmt.Println("--- Welcome to Todo List Manager ---")
	for {
		fmt.Println("Choose one option")
		for _, v := range todos.Ops {
			fmt.Printf("%d: %s \n", v, v.String())
		}
		var inp todos.Operand
		fmt.Scanf("%d", &inp)
		switch inp {
		case todos.Add:
			fmt.Println("Enter your new todo")
			var newtask todos.Todo
			fmt.Scanf("%s", &newtask.Task)
			todoList.Add(newtask)
		case todos.Remove:
			fmt.Println("Enter the ID of todo to delete")
			var id int8
			fmt.Scanf("%d", &id)
			todoList.Remove(id)
		case todos.Show:
			todoList.Show()
		}
	}
}
