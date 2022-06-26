package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    &sync.Mutex{},
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newItem := cacheItem{
		key,
		value,
	}
	_, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(l.items[key])
		l.queue.Front().Value = newItem
	} else {
		if l.queue.Len() == l.capacity {
			delete(l.items, l.queue.Back().Value.(cacheItem).key)
			l.queue.Remove(l.queue.Back())
		}
		l.items[key] = l.queue.PushFront(newItem)
	}
	return ok
}

// Получить элемент из кеша по ключу.
func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	el, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(l.items[key])
		return el.Value.(cacheItem).value, true
	}
	return nil, false
}

// Очистить кэш.
func (l *lruCache) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}
