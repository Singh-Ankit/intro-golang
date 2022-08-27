package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Print("Learning Race Condition")

	//let's create a waitgroup which is actually going to ask main() to wait until all Goroutines have done their tasks
	//will store reference of waitgroup will pass pointer to the functions
	//WaitGroup step 1: CREATE
	wg := &sync.WaitGroup{}

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
	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine called A, this contribute to change in variable value called score")
		//add new values to score variable
		score = append(score, 1)
		//WaitGroup step 2: INTIMATE when work is done
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine called B, this also contributes to change in variable value called score")
		//add new values to score variable
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine called C, this also contributes to change in variable value called score")
		//add new values to score variable
		score = append(score, 3)
		wg.Done()
	}(wg)
	//let's notify our main method that we have enabled the waitgroup, and it must wait
	//until all waitGroups have completed their execution
	wg.Wait()
	fmt.Println(score)
	//below function actually won't work because these fucntios are not having waitgroups as input
	/*
		go func() {
			fmt.Println("Goroutine called A, this contribute to change in variable value called score")
			//add new values to score variable
			score = append(score,1)
		} ()

		go func() {
			fmt.Println("Goroutine called B, this also contributes to change in variable value called score")
			//add new values to score variable
			score = append(score,2)
		} ()
	*/

}
