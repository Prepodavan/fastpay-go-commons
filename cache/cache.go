package cache

import "github.com/SolarLabRU/fastpay-go-commons/object"

type Cache struct {
	table Table
}

func (c *Cache) Add(request, response object.Hashable) {
	c.table[request.Hash()] = response
}

func (c *Cache) Get(request object.Hashable) (response object.Hashable) {
	return c.table[request.Hash()]
}

func (c *Cache) Has(request object.Hashable) (ok bool) {
	_, ok = c.table[request.Hash()]
	return
}

func (c *Cache) Find(request object.Hashable) (response object.Hashable, ok bool) {
	response, ok = c.table[request.Hash()]
	return
}

func NewCache() *Cache {
	return &Cache{table: make(Table)}
}
