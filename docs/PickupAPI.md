# \PickupAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPickups**](PickupAPI.md#CancelPickups) | **Put** /api/v1/pickups/cancel | Cancel Pickups
[**CancelledPickupDocument**](PickupAPI.md#CancelledPickupDocument) | **Post** /api/v1/pickups/document/cancelled | Cancelled Pickup Document
[**GetPickupDocument**](PickupAPI.md#GetPickupDocument) | **Get** /api/v1/pickups/{pickupId}/document | Get Pickup Document
[**GetPickups**](PickupAPI.md#GetPickups) | **Get** /api/v1/pickups | Get Pickups
[**SchedulePickup**](PickupAPI.md#SchedulePickup) | **Post** /api/v1/pickups | Schedule Pickup



## CancelPickups

> SchedulePickupCancelResponse CancelPickups(ctx).SchedulePickupCancelRequest(schedulePickupCancelRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Cancel Pickups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func main() {
	schedulePickupCancelRequest := *openapiclient.NewSchedulePickupCancelRequest() // SchedulePickupCancelRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupAPI.CancelPickups(context.Background()).SchedulePickupCancelRequest(schedulePickupCancelRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupAPI.CancelPickups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPickups`: SchedulePickupCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `PickupAPI.CancelPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulePickupCancelRequest** | [**SchedulePickupCancelRequest**](SchedulePickupCancelRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**SchedulePickupCancelResponse**](SchedulePickupCancelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelledPickupDocument

> GetPickupCancelledDocumentResponse CancelledPickupDocument(ctx).GetPickupCancelledDocumentRequest(getPickupCancelledDocumentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Cancelled Pickup Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func main() {
	getPickupCancelledDocumentRequest := *openapiclient.NewGetPickupCancelledDocumentRequest() // GetPickupCancelledDocumentRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupAPI.CancelledPickupDocument(context.Background()).GetPickupCancelledDocumentRequest(getPickupCancelledDocumentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupAPI.CancelledPickupDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelledPickupDocument`: GetPickupCancelledDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `PickupAPI.CancelledPickupDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelledPickupDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPickupCancelledDocumentRequest** | [**GetPickupCancelledDocumentRequest**](GetPickupCancelledDocumentRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetPickupCancelledDocumentResponse**](GetPickupCancelledDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPickupDocument

> GetPickupDocument GetPickupDocument(ctx, pickupId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Get Pickup Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func main() {
	pickupId := "pickupId_example" // string | Indicates pickupId.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupAPI.GetPickupDocument(context.Background(), pickupId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupAPI.GetPickupDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPickupDocument`: GetPickupDocument
	fmt.Fprintf(os.Stdout, "Response from `PickupAPI.GetPickupDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickupId** | **string** | Indicates pickupId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPickupDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetPickupDocument**](GetPickupDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPickups

> GetAllShipments GetPickups(ctx).Carrier(carrier).StartDate(startDate).EndDate(endDate).Status(status).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Get Pickups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func main() {
	carrier := "carrier_example" // string | Indicates CarrierID.
	startDate := "startDate_example" // string | Indicates start date.
	endDate := "endDate_example" // string | Indicates end date.
	status := "status_example" // string | Indicates status of the pickup (schedule or cancel)
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupAPI.GetPickups(context.Background()).Carrier(carrier).StartDate(startDate).EndDate(endDate).Status(status).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupAPI.GetPickups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPickups`: GetAllShipments
	fmt.Fprintf(os.Stdout, "Response from `PickupAPI.GetPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrier** | **string** | Indicates CarrierID. | 
 **startDate** | **string** | Indicates start date. | 
 **endDate** | **string** | Indicates end date. | 
 **status** | **string** | Indicates status of the pickup (schedule or cancel) | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetAllShipments**](GetAllShipments.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulePickup

> SchedulePickup200Response SchedulePickup(ctx).SchedulePickupRequest(schedulePickupRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Schedule Pickup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-shipping-go"
)

func main() {
	schedulePickupRequest := openapiclient.schedulePickup_request{SchedulePickupDHLEXPRequest: openapiclient.NewSchedulePickupDHLEXPRequest()} // SchedulePickupRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupAPI.SchedulePickup(context.Background()).SchedulePickupRequest(schedulePickupRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupAPI.SchedulePickup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulePickup`: SchedulePickup200Response
	fmt.Fprintf(os.Stdout, "Response from `PickupAPI.SchedulePickup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulePickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulePickupRequest** | [**SchedulePickupRequest**](SchedulePickupRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**SchedulePickup200Response**](SchedulePickup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

