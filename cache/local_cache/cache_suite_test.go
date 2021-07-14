package local_cache_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dsykes16/gofib/cache/local_cache"
	. "github.com/dsykes16/gofib/cache/shared_tests"
)

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LocalCache Suite")
}

var _ = Describe("Local Cache Tests", func() {
	SharedCacheTests(local_cache.NewLocalCache)
})
