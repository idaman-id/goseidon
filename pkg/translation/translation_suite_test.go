package translation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"idaman.id/storage/pkg/translation"

	. "github.com/onsi/gomega"
)

func TestTranslation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Translation Suite")
}

var _ = Describe("Translation Service", func() {
	Describe("Service instance", func() {
		It("should be defined", func() {
			Expect(translation.Service).NotTo(BeNil())
		})
	})
})
