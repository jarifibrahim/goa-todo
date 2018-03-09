package main

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
}

type todoRepository struct{}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (p Todo) TableName() string {
	return "todos"
}

func (t *todoRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	if err := db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoRepository) Delete(ctx context.Context, ID uint) error {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return errors.New("Failed to connect to database")
	}
	defer db.Close()

	var tempTodo Todo
	if err := db.First(&tempTodo, ID).Error; err != nil {
		return errors.New("Failed to fetch item from database")
	}
	if err := db.Delete(&tempTodo).Error; err != nil {
		return errors.New("Failed to delete item from database")
	}
	return nil
}

func (t *todoRepository) List(ctx context.Context) ([]Todo, error) {
	return nil, nil
}

func (t *todoRepository) Show(ctx context.Context, ID uint) (*Todo, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}
	defer db.Close()

	var tempTodo Todo
	if err := db.First(&tempTodo, ID).Error; err != nil {
		return nil, errors.New("Failed to fetch item from database")
	}
	return &tempTodo, nil
}

func (t *todoRepository) Update(ctx context.Context, newTodo *Todo, ID uint) (*Todo, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}
	defer db.Close()

	var tempTodo Todo
	if err := db.First(&tempTodo, ID).Error; err != nil {
		// Todo - Raise this error diffently
		return nil, errors.New("Failed to fetch item from database")
	}
	// Update Title
	if tempTodo.Title != newTodo.Title {
		newDB := db.Model(&tempTodo).Update("Title", newTodo.Title)
		if newDB.Error != nil {
			return &Todo{}, newDB.Error
		}
	}
	// Update Description
	if tempTodo.Description != newTodo.Description {
		newDB := db.Model(&tempTodo).Update("Description", newTodo.Description)
		if newDB.Error != nil {
			return &Todo{}, newDB.Error
		}
	}
	return &tempTodo, nil
}
