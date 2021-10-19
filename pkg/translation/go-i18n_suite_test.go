package translation_test

import (
	. "github.com/onsi/ginkgo"
	"idaman.id/storage/pkg/translation"

	. "github.com/onsi/gomega"
)

var _ = Describe("GoI18n Service", func() {
	Describe("NewGoI18nService function", func() {
		var (
			localizer translation.Localizer
		)

		It("should return GoI18nService instance", func() {
			service := translation.NewGoI18nService(localizer)

			expected := &translation.GoI18nService{}
			Expect(service).To(Equal(expected))
		})
	})

	Describe("Translate method", func() {
		var (
			localizer translation.Localizer
			service   *translation.GoI18nService
			param     translation.TranslatorDto
		)

		BeforeEach(func() {
			localizer = &MockGoI18n{
				success: true,
			}
			service = translation.NewGoI18nService(localizer)
			param = translation.TranslatorDto{
				MessageId: "FAKE_MSG_ID",
			}
		})

		When("localizer is not specified", func() {
			It("should return messageId", func() {
				service = &translation.GoI18nService{}
				res := service.Translate(param)

				Expect(res).To(Equal(param.MessageId))
			})
		})

		When("fail to localize message", func() {
			It("should return messageId", func() {
				localizer = &MockGoI18n{
					success: false,
				}
				service = translation.NewGoI18nService(localizer)
				res := service.Translate(param)

				Expect(res).To(Equal(param.MessageId))
			})
		})

		When("success localize message", func() {
			It("should return translated message", func() {
				res := service.Translate(param)

				Expect(res).To(Equal("localized"))
			})
		})
	})

})
