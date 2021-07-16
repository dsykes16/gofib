package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dsykes16/gofib"
)

var _ = Describe("Server Initialization", func() {
	It("Loads Configuration", func() {
		expected := Config{
			DBDriver:      "postgres",
			DBHost:        "localhost",
			DBUser:        "postgres",
			DBPass:        "testpass",
			DBPort:        5432,
			DBName:        "gofib",
			DBSSL:         false,
			ServerAddress: "0.0.0.0:9000",
		}
		config, err := LoadConfig("./")
		Expect(err).NotTo(HaveOccurred())
		Expect(config).To(Equal(expected))
	})
})
