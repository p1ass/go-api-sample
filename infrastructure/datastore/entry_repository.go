package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type entryRepository struct {
	db *gorm.DB
}

type EntryRepository interface {
	Store(entry *model.Entry) error
	FindAll(entries []*model.Entry) ([]*model.Entry, error)
}

func NewEntryRepository(db *gorm.DB) EntryRepository {
	return &entryRepository{db}
}

func (eR *entryRepository) Store(entry *model.Entry) error {
	return eR.db.Save(entry).Error
}

func (eR *entryRepository) FindAll(entries []*model.Entry) ([]*model.Entry, error) {
	err := eR.db.Find(&entries).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return entries, nil
}
