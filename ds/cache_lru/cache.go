package cache_lru

import "container/list"

type Item struct {
	Payload interface{}
	Ptr     *list.Element
}

type CacheLRU struct {
	Items    map[string]*Item
	Queue    *list.List
	Capacity int
}

func New(capacity int) *CacheLRU {
	return &CacheLRU{
		Items:    make(map[string]*Item),
		Queue:    list.New(),
		Capacity: capacity,
	}
}

func (c *CacheLRU) Get(key string) interface{} {
	if item, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(item.Ptr)

		return item.Payload
	}

	return nil
}

func (c *CacheLRU) Put(key string, value interface{}) {
	if item, ok := c.Items[key]; ok {
		item.Payload = value
		c.Items[key] = item
		c.Queue.MoveToFront(item.Ptr)

		return
	}

	if len(c.Items) == c.Capacity {
		back := c.Queue.Back()
		c.Queue.Remove(back)
		delete(c.Items, back.Value.(string))
	}

	c.Items[key] = &Item{
		Payload: value,
		Ptr:     c.Queue.PushFront(key),
	}
}
