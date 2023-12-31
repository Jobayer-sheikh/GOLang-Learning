package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// region String conversion

	i, e := strconv.Atoi("11")
	fmt.Println(i, e)

	a, r := strconv.ParseInt("321", 10, 0)
	fmt.Println(a, r)

	fmt.Println(rand.Intn(100))

	// endregion

	// region Atomic

	var atomicInt uint64
	var nonAtomicInt uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&atomicInt, 1)
				nonAtomicInt++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", atomicInt, nonAtomicInt)

	// endregion

	// region Mutex

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	// endregion
}
