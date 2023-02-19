package main

type strategy interface {
	evict(c *cache)
}
