package forwarder

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/grafana/tempo/modules/distributor/forwarder/otlpgrpc"
)

const (
	OTLPGRPCBackend = "otlpgrpc"
)

type Config struct {
	Name     string          `yaml:"name"`
	Backend  string          `yaml:"backend"`
	OTLPGRPC otlpgrpc.Config `yaml:"otlpgrpc"`
}

func (cfg *Config) Validate() error {
	if cfg.Name == "" {
		return errors.New("name is empty")
	}

	switch cfg.Backend {
	case OTLPGRPCBackend:
		return cfg.OTLPGRPC.Validate()
	default:
	}

	return fmt.Errorf("%s backend is not supported", cfg.Backend)
}

type ConfigList []Config

func (cfgs ConfigList) Validate() error {
	for i, cfg := range cfgs {
		if err := cfg.Validate(); err != nil {
			return fmt.Errorf("failed to validate config at index=%d: %w", i, err)
		}
	}

	return nil
}
