package compression

import (
	"github.com/go-chi/chi/middleware"
	"github.com/kisamoto/janus/pkg/plugin"
	"github.com/kisamoto/janus/pkg/proxy"
)

func init() {
	plugin.RegisterPlugin("compression", plugin.Plugin{
		Action: setupCompression,
	})
}

func setupCompression(def *proxy.RouterDefinition, rawConfig plugin.Config) error {
	def.AddMiddleware(middleware.DefaultCompress)
	return nil
}
