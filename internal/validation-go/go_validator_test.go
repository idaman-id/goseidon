package validation_go_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

func TestGoValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoValidation Package")
}
