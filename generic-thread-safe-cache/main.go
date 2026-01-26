/*
The "Generic Thread-Safe Cache" 🗄️
Now, let's take Generics and combine them with Concurrency. We are going to build a reusable, thread-safe in-memory cache—essentially a type-safe, mini-Redis for your internal Go apps.

📋 The Requirements:
The Struct: Create a generic struct Cache[K comparable, V any]. 📦

K must be comparable (Go requirement for map keys).

V can be any.

It should contain a map[K]V.

Concurrency Safety: Add a sync.RWMutex to the struct. 🔐

The Methods:

Set(key K, value V): Use a Write Lock (Lock()) to add the item.

Get(key K) (V, bool): Use a Read Lock (RLock()) to retrieve the item. 📖

Delete(key K): Use a Write Lock (Lock()) to remove the item.

The Test: In main():

Initialize a Cache[string, int] for something like user scores. 📈

Initialize a Cache[int, string] for something like product IDs to names. 🛒

Demonstrate that you can Set, Get, and Delete from both.
*/
package main

import (
	"fmt"
	"sync"
)

type Cache[K comparable, V any] struct {
	myMap map[K]V
	mu    sync.RWMutex
}

func NewCache[K comparable,V any] () *Cache[K,V] {
	return &Cache[K,V]{
		myMap: make(map[K]V),
	}
}

func (cache *Cache[K,V]) Set(key K, value V) {
	cache.mu.Lock()
	cache.myMap[key] = value
	cache.mu.Unlock()
}

func (cache *Cache[K,V]) Get(key K) (V, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	var returnValue V
	v,exist:= cache.myMap[key]
	if exist {
		return v, true
	} else {
		return returnValue, false
	}
}

func (cache *Cache[K,V]) Delete(key K) {
	cache.mu.Lock()
	delete(cache.myMap,key);
	cache.mu.Unlock()
}

func main() {
	userScoresCache:=NewCache[string, int]()
	userScoresCache.Set("deepak",10)
	score,scoreExist:=userScoresCache.Get("deepak")
	if scoreExist {
		fmt.Println(score)
	} else {
		fmt.Println("Score for deepak doesn't exist")
	}
	userScoresCache.Delete("deepak")
	productIdNamesCache:=NewCache[int,string]()

	productIdNamesCache.Set(1,"car")
	productName,productExist:=productIdNamesCache.Get(1)
	if productExist {
		fmt.Println(productName)
	} else {
		fmt.Println("Product with id 1 doesn't exist")
	}
	productIdNamesCache.Delete(1)
}