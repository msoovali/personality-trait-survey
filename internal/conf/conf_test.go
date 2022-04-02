package conf

import (
	"os"
	"testing"
)

func TestConf_ParseConfig(t *testing.T) {
	t.Run("envValuesNotProvided_returnsConfigWithDefaultValues", func(t *testing.T) {
		config := ParseConfig()

		if config.Addr != defaultAddr {
			t.Errorf("Expected %s, got %s", defaultAddr, config.Addr)
		}
	})

	const (
		addrValue = ":9000"
	)
	t.Run("envValuesProvided_returnsConfigWithProvidedValues", func(t *testing.T) {
		os.Setenv(addrEnvKey, addrValue)
		config := ParseConfig()

		if config.Addr != addrValue {
			t.Errorf("Expected address %s, got %s", addrValue, config.Addr)
		}
		os.Clearenv()
	})
}
