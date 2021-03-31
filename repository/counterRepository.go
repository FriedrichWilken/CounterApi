package repository

import "counter/models"

type CounterRepository struct {
	Counter int
}

func New(initialCounter int) *CounterRepository {
	return &CounterRepository{
		Counter: initialCounter,
	}
}

func (repo *CounterRepository) IncreaseBy(increaseBy int) {
	repo.Counter += increaseBy
}

func (repo *CounterRepository) DecreaseBy(decreaseBy int) {
	repo.Counter -= decreaseBy
}

func (repo *CounterRepository) ToDto() *models.CounterDto {
	return &models.CounterDto{
		Value: repo.Counter,
	}
}
