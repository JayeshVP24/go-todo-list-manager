# Todo List Manager using Go Lang

## Factories Pattern

### Todo Factory
Has all the logic of Todo and all manipulation of Todo Slice
- Todo Struct
    - Task string
- Todos Struct
    - todoList []Todo
    - Add(t Todo)
    - Remove(idx int8)
    - Show()
- NewTodos() Todos

### Operand Factory
Is an ENUM for operations allowd on Todo Factory
- Add, Remove, Show 
- String() string

## Todo
Currently a struct that includes just a string called task. Kept it a struct so we can add more properties when needed in future.

## Todos
Extends Todo as an slice. Will be the array that contains all todos.
### Add
Adds a Todo to the slice. Requires a Todo as argument
### Remove
Removed a Todo from the slice. Requires inx as an argument
### Show
Prints all Todos in a formatted manner.

## Operand
Just an ENUM that contains all allowed operations on Todo Factory.
All values are integers automatically incremented by IOTA starting
index is 1.
### String
Returns string version of current operand.
