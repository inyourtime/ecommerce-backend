package pkg_test

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	configs.InitConfigMock()
}

func TestS3(t *testing.T) {
	t.Run("Get c3 client success", func(t *testing.T) {
		s3Servive := pkg.NewS3Service()
		client, err := s3Servive.GetS3Client(configs.Cfg)
		assert.Nil(t, err)
		assert.NotNil(t, client)
	})

	t.Run("Get s3 client error", func(t *testing.T) {
		configs.Cfg.S3.AccountID = "asddsadasdas"
		s3Servive := pkg.NewS3Service()
		client, err := s3Servive.GetS3Client(configs.Cfg)
		_ = client
		_ = err
		// assert.Error(t, err)
		// assert.Nil(t, client)
	})
}
