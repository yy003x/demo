package issue

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}
