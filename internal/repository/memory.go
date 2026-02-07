package repository

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MemoryRepo struct {
	items []Item
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		items: []Item{},
	}
}

func (m *MemoryRepo) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m *MemoryRepo) Save(item Item) {
	m.items = append(m.items, item)
}

func (m *MemoryRepo) GetAll() []Item {
	return m.items
}
