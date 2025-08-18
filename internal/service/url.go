package service

import (
	"github.com/zulfahmidev/shorturl/internal/dto"
	"gorm.io/gorm"
)

type IURLRepository interface {
	Create(url string) (*dto.URLData, error)
	FindByURL(url string) (*dto.URLData, error)
	FindAll() (*[]dto.URLData, error)
	FindByToken(token string) (*dto.URLData, error)
	IncrementVisitCount(token string) error
}

type URLService struct {
	urlRepository IURLRepository
}

func NewURLService(urlRepository IURLRepository) *URLService {
	return &URLService{
		urlRepository: urlRepository,
	}
}

func (s *URLService) GenerateShortURL(url string, idempotent bool) (*dto.URLData, error) {
	var (
		data *dto.URLData
		err  error
	)

	if idempotent {
		data, err = s.urlRepository.FindByURL(url)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}

		if data != nil {
			return data, nil
		}
	}

	data, err = s.urlRepository.Create(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *URLService) GetRedirectURL(token string) (*dto.URLData, error) {
	data, err := s.urlRepository.FindByToken(token)
	if err != nil {
		return nil, err
	}

	err = s.urlRepository.IncrementVisitCount(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *URLService) GetStatsURL(token string) (*dto.URLData, error) {
	data, err := s.urlRepository.FindByToken(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *URLService) GetListURL() (*[]dto.URLData, error) {
	data, err := s.urlRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}