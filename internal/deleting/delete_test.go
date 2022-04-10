package deleting_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/deleting"
	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/file"
	repo "idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/text"
)

/**
 * ----------------------------------------------------------------
 * storageSpy
 * ----------------------------------------------------------------
 */
type storageSpy struct {
	DeleteFileWasCalled          bool
	DeleteFileLastParamLocalPath string
	ErrorResultOfDeleteFile      error
}

func (s *storageSpy) DeleteFile(localPath string) error {
	s.DeleteFileWasCalled = true
	s.DeleteFileLastParamLocalPath = localPath
	return s.ErrorResultOfDeleteFile
}

/**
 * ----------------------------------------------------------------
 * fileRepositorySpy
 * ----------------------------------------------------------------
 */
type fileRepositorySpy struct {
	FindByIdentifierWasCalled           bool
	FindByIdentifierLastParamIdentifier string
	FileModelResultOfFindByIdentifier   *repo.FileModel
	ErrorResultOfFindByIdentifier       error

	SaveWasCalled        bool
	SaveLastParamPayload repo.SaveFileParam
	ErrorResultOfSave    error

	DeleteWasCalled      bool
	DeleteLastIdentifier string
	ErrorResultOfDelete  error
}

func (r *fileRepositorySpy) FindByIdentifier(identifier string) (*repo.FileModel, error) {
	r.FindByIdentifierWasCalled = true
	r.FindByIdentifierLastParamIdentifier = identifier
	return r.FileModelResultOfFindByIdentifier, r.ErrorResultOfFindByIdentifier
}

func (r *fileRepositorySpy) Save(payload repo.SaveFileParam) error {
	r.SaveWasCalled = true
	r.SaveLastParamPayload = payload
	return r.ErrorResultOfSave
}

func (r *fileRepositorySpy) Delete(identifier string) error {
	r.DeleteWasCalled = true
	r.DeleteLastIdentifier = identifier
	return r.ErrorResultOfDelete
}

/**
 * ----------------------------------------------------------------
 * Main
 * ----------------------------------------------------------------
 */
var _ = Describe("Test Delete Service", func() {
	var (
		fileId              string = "identifier"
		fileIdWithExtension string = "identifier.jpeg"
		theStorageSpy       *storageSpy
		repoSpy             *fileRepositorySpy
		service             deleting.DeleteService
		fileService         file.FileRemover
	)

	setFileAsExist := func() *repo.FileModel {
		justTime := time.Unix(1000000, 0)
		fileModel := &repo.FileModel{
			Id:           1,
			UniqueId:     "UniqueId",
			OriginalName: "OriginalName",
			Name:         "Name",
			Extension:    "Extension",
			Size:         1000,
			Mimetype:     "Mimetype",
			FileLocation: "FileLocation",
			FileName:     "FileName",
			CreatedAt:    &justTime,
			UpdatedAt:    &justTime,
			DeletedAt:    &justTime,
		}
		repoSpy.FileModelResultOfFindByIdentifier = fileModel
		repoSpy.ErrorResultOfFindByIdentifier = nil
		return fileModel
	}

	setFileAsNOTExist := func() error {
		notFoundError := app_error.NewNotfoundError("error message from repo")
		repoSpy.FileModelResultOfFindByIdentifier = nil
		repoSpy.ErrorResultOfFindByIdentifier = notFoundError
		return notFoundError
	}

	setFileDataDeletionSucceeded := func() {
		repoSpy.ErrorResultOfDelete = nil
	}

	setFileDataDeletionFailed := func() error {
		theError := app_error.NewUnknownError("Unexpected error happened")
		repoSpy.ErrorResultOfDelete = theError
		return theError
	}

	/**
	 * ----------------------------------------------------------------
	 * Test cases
	 * ----------------------------------------------------------------
	 */

	BeforeEach(func() {
		textService := text.NewTextService()
		fileService = file.NewFileService(textService)

		repoSpy = &fileRepositorySpy{}
		theStorageSpy = &storageSpy{}
		service = deleting.NewDeleteService(theStorageSpy, repoSpy, fileService)
	})

	It("check file exist by calls repo method with correct argument", func() {
		setFileAsNOTExist()
		service.DeleteFile(fileIdWithExtension)
		Expect(repoSpy.FindByIdentifierWasCalled).To(Equal(true))
		Expect(repoSpy.FindByIdentifierLastParamIdentifier).To(Equal(fileId))
	})

	It("return apropriate error if file not exist", func() {
		notFoundError := setFileAsNOTExist()
		// Perhaps this is code smell because how we guarantee repository would throw not-found-error object?
		// Repo is object we mock because repo is "outsider" in point of view Delete Service hence we need a contract for aligning the communication between the two.
		// But currently the mock is only based on interface contract. I guess we need DTO contract (response contract) as well.
		Expect(service.DeleteFile(fileIdWithExtension)).To(Equal(notFoundError))
	})

	It("delete file data by calls repo delete method with correct argument", func() {
		setFileAsExist()
		service.DeleteFile(fileIdWithExtension)
		Expect(repoSpy.DeleteWasCalled).To(Equal(true))
		Expect(repoSpy.DeleteLastIdentifier).To(Equal(fileId))
	})

	It("return repo's error if repo delete method return error", func() {
		setFileAsExist()
		setFileDataDeletionFailed()
		theError := app_error.NewUnknownError("Unexpected error happened")
		repoSpy.ErrorResultOfDelete = theError
		Expect(service.DeleteFile(fileIdWithExtension)).To(Equal(theError))
	})

	It("calls storage method with correct argument for delete file object", func() {
		fileModel := setFileAsExist()
		setFileDataDeletionSucceeded()
		service.DeleteFile(fileIdWithExtension)
		Expect(theStorageSpy.DeleteFileWasCalled).To(Equal(true))
		Expect(theStorageSpy.DeleteFileLastParamLocalPath).To(Equal(fileModel.FileName))
	})

	It("returns storage's error if storage delete method return error", func() {
		setFileAsExist()
		setFileDataDeletionSucceeded()
		theError := app_error.NewUnknownError("Unexpected error happened")
		theStorageSpy.ErrorResultOfDeleteFile = theError
		Expect(service.DeleteFile(fileIdWithExtension)).To(Equal(theError))
	})

	It("returns nil when success", func() {
		setFileAsExist()
		setFileDataDeletionSucceeded()
		Expect(service.DeleteFile(fileIdWithExtension)).To(BeNil())
	})
})
