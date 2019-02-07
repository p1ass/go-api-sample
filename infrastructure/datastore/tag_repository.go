package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type tagRepository struct {
	db *gorm.DB
}

type TagRepository interface {
	Store(tag *model.Tag) error
	FindAll(entries []*model.Tag) ([]*model.Tag, error)
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db}
}

func (eR *tagRepository) Store(tag *model.Tag) error {
	return eR.db.Save(tag).Error
}

func (eR *tagRepository) FindAll(entries []*model.Tag) ([]*model.Tag, error) {
	err := eR.db.Find(&entries).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return entries, nil
}
