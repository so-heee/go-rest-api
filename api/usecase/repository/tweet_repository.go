package repository

import "github.com/so-heee/go-rest-api/api/domain/model"

type TweetRepository interface {
	FindById(id int) (model.Tweet, error)
}
