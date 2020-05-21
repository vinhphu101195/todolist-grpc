package models

import "github.com/jinzhu/gorm"

type (
	// TodoModel describes a todoModel type
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
		Userid    int    `json:"userid"`
	}
	// TransformedTodo represents a formatted todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
		Userid    int    `json:"userid"`
	}
)
