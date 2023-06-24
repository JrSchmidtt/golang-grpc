package service

import (
	"context"

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

// CreateCategory is a function handler that creates a new category
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

// ListCategories is a function handler that lists all categories
func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category

	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoryList}, nil
}