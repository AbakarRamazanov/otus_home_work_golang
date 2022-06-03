package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	PrintQueue()
	Clear()
}

type lruCache struct {
	mutex    *sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache lruCache) PrintQueue() {
	l := cache.queue
	for i := l.Front(); i != nil; i = i.Next {
		if i.Next != nil {
		} else {
		}
		if i.Prev != nil {
		} else {
		}
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	_, exist := cache.items[key]
	if exist {
		item := cache.items[key].Value.(*cacheItem)
		item.value = value
		cache.queue.MoveToFront(cache.items[key])
	} else {
		if cache.capacity == cache.queue.Len() {
			item := cache.queue.Back().Value.(*cacheItem)
			delete(cache.items, item.key)
			cache.queue.Remove(cache.queue.Back())
		}
		cache.items[key] = cache.queue.PushFront(&cacheItem{key: key, value: value})
	}
	cache.mutex.Unlock()
	return exist
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	var result interface{}
	cache.mutex.RLock()
	_, exist := cache.items[key]
	if exist {
		cache.queue.MoveToFront(cache.items[key])
		item := cache.items[key].Value.(*cacheItem)
		result = item.value
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
