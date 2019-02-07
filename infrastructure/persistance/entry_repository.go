package persistance

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
)

type entryRepository struct {
	db *gorm.DB
}

func NewEntryRepository(db *gorm.DB) repository.EntryRepository {
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
