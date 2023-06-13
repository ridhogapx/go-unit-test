package service

import (
	"errors"
	"learn_unit_testing/entity"
	"learn_unit_testing/repository"
)

type CateogryService struct {
	Repository repository.CategoryRepository
}

func (service CateogryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")

	} else {
		return category, nil
	}
}
