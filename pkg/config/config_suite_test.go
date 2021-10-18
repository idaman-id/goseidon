package config_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/pkg/app"
	config "idaman.id/storage/pkg/config"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config Service", func() {

	Describe("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(config.CONFIG_VIPER).To(Equal("viper"))
		})
	})

	Describe("Init function", func() {
		BeforeEach(func() {
			config.Service = NewMockconfig()
		})

		When("service is not available ", func() {
			It("should return NotfoundError", func() {
				config.Service = nil
				err := config.Init()

				expected := &app.NotFoundError{
					Message: app.STATUS_NOT_FOUND,
					Context: "Config",
				}

				Expect(err).To(MatchError(expected))
			})
		})

		When("failed load configuration", func() {
			It("should return err", func() {
				expected := errors.New("Mocked error")
				config.Service = &MockConfig{
					loadConfigError: expected,
				}

				err := config.Init()

				Expect(err).To(MatchError(expected))
			})
		})

		When("success initialization", func() {
			It("should return nil", func() {
				err := config.Init()

				Expect(err).To(BeNil())
			})

			It("should set default data", func() {
				config.Init()

				Expect(config.Service.GetString("APP_HOST")).To((Equal("localhost")))
				Expect(config.Service.GetInt("APP_PORT")).To((Equal(3000)))
				Expect(config.Service.GetString("APP_DEFAULT_LOCALE")).To((Equal("en")))
				Expect(config.Service.GetInt("MIN_UPLOADED_FILE")).To((Equal(1)))
				Expect(config.Service.GetInt("MAX_UPLOADED_FILE")).To((Equal(5)))
				Expect(config.Service.GetInt("MIN_FILE_SIZE")).To((Equal(1)))
				Expect(config.Service.GetInt("MAX_FILE_SIZE")).To((Equal(134217728)))
			})
		})
	})

	Describe("CreateConfig function", func() {
		var (
			provider string
		)
		BeforeEach(func() {
			provider = config.CONFIG_VIPER
		})

		When("provider is not supported", func() {
			It("should return NotfoundError", func() {
				provider = "invalid"
				config, err := config.CreateConfig(provider)

				expected := &app.NotSupportedError{
					Message: app.STATUS_NOT_SUPPORTED,
					Context: "Config",
				}

				Expect(config).To(BeNil())
				Expect(err).To(MatchError(expected))
			})
		})

		When("provider is supported", func() {
			It("should return valid config", func() {
				expected := &config.ViperConfig{
					FileName: ".env",
				}
				config, err := config.CreateConfig(provider)

				Expect(config).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})
	})

})
