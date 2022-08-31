package song

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	repo := &RepoMock{}
	uc := New(repo)
	require.NotEmpty(t, uc)
}
