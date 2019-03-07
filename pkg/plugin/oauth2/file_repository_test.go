package oauth2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileSystemRepository(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	assert.Contains(t, wd, "github.com/kisamoto/janus")

	// .../github.com/kisamoto/janus/pkg/api/../../assets/auth
	exampleAPIsPath := filepath.Join(wd, "..", "..", "..", "assets", "stubs", "auth")
	info, err := os.Stat(exampleAPIsPath)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())

	fsRepo, err := NewFileSystemRepository(exampleAPIsPath)
	assert.NoError(t, err)
	assert.NotNil(t, fsRepo)
}
