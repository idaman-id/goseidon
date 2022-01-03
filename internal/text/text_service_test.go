package text_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"idaman.id/storage/internal/text"
)

var _ = Describe("Text Service", func() {
	var (
		textService text.TextService
		message     string
	)

	BeforeEach(func() {
		textService = text.NewTextService()
		message = "Long File Name .JPG"
	})

	Context("Slugify method", func() {
		When("method called", func() {
			It("should return slugified message", func() {
				slug := textService.Slugify(message)

				Expect(slug).To(Equal("long-file-name-jpg"))
			})
		})
	})

	Context("ParseString method", func() {
		var (
			param interface{}
		)

		When("param is string", func() {
			It("should return original value", func() {
				param = "original_value"
				res := textService.ParseString(param)

				Expect(res).To(Equal(param))
			})
		})

		When("param is boolean", func() {
			It("should return the string representation", func() {
				param = true
				true := textService.ParseString(param)

				param = false
				false := textService.ParseString(param)

				Expect(true).To(Equal("true"))
				Expect(false).To(Equal("false"))
			})
		})

		When("param is float", func() {
			It("should return the string representation", func() {
				param = float32(23.22)
				float32 := textService.ParseString(param)

				param = float64(23.22)
				float64 := textService.ParseString(param)

				Expect(float32).To(Equal("23.22"))
				Expect(float64).To(Equal("23.22"))
			})
		})

		When("param is uint", func() {
			It("should return the string representation", func() {
				param = uint64(500000)
				uint64 := textService.ParseString(param)

				param = uint32(500000)
				uint32 := textService.ParseString(param)

				param = uint16(10000)
				uint16 := textService.ParseString(param)

				param = uint8(255)
				uint8 := textService.ParseString(param)

				Expect(uint64).To(Equal("500000"))
				Expect(uint32).To(Equal("500000"))
				Expect(uint16).To(Equal("10000"))
				Expect(uint8).To(Equal("255"))
			})
		})

		When("param is int", func() {
			It("should return the string representation", func() {
				param = int64(-500000)
				int64 := textService.ParseString(param)

				param = int32(-500000)
				int32 := textService.ParseString(param)

				param = int16(-10000)
				int16 := textService.ParseString(param)

				param = int8(-127)
				int8 := textService.ParseString(param)

				Expect(int64).To(Equal("-500000"))
				Expect(int32).To(Equal("-500000"))
				Expect(int16).To(Equal("-10000"))
				Expect(int8).To(Equal("-127"))
			})
		})

		When("param is unsupported", func() {
			It("should return empty string", func() {
				param = []string{
					"1", "2",
				}
				slice := textService.ParseString(param)

				Expect(slice).To(Equal(""))
			})
		})

	})

})
