package responsetransformer

import (
	"github.com/kisamoto/janus/pkg/plugin"
	"github.com/kisamoto/janus/pkg/proxy"
)

func init() {
	plugin.RegisterPlugin("response_transformer", plugin.Plugin{
		Action: setupResponseTransformer,
	})
}

func setupResponseTransformer(def *proxy.RouterDefinition, rawConfig plugin.Config) error {
	var config Config
	err := plugin.Decode(rawConfig, &config)
	if err != nil {
		return err
	}

	def.AddMiddleware(NewResponseTransformer(config))
	return nil
}
