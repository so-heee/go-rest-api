package service

import (
	"github.com/so-heee/go-rest-api/api/domain/model"
	"github.com/so-heee/go-rest-api/api/usecase/repository"
)

// TweetService.
type TweetService struct {
	TweetRepository repository.TweetRepository
}

func (s *TweetService) TweetById(id int) (tweet *model.Tweet, err error) {
	tweet, err = s.TweetRepository.FindById(id)
	return
}
