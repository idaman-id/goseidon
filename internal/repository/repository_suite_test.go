package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/repository"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("Repository Service", func() {

	Context("NewRepository function", func() {
		var (
			provider string
		)

		BeforeEach(func() {
			provider = repository.DATABASE_MONGO
		})

		When("provider is not supported", func() {
			It("should return NotSupportedError", func() {
				provider = "unsupported"
				repo, err := repository.NewRepository(provider)

				expected := &app_error.NotSupportedError{
					Message: app_error.ERROR_NOT_SUPPORTED,
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
