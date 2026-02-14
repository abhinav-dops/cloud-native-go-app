package service

type Repo interface {
	AddItem(name string) error
	GetItems() ([]string, error)
}

type ItemService struct {
	repo Repo
}

func NewItemService(r Repo) *ItemService {
	return &ItemService{repo: r}
}

// func (s *ItemService) CreateItem(name string) repository.Item {
// 	return s.repo.CreateItem(name)
// }

// func (s *ItemService) ListItems() []repository.Item {
// 	return s.repo.ListItems()
// }

func (s *ItemService) Add(name string) error {
	return s.repo.AddItem(name)
}

func (s *ItemService) List() ([]string, error) {
	return s.repo.GetItems()
}
