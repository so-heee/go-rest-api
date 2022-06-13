package database

import "github.com/so-heee/go-rest-api/api/domain/model"

// TweetRepository.
type TweetRepository struct {
	SQLHandler
}

func (r *TweetRepository) FindById(id int) (*model.Tweet, error) {
	t := model.Tweet{}
	if err := r.First(&t, id).Error(); err != nil {
		return nil, err
	}
	return &t, nil
}
