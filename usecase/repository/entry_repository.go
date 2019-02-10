package repository

import (
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type EntryRepository interface {
	Store(entry *model.Entry) error
	Update(entry *model.Entry) error
	Delete(entry *model.Entry) error
	FindAll() ([]*model.Entry, error)
	FindByID(id int) (*model.Entry, error)
}
