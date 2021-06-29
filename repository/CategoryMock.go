package repository

import (
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
)

type CategoryMock struct {
	Mock mock.Mock
}

func (repository CategoryMock) FindById(id string) *entity.Category {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil{
		return nil
	} else {
		category := argument.Get(0).(entity.Category)
		return &category
	}
}
