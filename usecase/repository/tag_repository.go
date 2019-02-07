package repository

import (
	"github.com/naoki-kishi/go-api-sample/domain/model"
)

type TagRepository interface {
	Store(tag *model.Tag) error
	FindAll(entries []*model.Tag) ([]*model.Tag, error)
}
