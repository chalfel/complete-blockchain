package core

type Storage interface {
	Put(*Block) error
}

type InMemoryStore struct{}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (im *InMemoryStore) Put(b *Block) error {
	return nil
}
