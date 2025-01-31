/*
Shipping APIs

Testing ShipmentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sp360shipping

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func Test_sp360shipping_ShipmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShipmentAPIService CancelShipmentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.ShipmentAPI.CancelShipmentById(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService CancelStampsERR", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.CancelStampsERR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService CreateReturnLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.ShipmentAPI.CreateReturnLabel(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService CreateShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.CreateShipment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetAllShipments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetAllShipments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetCarrierAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetCarrierAccount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetCarriers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetCarriers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetCountries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetCountries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetParcelTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetParcelTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetRates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetRates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService GetSpecialServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.GetSpecialServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ReprintShipmentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.ShipmentAPI.ReprintShipmentById(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentById(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
