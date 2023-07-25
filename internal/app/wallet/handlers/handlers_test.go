package handlers

import (
	"testing"

	"github.com/lifebackend/vp/internal/testutils"
	"github.com/stretchr/testify/suite"
)

func TestCoreSuite(t *testing.T) {
	t.Parallel()

	// nolint:exhaustivestruct
	suite.Run(t, &CoreSuite{
		ServiceTestSuite: testutils.ServiceTestSuite{},
		Client:           nil,
	})
}
