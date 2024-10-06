package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) addTodo(title string) {
	*todos = append(*todos, Todo{Title: title, Completed: false, CreatedAt: time.Now(), CompletedAt: nil})
}

func (todos *Todos) isIndexValid(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("index out of bounds")
	}
	return nil
}

func (todos *Todos) removeTodo(index int) error {
	temp := *todos

	if err := todos.isIndexValid(index); err != nil {
		return err
	}

	*todos = append(temp[:index], temp[index+1:]...)
	return nil
}

func (todos *Todos) toggleCompleted(index int) error {

	temp := *todos

	if err := todos.isIndexValid(index); err != nil {
		return err
	}

	var isCompleted bool = temp[index].Completed
	if !isCompleted {
		var currentCompletionTime time.Time = time.Now()
		temp[index].CompletedAt = &currentCompletionTime
	}

	temp[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) updateTitle(index int, updatedTitle string) error {
	if err := todos.isIndexValid(index); err != nil {
		return err
	}

	var temp Todos = *todos
	temp[index].Title = updatedTitle
	return nil
}

func (todos *Todos) printList() {
	var isCompletedEmoji string = ""
	for index, item := range *todos {
		if item.Completed {
			isCompletedEmoji = "✅"
		} else {
			isCompletedEmoji = "❌"
		}
		fmt.Println(index, " ", item.Title, " ", item.Completed, " ", isCompletedEmoji)
	}
}

func main() {
	for {
		
	}
}
