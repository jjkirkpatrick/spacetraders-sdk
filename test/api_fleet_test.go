/*
SpaceTraders API

Testing FleetAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package spacetraders_sdk

import (
	"context"
	"testing"

	openapiclient "github.com/jjkirkpatrick/spacetraders-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_spacetraders_sdk_FleetAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FleetAPIService CreateChart", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.CreateChart(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService CreateShipShipScan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.CreateShipShipScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService CreateShipSystemScan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.CreateShipSystemScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService CreateShipWaypointScan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.CreateShipWaypointScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService CreateSurvey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.CreateSurvey(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService DockShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.DockShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService ExtractResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.ExtractResources(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService ExtractResourcesWithSurvey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.ExtractResourcesWithSurvey(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetMounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.GetMounts(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetMyShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.GetMyShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetMyShipCargo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.GetMyShipCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetMyShips", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FleetAPI.GetMyShips(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetShipCooldown", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.GetShipCooldown(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService GetShipNav", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.GetShipNav(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService InstallMount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.InstallMount(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService Jettison", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.Jettison(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService JumpShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.JumpShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService NavigateShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.NavigateShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService NegotiateContract", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.NegotiateContract(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService OrbitShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.OrbitShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService PatchShipNav", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.PatchShipNav(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService PurchaseCargo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.PurchaseCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService PurchaseShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FleetAPI.PurchaseShip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService RefuelShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.RefuelShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService RemoveMount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.RemoveMount(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService SellCargo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.SellCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService ShipRefine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.ShipRefine(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService SiphonResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.SiphonResources(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService TransferCargo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.TransferCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetAPIService WarpShip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetAPI.WarpShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
