package main

import (
	"fmt"
	"runtime"
)

func allocateHeapMemory(size int) *[]int {
	// Allocate a large slice on the heap
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return &data
}

func main() {
	fmt.Print("Starting Memory Allocation Example\n")

	// Create a large heap allocation
	fmt.Println("Allocating memory on the heap...")
	_ = allocateHeapMemory(10_000_000) // Allocate ~80MB (10 million integers)

	// Display memory stats before garbage collection
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)
	fmt.Println("\nMemory Stats Before Garbage Collection:")
	fmt.Printf("Allocated: %d bytes\n", memStatsBefore.Alloc)
	fmt.Printf("Heap Allocated: %d bytes\n", memStatsBefore.HeapAlloc)

	// Trigger garbage collection and observe memory stats
	fmt.Print("\nTriggering Garbage Collection...\n")
	for i := 1; i <= 3; i++ {
		runtime.GC() // Manually trigger garbage collection

		var memStatsAfter runtime.MemStats
		runtime.ReadMemStats(&memStatsAfter)

		fmt.Printf("After GC %d:\n", i)
		fmt.Printf("  Allocated: %d bytes\n", memStatsAfter.Alloc)
		fmt.Printf("  Heap Allocated: %d bytes\n\n", memStatsAfter.HeapAlloc)
	}

	fmt.Println("Program Finished!")
}
