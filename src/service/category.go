package service

import (
	"github.com/JrSchmidtt/golang-grpc/src/database"
	"github.com/JrSchmidtt/golang-grpc/src/pb"
)

// CategoryService is a struct that implements the CategoryServiceServer interface
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

// NewCategoryService is a function that returns a new CategoryService
func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}