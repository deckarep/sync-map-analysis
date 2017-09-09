package main

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func populateRegular(m *RegularIntMap) {

}

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func populateMap(n int, rm *RegularIntMap) {
	nums := nrand(n)
	for _, v := range nums {
		rm.Store(v, v)
	}
}

func populateSyncMap(n int, sm *sync.Map) {
	nums := nrand(n)
	for _, v := range nums {
		sm.Store(v, v)
	}
}

func BenchmarkStoreRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	b.ResetTimer()
	for _, v := range nums {
		rm.Store(v, v)
	}
}

func BenchmarkStoreSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	b.ResetTimer()
	for _, v := range nums {
		sm.Store(v, v)
	}
}

func BenchmarkDeleteRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		rm.Delete(v)
	}
}

func BenchmarkDeleteSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		sm.Delete(v)
	}
}

func BenchmarkLoadRegularFound(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rm.Load(nums[i])
	}
}

func BenchmarkLoadRegularNotFound(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rm.Load(i)
	}
}

func BenchmarkLoadSyncFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Load(nums[i])
	}
}

func BenchmarkLoadSyncNotFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Load(i)
	}
}

func benchmarkRegularStableKeys(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	rm := NewRegularIntMap()
	populateMap(b.N, rm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			for i := 0; i < n; i++ {
				rm.Load(5)
			}
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkRegularStableKeys1(b *testing.B) {
	benchmarkRegularStableKeys(b, 1)
}

func BenchmarkRegularStableKeys2(b *testing.B) {
	benchmarkRegularStableKeys(b, 2)
}

func BenchmarkRegularStableKeys4(b *testing.B) {
	benchmarkRegularStableKeys(b, 4)
}

func BenchmarkRegularStableKeys8(b *testing.B) {
	benchmarkRegularStableKeys(b, 8)
}

func BenchmarkRegularStableKeys16(b *testing.B) {
	benchmarkRegularStableKeys(b, 16)
}

func BenchmarkRegularStableKeys32(b *testing.B) {
	benchmarkRegularStableKeys(b, 32)
}

func BenchmarkRegularStableKeys64(b *testing.B) {
	benchmarkRegularStableKeys(b, 64)
}

func BenchmarkSyncStableKeys1(b *testing.B) {
	benchmarkSyncStableKeys(b, 1)
}

func BenchmarkSyncStableKeys2(b *testing.B) {
	benchmarkSyncStableKeys(b, 2)
}

func BenchmarkSyncStableKeys4(b *testing.B) {
	benchmarkSyncStableKeys(b, 4)
}

func BenchmarkSyncStableKeys8(b *testing.B) {
	benchmarkSyncStableKeys(b, 8)
}

func BenchmarkSyncStableKeys16(b *testing.B) {
	benchmarkSyncStableKeys(b, 16)
}

func BenchmarkSyncStableKeys32(b *testing.B) {
	benchmarkSyncStableKeys(b, 32)
}

func BenchmarkSyncStableKeys64(b *testing.B) {
	benchmarkSyncStableKeys(b, 64)
}

func benchmarkSyncStableKeys(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	populateSyncMap(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			for i := 0; i < n; i++ {
				sm.Load(5)
			}
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}
