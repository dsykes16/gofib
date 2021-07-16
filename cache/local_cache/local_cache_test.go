package local_cache_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/dsykes16/gofib/cache/local_cache"
	. "github.com/dsykes16/gofib/cache/shared_tests"
)

var _ = Describe("Local Cache Tests", func() {
	SharedCacheTests(local_cache.New)
})
