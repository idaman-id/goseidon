package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/pkg/app"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("Repository Service", func() {

	Describe("Init function", func() {

		var (
			provider string
		)

		BeforeEach(func() {
			provider = repository.DATABASE_MONGO
			repository.FileRepo = nil
		})

		When("provider is not supported", func() {
			It("should return NotSupportedError", func() {
				provider = "unsupported"
				err := repository.Init(provider)

				expected := &app.NotSupportedError{
					Message: app.STATUS_NOT_SUPPORTED,
					Context: "Database",
				}

				Expect(err).To(MatchError(expected))
			})
		})

		When("provider is supported", func() {
			It("should return nil", func() {
				err := repository.Init(provider)

				Expect(err).To(BeNil())
			})
		})

		When("not initialized", func() {
			It("should be nil", func() {
				Expect(repository.FileRepo).To(BeNil())
			})
		})

		When("initialized", func() {
			It("should not be nil", func() {
				repository.Init(provider)

				Expect(repository.FileRepo).ToNot(BeNil())
			})
		})
	})
})

var _ = Describe("Repository Contract", func() {
	Describe("Contract constant", func() {
		It("should be defined", func() {
			Expect(repository.DATABASE_MONGO).To(Equal("mongo"))
		})
	})
})
