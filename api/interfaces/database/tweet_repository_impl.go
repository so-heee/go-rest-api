package database

import "github.com/so-heee/go-rest-api/api/domain/model"

// TweetRepository.
type TweetRepository struct {
	SQLHandler
}

func (repo *TweetRepository) FindById(id int) (tweet *model.Tweet, err error) {
	if err = repo.First(&tweet, id).Error(); err != nil {
		return
	}
	return
}
