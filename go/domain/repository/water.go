package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Water は水分量の永続化を担当するリポジトリです。
type Water interface {
	FindByID(id string) (*entity.Water, error)
	FindAll() (*entity.Waters, error)
	Store(water *entity.Water)  (*entity.Water, error)
	UpdateByID(id string, water *entity.Water) (*entity.Water, error)
	DeleteByID(id string) (*entity.Water, error)
}