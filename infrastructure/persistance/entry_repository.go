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

func (eR *entryRepository) FindByID(id int) (*model.Entry, error) {
	entry := model.Entry{ID: id}
	err := eR.db.Preload("Tags").First(&entry).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return &entry, nil
}

func (eR *entryRepository) Store(entry *model.Entry) error {
	return eR.db.Save(entry).Error
}

func (eR *entryRepository) Update(entry *model.Entry) error {
	// Save will include all fields when perform the Updating SQL, even it is not changed
	return eR.db.Model(&model.Entry{ID: entry.ID}).Updates(entry).Error
}

func (eR *entryRepository) Delete(entry *model.Entry) error {
	return eR.db.Delete(entry).Error
}

func (eR *entryRepository) FindAll() ([]*model.Entry, error) {
	entries := []*model.Entry{}

	err := eR.db.Preload("Tags").Find(&entries).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return entries, nil
}
