# \MultipieceShipmentAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipieceRates**](MultipieceShipmentAPI.md#MultipieceRates) | **Post** /api/v1/multipiece/rates | Multipiece Rates
[**MultipieceShipment**](MultipieceShipmentAPI.md#MultipieceShipment) | **Post** /api/v1/multipiece/shipments | Multipiece Shipment
[**MultipieceShipmentCancel**](MultipieceShipmentAPI.md#MultipieceShipmentCancel) | **Put** /api/v1/multipiece/shipments/{shipmentId}/cancel | Cancel Multipiece Shipment
[**MultipieceShipmentReprint**](MultipieceShipmentAPI.md#MultipieceShipmentReprint) | **Get** /api/v1/multipiece/shipments/{shipmentId}/reprint | Reprint Multipiece Shipment



## MultipieceRates

> MultipieceRates200Response MultipieceRates(ctx).MultipieceRatesRequest(multipieceRatesRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Multipiece Rates



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
	multipieceRatesRequest := openapiclient.multipieceRates_request{MultipieceRateShopRequest: openapiclient.NewMultipieceRateShopRequest()} // MultipieceRatesRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceShipmentAPI.MultipieceRates(context.Background()).MultipieceRatesRequest(multipieceRatesRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceShipmentAPI.MultipieceRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceRates`: MultipieceRates200Response
	fmt.Fprintf(os.Stdout, "Response from `MultipieceShipmentAPI.MultipieceRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipieceRatesRequest** | [**MultipieceRatesRequest**](MultipieceRatesRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceRates200Response**](MultipieceRates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipment

> MultipieceShipmentResponse MultipieceShipment(ctx).MultipieceShipmentRequest(multipieceShipmentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Multipiece Shipment



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
	multipieceShipmentRequest := *openapiclient.NewMultipieceShipmentRequest() // MultipieceShipmentRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceShipmentAPI.MultipieceShipment(context.Background()).MultipieceShipmentRequest(multipieceShipmentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceShipmentAPI.MultipieceShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipment`: MultipieceShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `MultipieceShipmentAPI.MultipieceShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipieceShipmentRequest** | [**MultipieceShipmentRequest**](MultipieceShipmentRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceShipmentResponse**](MultipieceShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipmentCancel

> CancelShipment MultipieceShipmentCancel(ctx, shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Cancel Multipiece Shipment



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
	shipmentId := "shipmentId_example" // string | It specifies the shipmentId of onward shipment against which return label has to be created.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceShipmentAPI.MultipieceShipmentCancel(context.Background(), shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceShipmentAPI.MultipieceShipmentCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipmentCancel`: CancelShipment
	fmt.Fprintf(os.Stdout, "Response from `MultipieceShipmentAPI.MultipieceShipmentCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | It specifies the shipmentId of onward shipment against which return label has to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**CancelShipment**](CancelShipment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipmentReprint

> MultipieceShipmentResponse MultipieceShipmentReprint(ctx, shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Reprint Multipiece Shipment



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
	shipmentId := "shipmentId_example" // string | It specifies the shipmentId of onward shipment against which return label has to be created.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceShipmentAPI.MultipieceShipmentReprint(context.Background(), shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceShipmentAPI.MultipieceShipmentReprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipmentReprint`: MultipieceShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `MultipieceShipmentAPI.MultipieceShipmentReprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | It specifies the shipmentId of onward shipment against which return label has to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentReprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceShipmentResponse**](MultipieceShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

