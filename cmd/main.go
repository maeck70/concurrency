package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type chanData struct {
	index int
	value Data
}

func worker(ch <-chan chanData, wg *sync.WaitGroup) {
	for task := range ch {
		delay := time.Duration(rand.Float32()*2000) * time.Millisecond
		data[task.index].Calculate(2.5, delay)
		atomic.StoreInt32(&data[task.index].Result, data[task.index].Result)
	}
	wg.Done()
}

func main() {
	ch := make(chan chanData)
	var wg sync.WaitGroup

	// Launch workers
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(ch, &wg)
	}

	// Send tasks to workers
	for i, value := range data {
		ch <- chanData{i, value}
	}

	close(ch)
	wg.Wait()

	// Print results
	fmt.Println("Done\n\nResults:")
	for _, d := range data {
		d.Print()
	}
}

func (d Data) Print() {
	fmt.Println(d.Name, d.Age, d.Result)
}

func (d *Data) Calculate(factor float32, delay time.Duration) {
	d.Wait(delay)
	d.Result = int32(float32(d.Age) * factor)
	fmt.Printf("Calculated for %s: %d\n", d.Name, d.Result)
}

func (d Data) Wait(delay time.Duration) {
	fmt.Printf("...Waiting %s for %v\n", d.Name, delay)
	time.Sleep(delay)
}
