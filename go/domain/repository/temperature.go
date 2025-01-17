package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Temperature は気温の永続化を担当するリポジトリです。
type Temperature interface {
	FindByID(id int) (*entity.Temperature, error)
	FindAll() (*entity.Temperatures, error)
	Store(temperature *entity.Temperature) (*entity.Temperature, error)
	UpdateByID(temperature *entity.Temperature) (*entity.Temperature, error)
	DeleteByID(id int) error
}
