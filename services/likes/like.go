package services

import (
	"mini-project/models"
	"mini-project/repositories"
)

type LikeService struct {
	Repository repositories.LikeRepository
}

func NewLikeService(repository repositories.LikeRepository) LikeService {
	return LikeService{
		Repository: repository,
	}
}

func (ls *LikeService) GetAll() []models.Like {
	return ls.Repository.GetAll()
}

func (ls *LikeService) GetByBlogID(blog_id int) []models.Like {
	return ls.Repository.GetByBlogID(blog_id)
}

func (ls *LikeService) Create(likeRequest models.LikeRequest) models.Like {
	return ls.Repository.Create(likeRequest)
}

func (ls *LikeService) Delete(id int) bool {
	return ls.Repository.Delete(id)
}
