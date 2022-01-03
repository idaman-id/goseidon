package uploading_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUploading(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uploading Package")
}
