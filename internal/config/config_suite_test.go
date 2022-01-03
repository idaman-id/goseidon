package config_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	config "idaman.id/storage/internal/config"
	"idaman.id/storage/internal/error"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config Service", func() {

	Context("InitConfig function", func() {
		var (
			configService config.ConfigService
		)

		BeforeEach(func() {
			configService = NewStubconfig()
		})

		When("failed load configuration", func() {
			It("should return err", func() {
				expected := errors.New("Stubed error")
				s := &StubConfig{
					loadConfigError: expected,
				}

				err := config.InitConfig(s)

				Expect(err).To(MatchError(expected))
			})
		})

		When("success initialization", func() {
			It("should return nil", func() {
				err := config.InitConfig(configService)

				Expect(err).To(BeNil())
			})

			It("should set default data", func() {
				config.InitConfig(configService)

				Expect(configService.GetString("APP_HOST")).To((Equal("localhost")))
				Expect(configService.GetInt("APP_PORT")).To((Equal(3000)))
				Expect(configService.GetString("APP_DEFAULT_LOCALE")).To((Equal("en")))
				Expect(configService.GetInt("MIN_UPLOADED_FILE")).To((Equal(1)))
				Expect(configService.GetInt("MAX_UPLOADED_FILE")).To((Equal(5)))
				Expect(configService.GetInt("MIN_FILE_SIZE")).To((Equal(1)))
				Expect(configService.GetInt("MAX_FILE_SIZE")).To((Equal(134217728)))
			})
		})
	})

	Context("NewConfig function", func() {
		var (
			provider string
		)
		BeforeEach(func() {
			provider = config.CONFIG_VIPER
		})

		When("provider is not supported", func() {
			It("should return NotfoundError", func() {
				provider = "invalid"
				config, err := config.NewConfig(provider)

				expected := &error.NotSupportedError{
					Message: error.ERROR_NOT_SUPPORTED,
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
				config, err := config.NewConfig(provider)

				Expect(config).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})
	})

})

var _ = Describe("Config Contract", func() {
	Context("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(config.CONFIG_VIPER).To(Equal("viper"))
		})
	})
})
