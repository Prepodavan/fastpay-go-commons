package cache

import "github.com/SolarLabRU/fastpay-go-commons/object"

type Table map[string]object.Hashable

func (t Table) Add(obj object.Hashable) {
	t[obj.Hash()] = obj
}

func (t Table) Has(obj object.Hashable) (ok bool) {
	_, ok = t[obj.Hash()]
	return
}

func (t Table) Find(h string) (obj object.Hashable, ok bool) {
	obj, ok = t[h]
	return
}

func (t Table) Get(h string) object.Hashable {
	return t[h]
}
