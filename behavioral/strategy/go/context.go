package main

import "fmt"

type cache struct {
	evictionAlgorithm strategy
	stored            int
	storage           map[string]string
	storageCapacity   int
}

func initCache() *cache {
	return &cache{
		evictionAlgorithm: &fifo{},
		storage:           make(map[string]string),
		stored:            0,
		storageCapacity:   2,
	}
}

func (c *cache) add(key, value string) {
	if c.stored == c.storageCapacity {
		fmt.Println("Cache full. Will evict:")
		c.evict()
	}

	c.storage[key] = value
	c.stored++
}

func (c *cache) evict() {
	c.evictionAlgorithm.evict(c)
	c.stored--
}
