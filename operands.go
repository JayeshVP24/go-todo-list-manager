package main

type Operand int8

const (
	Add Operand = iota + 1
	Remove
	Show
)

func (o Operand) String() string {
	switch o {
	case Add:
		return "Add a todo"
	case Remove:
		return "Remove a todo"
	case Show:
		return "Show all todos"
	}
	return "unknown"
}
