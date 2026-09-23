package functional

import (
	"runtime"
	"sync"
)

func Map[T, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}

func Filter[T any](items []T, keep func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

func ParallelMap[T, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))
	workers := min(len(items), runtime.GOMAXPROCS(0))
	var wg sync.WaitGroup
	wg.Add(workers)
	for worker := range workers {
		go func(start int) {
			defer wg.Done()
			for i := start; i < len(items); i += workers {
				result[i] = transform(items[i])
			}
		}(worker)
	}
	wg.Wait()
	return result
}

func GroupBy[T any, K comparable, V any](items []T, keyValue func(T) (K, V)) map[K][]V {
	result := make(map[K][]V)
	for _, item := range items {
		key, value := keyValue(item)
		result[key] = append(result[key], value)
	}
	return result
}
