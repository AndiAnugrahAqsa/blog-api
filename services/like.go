package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type LikeService struct {
	Repository repositories.LikeRepository
}

func NewLikeService() LikeService {
	return LikeService{
		Repository: &repositories.LikeRepositoryImpl{},
	}
}

func (cs *LikeService) GetAll() []models.Like {
	return cs.Repository.GetAll()
}

func (cs *LikeService) GetByBlogID(blog_id int) []models.Like {
	return cs.Repository.GetByBlogID(blog_id)
}

func (cs *LikeService) Create(likeRequest models.LikeRequest) models.Like {
	return cs.Repository.Create(likeRequest)
}

func (cs *LikeService) Delete(id int) bool {
	return cs.Repository.Delete(id)
}
