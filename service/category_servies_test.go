package service

import (
	"learn_unit_testing/entity"
	"learn_unit_testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CateogryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceSuccess(t *testing.T) {
	var category entity.Category = entity.Category{
		Id:   "1",
		Name: "Electro",
	}
}
