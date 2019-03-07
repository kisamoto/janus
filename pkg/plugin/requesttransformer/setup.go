package requesttransformer

import (
	"github.com/kisamoto/janus/pkg/plugin"
	"github.com/kisamoto/janus/pkg/proxy"
)

func init() {
	plugin.RegisterPlugin("request_transformer", plugin.Plugin{
		Action: setupRequestTransformer,
	})
}

func setupRequestTransformer(def *proxy.RouterDefinition, rawConfig plugin.Config) error {
	var config Config
	err := plugin.Decode(rawConfig, &config)
	if err != nil {
		return err
	}

	def.AddMiddleware(NewRequestTransformer(config))
	return nil
}
