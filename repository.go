package foo

import (
	"context"
	"errors" // standard go errors package
	errs "github.com/pkg/errors" // imports another errors package and gives it the name "errs"
	// github.com/fabric8-services/fabric8-wit/errors 

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
}

type todoRepository struct{
	db *gorm.DB
}

/*
// in main package
func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()
	repo := foo.NewTodoRepsitory(db)
}
*/

func NewTodoRepository(db *gorm.DB) todoRepository {	
	todoRepository{
		db: db,
	}
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (p Todo) TableName() string {
	return "todos"
}

func (t *todoRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	if err := t.db.Create(todo).Error; err != nil {
		return nil, errs.Wrap(err, "failed to create todo item")
	}
	return todo, nil
}

func (t *todoRepository) Delete(ctx context.Context, ID uint) error {
	tempTodo := Todo{ID: ID}
	if err := t.db.Delete(&tempTodo).Error; err != nil {
		return errs.Wrap(err, "failed to delete item from database")
	}
	return nil
}

func (t *todoRepository) List(ctx context.Context) ([]Todo, error) {
	var todolist []Todo
	if err := t.db.Find(&todolist).Error; err != nil {
		return nil, errs.Wrap(err, "failed to fetch all todo items from database")
	}
	return todolist, nil
}

func (t *todoRepository) Show(ctx context.Context, ID uint) (*Todo, error) {
	var tempTodo Todo
	if err := t.db.First(&tempTodo, ID).Error; err != nil {
		return nil, errs.Wrap(err, "failed to fetch item from database")
	}
	return &tempTodo, nil
}

func (t *todoRepository) Update(ctx context.Context, newTodo Todo, ID uint) (*Todo, error) {
	var tempTodo Todo
	if err := t.db.First(&tempTodo, ID).Error; err != nil {
		// Todo - Raise this error diffently
		return nil, errors.New("failed to fetch item from database")
	}
	tempTodo.Title = newTodo.Title
	tempTodo.Description = newTodo.Description
	db := t.db.Save(&tempTodo)
	if db.Error != nil {
		return nil, errs.Wrap(db.Error, "failed to update todo")
	}
	return &tempTodo, nil
}
