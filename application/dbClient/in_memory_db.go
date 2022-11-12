package dbClient

import (
	"sync"
)

type InMemoryClient interface {
	CreateSync(key string, value string)
	GetSync(key string) string
}

type inMemoryClient struct {
	client    map[string]string
	writeLock sync.RWMutex
}

func NewInMemoryClient() InMemoryClient {
	return &inMemoryClient{
		client:    map[string]string{},
		writeLock: sync.RWMutex{},
	}
}

func (i *inMemoryClient) CreateSync(key string, value string) {
	i.writeLock.Lock()
	defer i.writeLock.Unlock()
	i.client[key] = value
}

func (i *inMemoryClient) GetSync(key string) string {
	i.writeLock.RLock()
	defer i.writeLock.RUnlock()
	return i.client[key]
}
