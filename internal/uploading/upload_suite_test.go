package uploading_test

import (
	"mime/multipart"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/uploading"
)

func TestUploading(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uploading Suite")
}

var _ = Describe("Upload Contract", func() {
	Context("Contract constant", func() {
		It("should be defined", func() {
			Expect(uploading.UPLOAD_FAILED).To(Equal("failed"))
			Expect(uploading.UPLOAD_SUCCESS).To(Equal("success"))
		})
	})
})

var _ = Describe("Upload Rule", func() {
	var (
		provider string
		files    []*multipart.FileHeader
		param    uploading.UploadRuleParam
	)

	BeforeEach(func() {
		provider = "local"
		file := &multipart.FileHeader{
			Size: 1024,
		}
		files = append(files, file)
		param = uploading.UploadRuleParam{
			Provider:    provider,
			FileHeaders: files,
		}
	})

	Context("SetData method", func() {
		It("should set valid file rule", func() {
			rule := uploading.UploadRule{}
			rule.SetData(param)

			Expect(len(rule.Files)).To(Equal(len(files)))
			for index, file := range files {
				Expect(rule.Files[index].Size).To(Equal(uint64(file.Size)))
			}
		})
		It("should set valid provider rule", func() {
			rule := uploading.UploadRule{}
			rule.SetData(param)

			Expect(rule.Provider).To(Equal(provider))
		})
	})
})
