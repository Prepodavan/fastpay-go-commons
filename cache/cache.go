package cache

import "github.com/SolarLabRU/fastpay-go-commons/object"

type Cache map[string]object.Hashable

func (c Cache) Add(obj object.Hashable) {
	c[obj.Hash()] = obj
}

func (c Cache) Has(obj object.Hashable) (ok bool) {
	_, ok = c[obj.Hash()]
	return
}

func (c Cache) Find(h string) (obj object.Hashable, ok bool) {
	obj, ok = c[h]
	return
}

func (c Cache) Get(h string) object.Hashable {
	return c[h]
}
