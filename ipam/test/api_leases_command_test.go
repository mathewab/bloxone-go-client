/*
IP Address Management API

Testing LeasesCommandAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ipam

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func TestLeasesCommandAPIService(t *testing.T) {

	apiClient := ipam.NewAPIClient()

	t.Run("Test LeasesCommandAPIService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LeasesCommandAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
