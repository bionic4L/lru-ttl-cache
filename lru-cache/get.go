package lru_cache

import (
	"container/list"
	"time"
)

/*
Get returns Item.Value and move element to front of queue.
Also, checks if Item expired – if true delete it and return nil.
*/
func (c *LRU) Get(key string) interface{} {
	value, exists := c.Items.Load(key)
	if !exists {
		return nil
	}

	element := value.(*list.Element)
	item := element.Value.(*Item)

	if time.Now().After(item.ExpiresAt) {
		c.Queue.Remove(element)
		c.Items.Delete(key)
		return nil
	}

	c.Queue.MoveToFront(element)
	return element.Value.(*Item).Value
}
