package database_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/database"
)

var _ = Describe("Database Contract", func() {
	Context("Contract constant", func() {
		It("should contain valid value", func() {
			Expect(database.DATABASE_MYSQL).To(Equal("mysql"))
		})
	})
})
