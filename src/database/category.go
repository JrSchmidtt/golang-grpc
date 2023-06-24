package database

import (
	"database/sql"
	// "github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

// NewCategory is a function that returns a new Category
func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

// Create is a function that creates a new category in the database
// this function return a mock category for now
func (c *Category) Create(name string, description string) (Category, error) {
	// return a mock category
	return Category{
		ID:          "1",
		Name:        "mock category",
		Description: "mock description",
	}, nil
}


// FindAll is a function that finds all categories
// this function return a mock category for now
func (c *Category) FindAll() ([]Category, error) {
	// return a mock categoryList
	return []Category{
		{
			ID:          "1",
			Name:        "mock category",
			Description: "mock description",
		},
		{
			ID:          "2",
			Name:        "mock category 2",
			Description: "mock description 2",
		},
	}, nil
}

// FindByCourseID is a function that finds a category by course id
// this function return a mock category for now
func (c *Category) FindByCourseID(courseID string) (Category, error) {
	return Category{
		ID:          courseID,
		Name:        "mock category",
		Description: "mock description",
	}, nil
}

// Find is a function that finds a category by id
// this function return a mock category for now
func (c *Category) Find(id string) (Category, error) {
	return Category{
		ID:          id,
		Name:        "mock category",
		Description: "mock description",
	}, nil
}