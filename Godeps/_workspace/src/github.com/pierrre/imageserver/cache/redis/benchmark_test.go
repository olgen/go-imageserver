package redis

import (
	"testing"

	"github.com/pierrre/imageserver"
	cachetest "github.com/pierrre/imageserver/cache/_test"
	"github.com/pierrre/imageserver/testdata"
)

func BenchmarkGetSmall(b *testing.B) {
	benchmarkGet(b, testdata.Small)
}

func BenchmarkGetMedium(b *testing.B) {
	benchmarkGet(b, testdata.Medium)
}

func BenchmarkGetLarge(b *testing.B) {
	benchmarkGet(b, testdata.Large)
}

func BenchmarkGetHuge(b *testing.B) {
	benchmarkGet(b, testdata.Huge)
}

func BenchmarkGetAnimated(b *testing.B) {
	benchmarkGet(b, testdata.Animated)
}

func benchmarkGet(b *testing.B, image *imageserver.Image) {
	cache := newTestCache(b)
	defer cache.Close()

	cachetest.CacheBenchmarkGet(b, cache, image)
}
