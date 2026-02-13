package service

import "cloud-app-api/internal/repository"

type Repo interface {
	CreateItem(name string) repository.Item
	ListItems() []repository.Item
}

type ItemService struct {
	repo Repo
}

func NewItemService(r Repo) *ItemService {
	return &ItemService{repo: r}
}

func (s *ItemService) CreateItem(name string) repository.Item {
	return s.repo.CreateItem(name)
}

func (s *ItemService) ListItems() []repository.Item {
	return s.repo.ListItems()
}
