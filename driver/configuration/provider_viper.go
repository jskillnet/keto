package configuration

import (
	"fmt"
	"strings"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ory/x/corsx"
	"github.com/ory/x/tracing"
	"github.com/ory/x/viperx"
)

const ViperKeyDSN = "dsn"

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "4466")
}

type ViperProvider struct {
	l logrus.FieldLogger
}

func NewViperProvider(l logrus.FieldLogger) Provider {
	return &ViperProvider{l: l}
}

func (v *ViperProvider) ListenOn() string {
	return fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
}
func (v *ViperProvider) CORSEnabled() bool {
	return corsx.IsEnabled(v.l, "serve")
}

func (v *ViperProvider) CORSOptions() cors.Options {
	return corsx.ParseOptions(v.l, "serve")
}

func (v *ViperProvider) DSN() string {
	return viperx.GetString(v.l, ViperKeyDSN, "", "DATABASE_URL")
}

func (v *ViperProvider) TracingServiceName() string {
	return viperx.GetString(v.l, "tracing.service_name", "ORY Keto")
}

func (v *ViperProvider) TracingProvider() string {
	return viperx.GetString(v.l, "tracing.provider", "", "TRACING_PROVIDER")
}

func (v *ViperProvider) TracingJaegerConfig() *tracing.JaegerConfig {
	return &tracing.JaegerConfig{
		LocalAgentHostPort: viperx.GetString(v.l, "tracing.providers.jaeger.local_agent_address", "", "TRACING_PROVIDER_JAEGER_LOCAL_AGENT_ADDRESS"),
		SamplerType:        viperx.GetString(v.l, "tracing.providers.jaeger.sampling.type", "const", "TRACING_PROVIDER_JAEGER_SAMPLING_TYPE"),
		SamplerValue:       viperx.GetFloat64(v.l, "tracing.providers.jaeger.sampling.value", float64(1), "TRACING_PROVIDER_JAEGER_SAMPLING_VALUE"),
		SamplerServerURL:   viperx.GetString(v.l, "tracing.providers.jaeger.sampling.server_url", "", "TRACING_PROVIDER_JAEGER_SAMPLING_SERVER_URL"),
	}
}
