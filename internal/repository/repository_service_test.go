package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/repository"
)

var _ = Describe("Repository Service", func() {

	Context("NewRepository function", func() {
		var (
			provider string
		)

		BeforeEach(func() {
			provider = repository.DATABASE_MONGO
		})

		When("provider is not supported", func() {
			It("should return UnsupportedError", func() {
				provider = "unsupported"
				repo, err := repository.NewRepository(provider)

				expected := &app_error.UnsupportedError{
					Message: app_error.STATUS_NOT_SUPPORTED,
					Context: "Database",
				}

				Expect(repo).To(BeNil())
				Expect(err).To(MatchError(expected))
			})
		})

		When("provider is supported", func() {
			It("should return Repository instance", func() {
				repo, err := repository.NewRepository(provider)

				Expect(repo).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})

	})
})

var _ = Describe("Repository Contract", func() {
	Context("Contract constant", func() {
		It("should be defined", func() {
			Expect(repository.DATABASE_MONGO).To(Equal("mongo"))
		})
	})
})
