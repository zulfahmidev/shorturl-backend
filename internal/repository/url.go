package repository

import (
	"github.com/jinzhu/copier"
	"github.com/zulfahmidev/shorturl/internal/dto"
	"github.com/zulfahmidev/shorturl/internal/model"
	"github.com/zulfahmidev/shorturl/pkg/database"
	"github.com/zulfahmidev/shorturl/pkg/token"
)

type URLRepository struct{}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (r *URLRepository) Create(url string) (*dto.URLData, error) {
	token, err := token.GenerateUniqueToken()
	if err != nil {
		return nil, err
	}

	db := database.GetDB()
	tx := db.Begin()

	urlModel := model.URL{
		URL:   url,
		Token: token,
	}

	err = tx.Create(&urlModel).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	var data dto.URLData
	copier.Copy(&data, urlModel)

	return &data, nil
}

func (r *URLRepository) FindByURL(url string) (*dto.URLData, error) {
	db := database.GetDB()

	var urlModel model.URL
	if err := db.Where("url = ?", url).Order("created_at DESC").Last(&urlModel).Error; err != nil {
		return nil, err
	}

	var data dto.URLData
	copier.Copy(&data, urlModel)

	return &data, nil
}
