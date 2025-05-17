package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define the structure for a todo item
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const storageFile = "todos.json"

// Load todos from local storage (file)
func loadTodos() ([]Todo, error) {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		return []Todo{}, nil
	}
	bytes, err := os.ReadFile(storageFile)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	err = json.Unmarshal(bytes, &todos)
	return todos, err
}

// Save todos to local storage (file)
func saveTodos(todos []Todo) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, bytes, 0644)
}

// Add a new todo
func addTodo(title string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}
	todo := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
	return saveTodos(todos)
}

// List all todos
func listTodos() error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return nil
	}
	fmt.Println("To-Do List:")
	for _, todo := range todos {
		status := " "
		if todo.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Title)
	}
	return nil
}

// Complete a todo item by ID
func completeTodo(id int) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	found := false
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			found = true
		}
	}
	if !found {
		return fmt.Errorf("Todo with ID %d not found", id)
	}
	return saveTodos(todos)
}

// Remove a todo item by ID
func removeTodo(id int) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	newTodos := []Todo{}
	found := false
	for _, todo := range todos {
		if todo.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, todo)
	}
	if !found {
		return fmt.Errorf("Todo with ID %d not found", id)
	}
	return saveTodos(newTodos)
}

func printUsage() {
	fmt.Println("To-Do List Application")
	fmt.Println("Usage:")
	fmt.Println("  add <todo>          Add a new todo")
	fmt.Println("  list                List all todos")
	fmt.Println("  done <id>           Mark a todo as completed")
	fmt.Println("  remove <id>         Remove a todo by ID")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo description.")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		if err := addTodo(title); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Todo added.")
		}
	case "list":
		if err := listTodos(); err != nil {
			fmt.Println("Error:", err)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the todo ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			return
		}
		if err := completeTodo(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Todo marked as completed.")
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the todo ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			return
		}
		if err := removeTodo(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Todo removed.")
		}
	default:
		printUsage()
	}
}
