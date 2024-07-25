# \ShipmentByRateShopAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelShipmentByIdV2**](ShipmentByRateShopAPI.md#CancelShipmentByIdV2) | **Post** /api/v2/shipments/cancel | Cancel Shipment
[**CreateShipmentV2**](ShipmentByRateShopAPI.md#CreateShipmentV2) | **Post** /api/v2/shipments | Create Shipment
[**ReprintShipmentByIdV2**](ShipmentByRateShopAPI.md#ReprintShipmentByIdV2) | **Post** /api/v2/shipments/reprint | Reprint Shipment



## CancelShipmentByIdV2

> CancelShipmentV2 CancelShipmentByIdV2(ctx).ShipmentReprintCancelV2(shipmentReprintCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Cancel Shipment



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
	shipmentReprintCancelV2 := *openapiclient.NewShipmentReprintCancelV2() // ShipmentReprintCancelV2 | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners.<br /> When the developer is the only business partner, this field is not required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer Partner's location. (optional)
	xPBTransactionId := "XbxjOb5BR7lwMAe" // string | A unique Transaction ID provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.CancelShipmentByIdV2(context.Background()).ShipmentReprintCancelV2(shipmentReprintCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.CancelShipmentByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelShipmentByIdV2`: CancelShipmentV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.CancelShipmentByIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentReprintCancelV2** | [**ShipmentReprintCancelV2**](ShipmentReprintCancelV2.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners.&lt;br /&gt; When the developer is the only business partner, this field is not required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer Partner&#39;s location. | 
 **xPBTransactionId** | **string** | A unique Transaction ID provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**CancelShipmentV2**](CancelShipmentV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipmentV2

> DomesticShipmentResponseV2 CreateShipmentV2(ctx).CreateShipmentV2Request(createShipmentV2Request).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Create Shipment



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
	createShipmentV2Request := openapiclient.createShipmentV2_request{ShipmentDomesticByRateGroup: openapiclient.NewShipmentDomesticByRateGroup(*openapiclient.NewFromAddressV2("Pradeep Kumar", "638 Manitoba Ave", "Winnipeg", "MB", "R2W 2H2"), *openapiclient.NewToAddressV2("Paul Wright", "638 Manitoba Ave", "Winnipeg", "MB", "R2W 2H2"))} // CreateShipmentV2Request | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID assigned by PB, which is used in API for the communication with the strategic business partners. <br /> When the developer is the only business partner, this field is not required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer Partner's location. (optional)
	xPBTransactionId := "XbxjOb5BR7lwMAe" // string | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.CreateShipmentV2(context.Background()).CreateShipmentV2Request(createShipmentV2Request).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.CreateShipmentV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipmentV2`: DomesticShipmentResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.CreateShipmentV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipmentV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShipmentV2Request** | [**CreateShipmentV2Request**](CreateShipmentV2Request.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID assigned by PB, which is used in API for the communication with the strategic business partners. &lt;br /&gt; When the developer is the only business partner, this field is not required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer Partner&#39;s location. | 
 **xPBTransactionId** | **string** | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**DomesticShipmentResponseV2**](DomesticShipmentResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprintShipmentByIdV2

> ReprintShipmentV2 ReprintShipmentByIdV2(ctx).ShipmentReprintCancelV2(shipmentReprintCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Reprint Shipment



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
	shipmentReprintCancelV2 := *openapiclient.NewShipmentReprintCancelV2() // ShipmentReprintCancelV2 | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners. When the developer is the only business partner, this field is not required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer Partner's location. (optional)
	xPBTransactionId := "XbxjOb5BR7lwMAe" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.ReprintShipmentByIdV2(context.Background()).ShipmentReprintCancelV2(shipmentReprintCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.ReprintShipmentByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReprintShipmentByIdV2`: ReprintShipmentV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.ReprintShipmentByIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReprintShipmentByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentReprintCancelV2** | [**ShipmentReprintCancelV2**](ShipmentReprintCancelV2.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners. When the developer is the only business partner, this field is not required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer Partner&#39;s location. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**ReprintShipmentV2**](ReprintShipmentV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

