package cache

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[string]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[string]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key string, value interface{}) bool {
	newItem := cacheItem{key, value}

	if item, wasSet := c.items[key]; wasSet {
		item.Value = newItem
		c.queue.MoveToFront(item)

		return wasSet
	}

	c.items[key] = c.queue.PushFront(newItem)
	if c.queue.Len() > c.capacity {
		delKey := c.queue.Back().Value.(cacheItem).key
		delete(c.items, delKey)
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(key string) (value interface{}, wasSet bool) {
	item, wasSet := c.items[key]
	if wasSet {
		value = item.Value.(cacheItem).value
		c.queue.MoveToFront(item)
	}
	return
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = map[string]*ListItem{}
}
