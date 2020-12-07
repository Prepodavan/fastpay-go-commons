package sortedmap

import "sort"

type SortedStringsMap struct {
	m  map[string]interface{}
	sl sort.StringSlice
}

type UpdatePredicate func(interface{}) bool

func NewSortedStringsMap() *SortedStringsMap {
	return &SortedStringsMap{
		m:  make(map[string]interface{}),
		sl: make(sort.StringSlice, 0),
	}
}

func (ssm *SortedStringsMap) Keys() []string {
	return ssm.sl
}

func (ssm *SortedStringsMap) Get(key string) (val interface{}) {
	val, _ = ssm.m[key]
	return
}

func (ssm *SortedStringsMap) Has(key string) (ok bool) {
	_, ok = ssm.m[key]
	return
}

func (ssm *SortedStringsMap) Find(key string) (val interface{}, ok bool) {
	val, ok = ssm.m[key]
	return
}

func (ssm *SortedStringsMap) Set(key string, val interface{}) (out *SortedStringsMap) {
	out = ssm
	if !ssm.Has(key) {
		ssm.insertKey(key)
	}
	ssm.m[key] = val
	return
}

func (ssm *SortedStringsMap) insertKey(key string) (out *SortedStringsMap) {
	out = ssm
	i := ssm.sl.Search(key)
	ssm.sl = append(ssm.sl, "")
	copy(ssm.sl[i+1:], ssm.sl[i:])
	ssm.sl[i] = key
	return
}

func (ssm *SortedStringsMap) InsertOrUpdateIf(
	key string,
	val interface{},
	predicate UpdatePredicate,
) (out *SortedStringsMap) {
	out = ssm
	old, ok := ssm.Find(key)
	switch {
	case !ok:
		ssm.m[key] = val
		ssm.insertKey(key)
	case ok && predicate(old):
		ssm.m[key] = val
	}
	return
}

func (ssm *SortedStringsMap) Range(fn func(key string, val interface{})) (out *SortedStringsMap) {
	out = ssm
	for _, key := range ssm.sl {
		fn(key, ssm.m[key])
	}
	return
}
