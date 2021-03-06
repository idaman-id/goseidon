package translation_test

import (
	. "github.com/onsi/ginkgo/v2"
	"idaman.id/storage/internal/translation"

	. "github.com/onsi/gomega"
)

var _ = Describe("GoI18n Service", func() {
	Context("NewGoI18nService function", func() {
		var (
			localizer translation.Localizer
		)

		It("should return GoI18nService instance", func() {
			service := translation.NewGoI18nService(localizer)

			expected := &translation.GoI18nService{}
			Expect(service).To(Equal(expected))
		})
	})

	Context("Translate method", func() {
		var (
			localizer translation.Localizer
			service   *translation.GoI18nService
			param     translation.TranslatorParam
		)

		BeforeEach(func() {
			localizer = &FakeGoI18n{
				success: true,
			}
			service = translation.NewGoI18nService(localizer)
			param = translation.TranslatorParam{
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
				localizer = &FakeGoI18n{
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
