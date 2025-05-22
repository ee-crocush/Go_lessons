package benches

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

// бенчмарк
func BenchmarkSimple(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		Simple(data, n)
	}
}
