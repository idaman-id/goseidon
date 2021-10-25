package validation_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo"
	"idaman.id/storage/pkg/validation"

	. "github.com/onsi/gomega"
)

func TestValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation Suite")
}

var _ = Describe("Service Variable", func() {
	Describe("Init function", func() {

		BeforeEach(func() {
			validation.Service = nil
		})

		When("already initialized", func() {
			It("should return nil", func() {
				service, _ := validation.NewGoValidator(validator.New())
				validation.Service = service

				res := validation.Init()
				Expect(res).To(BeNil())
			})
		})

		When("not initialized", func() {
			It("should be defined", func() {
				Expect(validation.Service).To(BeNil())
				validation.Init()
				Expect(validation.Service).NotTo(BeNil())
			})

			It("should return nil", func() {
				res := validation.Init()
				Expect(res).To(BeNil())
			})
		})
	})
})
