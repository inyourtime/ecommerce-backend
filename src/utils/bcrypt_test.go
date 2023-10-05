package utils_test

import (
	"ecommerce-backend/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Run("Hash success", func(t *testing.T) {
		plain := "godofwar"
		hash, _ := utils.Hash(plain)
		assert.NotEqual(t, plain, hash)
	})

	t.Run("Compare hash success", func(t *testing.T) {
		plain := "godofwar"
		hash, _ := utils.Hash(plain)
		result := utils.CompareHash(plain, hash)
		assert.Equal(t, true, result)
	})

	t.Run("Compare hash fail", func(t *testing.T) {
		plain := "godofwar"
		liePlain := "eieieieiei"
		hash, _ := utils.Hash(plain)
		result := utils.CompareHash(liePlain, hash)
		assert.Equal(t, false, result)
	})

	t.Run("Hash plain overflow", func(t *testing.T) {
		plain := "oversizeoversizeoversizeoversizeoversizeoversizeoversizeoversizeoversizeoversizeoversize"
		hash, err := utils.Hash(plain)
		assert.Error(t, err)
		assert.Empty(t, hash)
	})
}
