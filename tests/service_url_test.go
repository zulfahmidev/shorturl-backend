package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zulfahmidev/shorturl/internal/dto"
	"github.com/zulfahmidev/shorturl/internal/service"
	"github.com/zulfahmidev/shorturl/tests/mocks"
	"gorm.io/gorm"
)

func TestURLService_GenerateShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mocks.NewMockIURLRepository(ctrl)
	mockService := 	service.NewURLService(mockRepo)
	url := "http://zulfahmidev.com"

	expectedData := &dto.URLData{
		Token: "abcdefghij",
		URL: url,
		VisitorCount: 0,
	}


	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindByURL(url).Return(nil, gorm.ErrRecordNotFound)
		mockRepo.EXPECT().Create(url).Return(expectedData, nil)

		result, err := mockService.GenerateShortURL(url, true)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedData.URL, result.URL)
		assert.Equal(t, expectedData.VisitorCount, result.VisitorCount)
		assert.NotEmpty(t, result.Token)
	})

	t.Run("idempotency set to false", func(t *testing.T) {
		mockRepo.EXPECT().Create(url).Return(expectedData, nil)

		result, err := mockService.GenerateShortURL(url, false)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedData.URL, result.URL)
		assert.Equal(t, expectedData.VisitorCount, result.VisitorCount)
		assert.NotEmpty(t, result.Token)
	})

	t.Run("create url error", func(t *testing.T) {
		mockRepo.EXPECT().FindByURL(url).Return(nil, gorm.ErrRecordNotFound)
		mockRepo.EXPECT().Create(url).Return(nil, errors.New("db error"))

		data, err := mockService.GenerateShortURL(url, true)
		assert.Error(t, err)
		assert.Nil(t, data)
	})


	t.Run("create url error on idempotency set to false", func(t *testing.T) {
		mockRepo.EXPECT().Create(url).Return(nil, errors.New("db error"))

		data, err := mockService.GenerateShortURL(url, false)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestURLService_GetRedirectURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIURLRepository(ctrl)
	mockService := service.NewURLService(mockRepo)

	token := "abcdefghij"
	expectedData := &dto.URLData{
		URL:   "https://zulfahmidev.com",
		Token: token,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindByToken(token).Return(expectedData, nil)
		mockRepo.EXPECT().IncrementVisitCount(token).Return(nil)

		data, err := mockService.GetRedirectURL(token)
		assert.NoError(t, err)
		assert.Equal(t, expectedData, data)
	})

	t.Run("find by token error", func(t *testing.T) {
		mockRepo.EXPECT().FindByToken(token).Return(nil, gorm.ErrRecordNotFound)

		data, err := mockService.GetRedirectURL(token)
		assert.Error(t, err)
		assert.Nil(t, data)
	})

	t.Run("increment visit count error", func(t *testing.T) {
		mockRepo.EXPECT().FindByToken(token).Return(expectedData, nil)
		mockRepo.EXPECT().IncrementVisitCount(token).Return(errors.New("db error"))

		data, err := mockService.GetRedirectURL(token)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestURLService_GetStatsURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIURLRepository(ctrl)
	mockService := service.NewURLService(mockRepo)

	token := "abcdefghij"
	expectedData := &dto.URLData{
		URL:   "https://zulfahmidev.com",
		Token: token,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindByToken(token).Return(expectedData, nil)

		data, err := mockService.GetStatsURL(token)
		assert.NoError(t, err)
		assert.Equal(t, expectedData, data)
	})

	t.Run("find by token error", func(t *testing.T) {
		mockRepo.EXPECT().FindByToken(token).Return(nil, gorm.ErrRecordNotFound)

		data, err := mockService.GetStatsURL(token)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}