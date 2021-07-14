package gofib_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRpcfib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoFib Suite")
}
