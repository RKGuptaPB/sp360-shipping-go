# \BatchAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkImportAPI**](BatchAPI.md#BulkImportAPI) | **Post** /api/v1/shipments/importUrl | Bulk Import Shipments
[**CreateBulkShipmentsAPI**](BatchAPI.md#CreateBulkShipmentsAPI) | **Post** /api/v1/bulkShipments | Create Bulk Shipments
[**GetBatchStatusAPI**](BatchAPI.md#GetBatchStatusAPI) | **Get** /api/v1/shipments/batch/{batchId}/status | Get Batch Status
[**GetShipmentDetailsForBatchAPI**](BatchAPI.md#GetShipmentDetailsForBatchAPI) | **Get** /api/v1/shipments/batch/{batchId}/shipments | Get Batch Shipment Details
[**GetSignatureImageERR**](BatchAPI.md#GetSignatureImageERR) | **Get** /api/v1/err/shipments/{shipmentId}/signaturefile | Signature Image ERR
[**ProcessBatchAPI**](BatchAPI.md#ProcessBatchAPI) | **Post** /api/v1/shipments/batch/{batchId}/process | Process Batch
[**VoidShippingLabel**](BatchAPI.md#VoidShippingLabel) | **Post** /api/v1/shipments/batch/{batchId}/void | Void Batch Shipping Labels



## BulkImportAPI

> ShipmentBatch BulkImportAPI(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Bulk Import Shipments



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
	body := *openapiclient.NewCreateBatchRequest("DOC_8X11", "SHIPPING_LABEL", "CarrierAccountId_example", "ServiceId_example", "ParcelType_example") // CreateBatchRequest |  This is the Request body to bulk import shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BulkImportAPI(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BulkImportAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportAPI`: ShipmentBatch
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BulkImportAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBatchRequest**](CreateBatchRequest.md) |  This is the Request body to bulk import shipments. | 
 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**ShipmentBatch**](ShipmentBatch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkShipmentsAPI

> BulkShipmentResponse CreateBulkShipmentsAPI(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Create Bulk Shipments



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
	body := openapiclient.CreateBulkShipmentsAPI_request{CreateBulkShipmentInternational: openapiclient.NewCreateBulkShipmentInternational("DOC_8X11", "SHIPPING_LABEL", "name", "abcd123", "PKG", "PMI", []openapiclient.ShipmentInternationalBatch{*openapiclient.NewShipmentInternationalBatch(*openapiclient.NewShipmentInternationalBatchFromAddress("27 Watervw Dr", "US", "06905", "CT"), *openapiclient.NewShipmentInternationalBatchParcel(float32(1), float32(2), float32(1), "IN", "OZ", float32(1)), "asas2223", "PKG", "PMI", *openapiclient.NewShipmentInternationalBatchToAddress("70 Hanlan RD", "CA", "L4L3P6", "ON"), *openapiclient.NewShipmentInternationalBatchCustoms(*openapiclient.NewShipmentInternationalBatchCustomsCustomsInfo("ReasonForExport_example", float32(123), "USD")))})} // CreateBulkShipmentsAPIRequest | This is the Request body to create bulk shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.CreateBulkShipmentsAPI(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.CreateBulkShipmentsAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkShipmentsAPI`: BulkShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.CreateBulkShipmentsAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkShipmentsAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBulkShipmentsAPIRequest**](CreateBulkShipmentsAPIRequest.md) | This is the Request body to create bulk shipments. | 
 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**BulkShipmentResponse**](BulkShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchStatusAPI

> GetStatusDetailedResponse GetBatchStatusAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Get Batch Status



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
	batchId := "batchId_example" // string | A unique identifier assigned to Batch which is automatically assigned by system while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatchStatusAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatchStatusAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchStatusAPI`: GetStatusDetailedResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatchStatusAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | A unique identifier assigned to Batch which is automatically assigned by system while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchStatusAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetStatusDetailedResponse**](GetStatusDetailedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentDetailsForBatchAPI

> GetShipmentsForBatch GetShipmentDetailsForBatchAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Status(status).Step(step).Execute()

Get Batch Shipment Details



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
	batchId := "batchId_example" // string | A unique identifier assigned to Batch which is automatically assigned by system while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)
	page := int32(56) // int32 | The page number to return. The default is 1. (optional)
	size := int32(56) // int32 | The number of records to return per page. The default is 20. (optional)
	status := "status_example" // string | The status of the shipment. The default is all. (optional)
	step := "step_example" // string | The step of the batch processing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetShipmentDetailsForBatchAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Status(status).Step(step).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetShipmentDetailsForBatchAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipmentDetailsForBatchAPI`: GetShipmentsForBatch
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetShipmentDetailsForBatchAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | A unique identifier assigned to Batch which is automatically assigned by system while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentDetailsForBatchAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 
 **page** | **int32** | The page number to return. The default is 1. | 
 **size** | **int32** | The number of records to return per page. The default is 20. | 
 **status** | **string** | The status of the shipment. The default is all. | 
 **step** | **string** | The step of the batch processing. | 

### Return type

[**GetShipmentsForBatch**](GetShipmentsForBatch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignatureImageERR

> SignatureFileResponse GetSignatureImageERR(ctx, shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Signature Image ERR



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
	shipmentId := "shipmentId_example" // string | This indicates the shipment identifier.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetSignatureImageERR(context.Background(), shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetSignatureImageERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignatureImageERR`: SignatureFileResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetSignatureImageERR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | This indicates the shipment identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignatureImageERRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**SignatureFileResponse**](SignatureFileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessBatchAPI

> ProcessShipmentResponse ProcessBatchAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Process Batch



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
	batchId := "batchId_example" // string | A unique identifier assigned to Batch which is automatically assigned by system while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.ProcessBatchAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.ProcessBatchAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessBatchAPI`: ProcessShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.ProcessBatchAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | A unique identifier assigned to Batch which is automatically assigned by system while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessBatchAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**ProcessShipmentResponse**](ProcessShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidShippingLabel

> VoidBatchResponse VoidShippingLabel(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).VoidBatchRequest(voidBatchRequest).Execute()

Void Batch Shipping Labels



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
	batchId := "batchId_example" // string | A unique identifier assigned to Batch which is automatically assigned by system while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)
	voidBatchRequest := *openapiclient.NewVoidBatchRequest() // VoidBatchRequest |  This is the Request body to Void batch shipping label. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.VoidShippingLabel(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).VoidBatchRequest(voidBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.VoidShippingLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidShippingLabel`: VoidBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.VoidShippingLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | A unique identifier assigned to Batch which is automatically assigned by system while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidShippingLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 
 **voidBatchRequest** | [**VoidBatchRequest**](VoidBatchRequest.md) |  This is the Request body to Void batch shipping label. | 

### Return type

[**VoidBatchResponse**](VoidBatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

