package retry

import (
	"testing"

	"github.com/kisamoto/janus/pkg/plugin"
	"github.com/kisamoto/janus/pkg/proxy"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	def := proxy.NewRouterDefinition(proxy.NewDefinition())
	err := setupRetry(def, make(plugin.Config))
	assert.NoError(t, err)

	assert.Len(t, def.Middleware(), 1)
}
