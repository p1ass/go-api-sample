package persistance

import (
	"github.com/jinzhu/gorm"
	"github.com/naoki-kishi/go-api-sample/domain/model"
	"github.com/naoki-kishi/go-api-sample/usecase/repository"
)

type tagRepository struct {
	db *gorm.DB
}

// NewTagRepository creates a repository using MySQL for datasotre.
func NewTagRepository(db *gorm.DB) repository.TagRepository {
	return &tagRepository{db}
}

func (tR *tagRepository) FindByID(id int) (*model.Tag, error) {
	tag := model.Tag{ID: id}
	err := tR.db.First(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (tR *tagRepository) Store(tag *model.Tag) error {
	return tR.db.Save(tag).Error
}

func (tR *tagRepository) Update(tag *model.Tag) error {
	// Save will include all fields when perform the Updating SQL, even it is not changed
	return tR.db.Model(&model.Tag{ID: tag.ID}).Updates(tag).Error
}

func (tR *tagRepository) Delete(tag *model.Tag) error {
	return tR.db.Delete(tag).Error
}

func (tR *tagRepository) FindAll() ([]*model.Tag, error) {
	tags := []*model.Tag{}

	err := tR.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (tR *tagRepository) FindOrCreateAll(tags []*model.Tag) error {
	for _, t := range tags {
		err := tR.db.FirstOrCreate(t, &t).Error
		if err != nil {
			return err
		}
	}

	return nil
}
