package file_test

import (
	"mime/multipart"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/text"
)

var _ = Describe("File Service", func() {
	var (
		fileService file.FileService
	)

	BeforeEach(func() {
		slugger := text.NewTextService()
		fileService = file.NewFileService(slugger)
	})

	Context("ParseOriginalName function", func() {

		var (
			fileHeader *multipart.FileHeader
		)

		BeforeEach(func() {
			fileHeader = &multipart.FileHeader{
				Filename: "image.jpeg",
			}
		})

		When("fileHeader is null", func() {
			It("should return empty string", func() {
				fileHeader = nil
				res := fileService.ParseOriginalName(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain empty file name", func() {
			It("should return empty string", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "",
				}
				res := fileService.ParseOriginalName(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader is filled", func() {
			It("should return original file name", func() {
				res := fileService.ParseOriginalName(fileHeader)

				Expect(res).To(Equal("image.jpeg"))
			})
		})

		When("fileHeader is containing capitcal word", func() {
			It("should return lowercased file name", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "Blue Dolpin.jpeg",
				}

				res := fileService.ParseOriginalName(fileHeader)

				Expect(res).To(Equal("blue dolpin.jpeg"))
			})
		})

	})

	Context("ParseName function", func() {
		var (
			fileHeader *multipart.FileHeader
		)

		BeforeEach(func() {
			fileHeader = &multipart.FileHeader{
				Filename: "image.jpeg",
			}
		})

		When("fileHeader is null", func() {
			It("should return empty string", func() {
				fileHeader = nil
				res := fileService.ParseName(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain empty file name", func() {
			It("should return empty string", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "",
				}
				res := fileService.ParseName(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain file extension", func() {
			It("should return file name without extension", func() {
				res := fileService.ParseName(fileHeader)

				Expect(res).To(Equal("image"))
			})
		})

		When("fileHeader contain multiple word seperated by spaces", func() {
			It("should return slugged file name with dashes", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "Blue Dolpin.jpeg",
				}
				res := fileService.ParseName(fileHeader)

				Expect(res).To(Equal("blue-dolpin"))
			})
		})
	})

	Context("ParseSize function", func() {
		var (
			fileHeader *multipart.FileHeader
		)

		BeforeEach(func() {
			fileHeader = &multipart.FileHeader{
				Size: 23456,
			}
		})

		When("fileHeader is null", func() {
			It("should return 0 number", func() {
				fileHeader = nil
				res := fileService.ParseSize(fileHeader)

				Expect(res).To(Equal(int64(0)))
			})
		})

		When("fileHeader is filled", func() {
			It("should return original file size", func() {
				res := fileService.ParseSize(fileHeader)

				Expect(res).To(Equal(int64(23456)))
			})
		})

		When("fileHeader is contain negative file size", func() {
			It("should return 0 number", func() {
				fileHeader = &multipart.FileHeader{
					Size: -928374,
				}
				res := fileService.ParseSize(fileHeader)

				Expect(res).To(Equal(int64(0)))
			})
		})
	})

	Context("ParseMimeType function", func() {
		var (
			fileHeader *multipart.FileHeader
		)

		BeforeEach(func() {
			fileHeader = &multipart.FileHeader{
				Header: map[string][]string{
					"Content-Type": {
						"image/jpeg",
					},
				},
			}
		})

		When("fileHeader is null", func() {
			It("should return empty string", func() {
				fileHeader = nil
				res := fileService.ParseMimeType(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader has no header", func() {
			It("should return empty string", func() {
				fileHeader = &multipart.FileHeader{}
				res := fileService.ParseMimeType(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader has no content type header", func() {
			It("should return empty string", func() {
				fileHeader = &multipart.FileHeader{
					Header: map[string][]string{
						"Content-Disposition": {},
					},
				}
				res := fileService.ParseMimeType(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain more than one content type value", func() {
			It("should return first mime type", func() {
				fileHeader = &multipart.FileHeader{
					Header: map[string][]string{
						"Content-Type": {
							"image/jpeg",
							"image/jpg",
						},
					},
				}
				res := fileService.ParseMimeType(fileHeader)

				Expect(res).To(Equal("image/jpeg"))
			})
		})

		When("fileHeader has content type", func() {
			It("should return mime type", func() {
				res := fileService.ParseMimeType(fileHeader)

				Expect(res).To(Equal("image/jpeg"))
			})
		})

	})

	Context("ParseExtension function", func() {
		var (
			fileHeader *multipart.FileHeader
		)

		BeforeEach(func() {
			fileHeader = &multipart.FileHeader{
				Filename: "Blue Dolpin.jpeg",
			}
		})

		When("fileHeader is null", func() {
			It("should return empty string", func() {
				fileHeader = nil
				res := fileService.ParseExtension(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain empty file name", func() {
			It("should return empty string", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "",
				}
				res := fileService.ParseExtension(fileHeader)

				Expect(res).To(Equal(""))
			})
		})

		When("fileHeader contain multiple dot file name", func() {
			It("should return last suffix extension", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "Blue.Dolpin.mkv.jpeg",
				}
				res := fileService.ParseExtension(fileHeader)

				Expect(res).To(Equal("jpeg"))
			})
		})

		When("fileHeader contain uppercase file extension", func() {
			It("should return lowercased file extension", func() {
				fileHeader = &multipart.FileHeader{
					Filename: "Image.JPEG",
				}
				res := fileService.ParseExtension(fileHeader)

				Expect(res).To(Equal("jpeg"))
			})
		})

	})

	Context("RemoveFileExtension function", func() {
		var (
			fileName string
		)

		BeforeEach(func() {
			fileName = "Blue Dolpin.jpeg"
		})

		When("fileName is empty string", func() {
			It("should return empty string", func() {
				fileName = ""
				res := fileService.RemoveFileExtension(fileName)

				Expect(res).To(Equal(""))
			})
		})

		When("fileName contain multiple dot", func() {
			It("should return file name without last dot suffix", func() {
				fileName = "Blue.Dolpin.mkv.jpeg"
				res := fileService.RemoveFileExtension(fileName)

				Expect(res).To(Equal("blue.dolpin.mkv"))
			})
		})

		When("fileName is uppercased", func() {
			It("should return lowercased fileName", func() {
				fileName = "UPPERCASE-FILE-NAME.jpeg"
				res := fileService.RemoveFileExtension(fileName)

				Expect(res).To(Equal("uppercase-file-name"))
			})
		})

	})

})
