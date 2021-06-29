package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_NotFound(t *testing.T)  {
	categoryRepository.Mock.On("FindById").Return(nil)
	category,err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, category,err)
}
