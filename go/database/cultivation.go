package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Cultivation = &CultivationRepository{}

// CultivationRepository は repository.CultivationRepository を満たす構造体です。
type CultivationRepository struct {
	db *gorm.DB
}

// NewCultivationRepository はCultivationRepositoryのポインタを生成する関数です。
func NewCultivationRepository(db *gorm.DB) *CultivationRepository {
	return &CultivationRepository{db: db}
}

// FindByID は指定されたIDを持つ栽培している植物を取得します。
func (r *CultivationRepository) FindByID(id string) (*entity.Cultivation, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ栽培している植物を取得します。
func (r *CultivationRepository) FindAll() (*entity.Cultivations, error) {
	return nil, nil
}
// Store は栽培している植物を新規保存します。
func (r *CultivationRepository) Store(cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	return nil, nil
}

// UpdateByID は栽培している植物の情報を更新します。
func (r *CultivationRepository) UpdateByID(id string, cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ栽培している植物を削除します。
func (r *CultivationRepository) DeleteByID(id string) (*entity.Cultivation, error) {
	return nil, nil
}