package main

import (
	"container/list"
	"fmt"
)

type lruInstance struct {
	linkList *list.List
	hash     map[string]*list.Element
}

var lI *lruInstance

func lru() {
	lI = &lruInstance{
		linkList: list.New(),
		hash:     map[string]*list.Element{},
	}

	set("key", "value")
	v, _ := get("key")
	fmt.Println(v)
}

func set(k, v string) {
	if el, ok := lI.hash[k]; ok {
		el.Value = v
		lI.linkList.MoveToBack(el)
		return
	}
	if lI.linkList.Len() >= 10 {
		front := lI.linkList.Front()
		if front != nil {
			delete(lI.hash, front.Value.(string))
			lI.linkList.Remove(front)
		}
	}
	newNode := lI.linkList.PushBack(v)
	lI.hash[k] = newNode
}

func get(k string) (string, bool) {
	if el, ok := lI.hash[k]; ok {
		lI.linkList.MoveToBack(el)
		return el.Value.(string), ok
	}
	return "", false
}

func linkListTest() {
	linkList := list.New()
	node := linkList.PushBack("v")
	fmt.Println(node)
}
