package repository

import (
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type EntryRepository interface {
	Store(entry *model.Entry) error
	FindAll(entries []*model.Entry) ([]*model.Entry, error)
	FindByID(id int) (*model.Entry, error)
}
