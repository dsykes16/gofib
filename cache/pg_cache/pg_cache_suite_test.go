package pg_cache_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPgCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PgCache Suite")
}
