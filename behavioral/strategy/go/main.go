package main

/*
The context will have several strategies for evicting entries to free up memory.
The cache struct will act as the context in this case.

Evictions are only printed. No actual implementation exists in this example.
*/
func main() {
	cache := initCache()

	cache.add("a", "1")
	cache.add("b", "2")

	// Will trigger eviction due to limit of 2 values in cache by default FIFO.
	cache.add("c", "3")

	cache.evictionAlgorithm = &lfu{}

	// Will trigger eviction due to limit of 2 values in cache by added LFU.
	cache.add("c", "3")

	cache.evictionAlgorithm = &lru{}

	// Will trigger eviction due to limit of 2 values in cache by added LRU.
	cache.add("c", "3")
}
