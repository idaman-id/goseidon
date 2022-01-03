package config_viper_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfigViper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConfigViper Package")
}
