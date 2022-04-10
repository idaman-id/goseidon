package deleting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeleting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deleting Suite")
}
