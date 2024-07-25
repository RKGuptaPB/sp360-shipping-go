# \BatchOperationsAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkImportAPIERR**](BatchOperationsAPI.md#BulkImportAPIERR) | **Post** /api/v1/err/shipments/importUrl | Bulk Import Shipments ERR
[**CreateBulkShipmentsAPIERR**](BatchOperationsAPI.md#CreateBulkShipmentsAPIERR) | **Post** /api/v1/err/bulkShipments | Create Bulk Shipments ERR



## BulkImportAPIERR

> ShipmentBatch BulkImportAPIERR(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Bulk Import Shipments ERR



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
	resp, r, err := apiClient.BatchOperationsAPI.BulkImportAPIERR(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsAPI.BulkImportAPIERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportAPIERR`: ShipmentBatch
	fmt.Fprintf(os.Stdout, "Response from `BatchOperationsAPI.BulkImportAPIERR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportAPIERRRequest struct via the builder pattern


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


## CreateBulkShipmentsAPIERR

> BulkShipmentResponse CreateBulkShipmentsAPIERR(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Create Bulk Shipments ERR



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
	body := openapiclient.CreateBulkShipmentsAPIERR_request{CreateBulkShipmentsERRCoversheet: openapiclient.NewCreateBulkShipmentsERRCoversheet("10", "COVERSHEET", "Name_example", "abcd123", "LTR", "FCM", []openapiclient.ShipmentERRCoversheet{*openapiclient.NewShipmentERRCoversheet(*openapiclient.NewShipmentFromAddress("27 Watervw Dr", "US", "06905", "CT"), *openapiclient.NewShipmentParcel(float32(1), float32(2), float32(1), "IN", "OZ", float32(1)), "asas2223", "LTR", "FCM", *openapiclient.NewShipmentToAddress("27 Watervw Dr", "US", "06905", "CT"))})} // CreateBulkShipmentsAPIERRRequest | This is the Request body to create bulk ERR shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchOperationsAPI.CreateBulkShipmentsAPIERR(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsAPI.CreateBulkShipmentsAPIERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkShipmentsAPIERR`: BulkShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchOperationsAPI.CreateBulkShipmentsAPIERR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkShipmentsAPIERRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBulkShipmentsAPIERRRequest**](CreateBulkShipmentsAPIERRRequest.md) | This is the Request body to create bulk ERR shipments. | 
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

