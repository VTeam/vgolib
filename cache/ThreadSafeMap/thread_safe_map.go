package ThreadSafeMap

import "sync"

type IndexFunc func(obj interface{}) ([]string, error)

type ThreadSafeMap struct {
	sync.RWMutex
	items  map[string]interface{}
	Indexs map[string]IndexFunc
	indices
}
