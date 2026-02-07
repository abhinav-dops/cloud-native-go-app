package service

import "cloud-app-api/internal/repository"

type ItemService struct {
	repo *repository.MemoryRepo
}

func NewItemService(r *repository.MemoryRepo) *ItemService {
	return &ItemService{repo: r}
}

func (s *ItemService) CreateItem(name string) repository.Item {
	item := repository.Item{
		ID:   len(s.repo.GetAll()) + 1,
		Name: name,
	}
	s.repo.Save(item)
	return item
}

func (s *ItemService) ListItems() []repository.Item {
	return s.repo.GetAll()
}
