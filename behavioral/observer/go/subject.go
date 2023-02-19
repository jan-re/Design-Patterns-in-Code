package main

import "fmt"

type subject interface {
	register(observer)
	deregister(observer)
	notifyAll()
}

type eshopItem struct {
	name      string
	inStock   bool
	observers map[string]observer
}

func newEshopItem(newName string) *eshopItem {
	return &eshopItem{
		name:      newName,
		inStock:   false,
		observers: make(map[string]observer),
	}
}

func (ei *eshopItem) makeAvailable() {
	ei.inStock = true
	fmt.Printf("Item %s is now available.\n", ei.name)
	ei.notifyAll()
}

func (ei *eshopItem) register(o observer) {
	oId := o.getID()

	if _, ok := ei.observers[oId]; ok {
		fmt.Printf("Observer %s is already subscribed. Aborting.\n", oId)
	} else {
		ei.observers[oId] = o
		fmt.Printf("Observer %s subscribed.\n", oId)
	}
}

func (ei *eshopItem) deregister(o observer) {
	oId := o.getID()

	if _, ok := ei.observers[oId]; ok {
		delete(ei.observers, oId)
		fmt.Printf("Observer %s unsubscribed.\n", oId)
	} else {
		fmt.Printf("Observer %s is not subscribed. Aborting.\n", oId)
	}
}

// notifyAll leverages the update() method from the observer interface.
func (ei *eshopItem) notifyAll() {
	for _, v := range ei.observers {
		v.update(ei.name)
	}
}
