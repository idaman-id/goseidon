package text_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/pkg/text"
)

func TestText(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Text Suite")
}

var _ = Describe("Text Service", func() {
	var (
		message string
	)

	BeforeEach(func() {
		message = "Long File Name .JPG"
	})

	Describe("Service variable", func() {
		It("should not be nil", func() {
			Expect(text.Service).NotTo(BeNil())
		})
	})

	Describe("Slugify method", func() {
		When("method called", func() {
			It("should return slugified message", func() {
				slug := text.Service.Slugify(message)

				Expect(slug).To(Equal("long-file-name-jpg"))
			})
		})
	})

})
