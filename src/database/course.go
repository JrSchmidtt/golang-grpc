package database

import (
	"database/sql"
	"github.com/google/uuid"
)

// Course is a struct that represents the course table in the database
type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

// NewCourse is a function that returns a new Course
func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

// Create is a function that creates a new course
// and returns a course struct
func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

// FindAll is a function that finds all courses
// and returns a slice of course struct
func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}

// FindByCategoryID is a function that finds all courses by category id
// and returns a slice of course struct
func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}

// Find is a function that finds a course by id
// and returns a course struct
func (c *Course) Find(id string) (Course, error) {
	var name, description, categoryID string
	err := c.db.QueryRow("SELECT name, description, category_id FROM courses WHERE id = $1", id).
		Scan(&name, &description, &categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}