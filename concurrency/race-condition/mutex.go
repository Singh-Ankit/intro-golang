package main

import (
	"fmt"
	"sync"
)

func ImplementingMutex() {
	fmt.Print("\n LOCK with Mutex No RACE Condition")

	//let's create a waitgroup which is actually going to ask main() to wait until all Goroutines have done their tasks
	//will store reference of waitgroup will pass pointer to the functions
	//WaitGroup step 1: CREATE
	wg := &sync.WaitGroup{}

	// TO FIX RACE CONDITION USE MUTEX :: mutually exclusive lock
	// when you access a shared memory you should lock it first and then do operation, after that release lock
	// Mutex step 1: CREATE
	mut := &sync.Mutex{}

	//based on the number of total goroutines which will be spun off, wg.ADD() must be used
	//to tell how many goroutins rae there
	//WaitGroup step 2: ADD
	wg.Add(3)

	//let's create a slice of int, which will store the score
	//there will be multiple goroutines which will add values to this variable
	var score = []int{}

	//let's create multiple go routines which is going to add new inputs to the score variable
	//suppose, these goroutines are anonymous. We will pass waitgroup values to these functions
	//WaitGroup step 2: Pass
	//Mutex step 2: Pass
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine called A, this contribute to change in variable value called score")
		//Mutex step 3: LOCK, acquire lock before touching shared resource below
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		//Mutex step 4: RELEASE LOCK, acquire lock before touching shared resource below
		//WaitGroup step 2: INTIMATE when work is done
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine called B, this also contributes to change in variable value called score")
		//Mutex step 3: LOCK, acquire lock before touching shared resource below
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		//Mutex step 4: RELEASE LOCK, acquire lock before touching shared resource below
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine called C, this also contributes to change in variable value called score")
		//Mutex step 3: LOCK, acquire lock before touching shared resource below
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		//Mutex step 4: RELEASE LOCK, acquire lock before touching shared resource below
		wg.Done()
	}(wg, mut)
	//let's notify our main method that we have enabled the waitgroup, and it must wait
	//until all waitGroups have completed their execution
	wg.Wait()
	fmt.Println(score)

}
