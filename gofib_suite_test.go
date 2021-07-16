package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGofib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gofib Suite")
}
