package service

type IURLRepository interface {

}

type URLService struct {
	urlRepository *IURLRepository
}

func NewURLService(urlRepository IURLRepository) *URLService {
	return &URLService{
		urlRepository: &urlRepository,
	}
}