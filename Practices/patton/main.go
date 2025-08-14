package main

import (
	"fmt"
	"sync"
)

type DBI struct {
	Name string
}

// GetDBInstance creates a new database instance.
// This function is used as the New function for sync.Pool.
func getDBInstance() interface{} {
	fmt.Println("Creating Connection")
	return &DBI{Name: "MongoDB"}
}

func main() {
	pool := &sync.Pool{
		New: getDBInstance,
	}

	// The first call to Get will create a new instance (calls New)
	instance := pool.Get()
	fmt.Printf("First Get: %+v\n", instance)

	// Put the instance back into the pool
	pool.Put(instance)
	pool.Put(instance)
	pool.Put(instance)
	pool.Put(instance)
	pool.Put(instance)

	// Now, Get will retrieve the same instance (no "Creating Connection" printed)
	pool.Get()
	pool.Get()
	pool.Get()
	pool.Get()

	fmt.Println(pool.Get())

	
}

/*
Why it "not working" (i.e., why "Creating Connection" is printed multiple times):

- sync.Pool is designed for temporary object reuse. When you call Get and the pool is empty, it calls the New function.
- When you Put an object back, it is available for the next Get, but only one object is stored unless multiple are Put.
- After you Get and Put once, only one object is in the pool. The next Get retrieves it, but after that, the pool is empty again.
- Each subsequent Get calls New, so "Creating Connection" is printed each time.

If you want to reuse the same instance multiple times, you need to Put it back after each use, or manage a pool of multiple objects.
*/
