package todo

import (
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (repository *TodoRepository) FindAll() []Todo {
	var todos []Todo
	repository.db.Order("ID desc").Find(&todos)
	return todos
}

func (repository *TodoRepository) Create(todo Todo) (Todo, error) {
	err := repository.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: database,
	}
}
