package repository

import "learn_unit_testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
