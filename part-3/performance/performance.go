package performance

import (
	"fmt"
	"runtime"
	"time"
)

// PerformanceProfile captures detailed performance metrics
type PerformanceProfile struct {
	AlgorithmName   string
	ExecutionTime   time.Duration
	MemoryAllocated uint64
	CPUUsage        float64
}

// AnalyzeAlgorithmPerformance provides comprehensive performance analysis
func AnalyzeAlgorithmPerformance(
	algorithm func(),
	iterations int,
) PerformanceProfile {
	var memStart, memEnd runtime.MemStats
	runtime.ReadMemStats(&memStart)

	startTime := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm()
	}

	executionTime := time.Since(startTime)

	runtime.ReadMemStats(&memEnd)
	memoryAllocated := memEnd.Alloc - memStart.Alloc

	return PerformanceProfile{
		ExecutionTime:   executionTime,
		MemoryAllocated: memoryAllocated,
	}
}

// PrintPerformanceReport generates a detailed performance report
func (pp *PerformanceProfile) PrintPerformanceReport() {
	fmt.Printf("Performance Report for %s:\n", pp.AlgorithmName)
	fmt.Printf("Total Execution Time: %v\n", pp.ExecutionTime)
	fmt.Printf("Memory Allocated: %d bytes\n", pp.MemoryAllocated)
}
