package repository

import (
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type TagRepository interface {
	Store(tag *model.Tag) error
	FindAll() ([]*model.Tag, error)
	FindByID(id int) (*model.Tag, error)
}
