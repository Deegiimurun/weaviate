package noop

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/entities/vectorindex/hnsw"
)

func Test_UpdateConfig(t *testing.T) {
	t.Run("hnsw: with skip==true", func(t *testing.T) {
		// the param we care about was not changed, do not error

		ind := NewIndex()
		err := ind.UpdateUserConfig(hnsw.UserConfig{
			Skip: true,
		})

		assert.Nil(t, err)
	})

	t.Run("hnsw: with skip==false", func(t *testing.T) {
		ind := NewIndex()
		err := ind.UpdateUserConfig(hnsw.UserConfig{
			Skip: false,
		})

		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "Delete and re-create")
	})

	t.Run("with unrecognized vector index config", func(t *testing.T) {
		ind := NewIndex()
		err := ind.UpdateUserConfig(nil)

		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "unrecognized vector index config")
	})
}
