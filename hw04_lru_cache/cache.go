package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache    // Remove me after realization.
	mutex    *sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
// 	key   Key
// 	value interface{}
// }

func (cache lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	_, exist := cache.items[key]
	if exist {
		cache.items[key].Value = value
		cache.queue.MoveToFront(cache.items[key])
	} else {
		cache.items[key] = cache.queue.PushFront(value)
	}
	cache.mutex.Unlock()
	return exist
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	var result interface{}
	cache.mutex.RLock()
	value, exist := cache.items[key]
	if exist {
		cache.queue.MoveToFront(cache.items[key])
		result = value.Value
	}
	cache.mutex.RUnlock()
	return result, exist
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.mutex.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    &sync.RWMutex{},
	}
}
