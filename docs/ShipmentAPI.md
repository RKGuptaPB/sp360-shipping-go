# \ShipmentAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelShipmentById**](ShipmentAPI.md#CancelShipmentById) | **Put** /api/v1/shipments/{shipmentId}/cancel | Cancel Shipment
[**CancelStampsERR**](ShipmentAPI.md#CancelStampsERR) | **Post** /api/v1/err/stamps/void | Cancel Stamps ERR
[**CreateReturnLabel**](ShipmentAPI.md#CreateReturnLabel) | **Post** /api/v1/shipments/{shipmentId}/return | Create Return label shipment
[**CreateShipment**](ShipmentAPI.md#CreateShipment) | **Post** /api/v1/shipments | Create Shipment
[**GetAllShipments**](ShipmentAPI.md#GetAllShipments) | **Get** /api/v1/shipments | Get All Shipments
[**GetCarrierAccount**](ShipmentAPI.md#GetCarrierAccount) | **Get** /api/v1/carrierAccounts | Get Carrier Accounts
[**GetCarriers**](ShipmentAPI.md#GetCarriers) | **Get** /api/v1/carriers | Get Carriers
[**GetCountries**](ShipmentAPI.md#GetCountries) | **Get** /api/v1/countries | Get Countries
[**GetParcelTypes**](ShipmentAPI.md#GetParcelTypes) | **Get** /api/v1/parcelTypes | Get Parcel Types
[**GetRates**](ShipmentAPI.md#GetRates) | **Post** /api/v1/rates | Rate Shop and Get Single Rate
[**GetServices**](ShipmentAPI.md#GetServices) | **Get** /api/v1/services | Get Services
[**GetSpecialServices**](ShipmentAPI.md#GetSpecialServices) | **Get** /api/v1/specialServices | Get Special Services
[**ReprintShipmentById**](ShipmentAPI.md#ReprintShipmentById) | **Get** /api/v1/shipments/{shipmentId}/reprint | Reprint Shipment
[**ShipmentById**](ShipmentAPI.md#ShipmentById) | **Get** /api/v1/shipments/{shipmentId} | Get Shipment by Id



## CancelShipmentById

> CancelShipment CancelShipmentById(ctx, shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

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
	shipmentId := "shipmentId_example" // string | The shipmentId, a unique identifier assigned for the Shipment.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.CancelShipmentById(context.Background(), shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.CancelShipmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelShipmentById`: CancelShipment
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.CancelShipmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | The shipmentId, a unique identifier assigned for the Shipment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentByIdRequest struct via the builder pattern


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


## CancelStampsERR

> CancelStampsResponseERR CancelStampsERR(ctx).CancelStampsRequestERR(cancelStampsRequestERR).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Cancel Stamps ERR



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
	cancelStampsRequestERR := *openapiclient.NewCancelStampsRequestERR() // CancelStampsRequestERR | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.CancelStampsERR(context.Background()).CancelStampsRequestERR(cancelStampsRequestERR).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.CancelStampsERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelStampsERR`: CancelStampsResponseERR
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.CancelStampsERR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelStampsERRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelStampsRequestERR** | [**CancelStampsRequestERR**](CancelStampsRequestERR.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**CancelStampsResponseERR**](CancelStampsResponseERR.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReturnLabel

> ReturnLabelResponse CreateReturnLabel(ctx, shipmentId).ReturnLabel(returnLabel).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Create Return label shipment



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
	returnLabel := *openapiclient.NewReturnLabel() // ReturnLabel | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.CreateReturnLabel(context.Background(), shipmentId).ReturnLabel(returnLabel).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.CreateReturnLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReturnLabel`: ReturnLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.CreateReturnLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | It specifies the shipmentId of onward shipment against which return label has to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnLabel** | [**ReturnLabel**](ReturnLabel.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**ReturnLabelResponse**](ReturnLabelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipment

> CreateShipment200Response CreateShipment(ctx).CreateShipmentRequest(createShipmentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

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
	createShipmentRequest := openapiclient.createShipment_request{ShipmentDomestic: openapiclient.NewShipmentDomestic("DOC_4X6", "SHIPPING_LABEL", *openapiclient.NewShipmentDomesticFromAddress("27 Waterview Dr", "US", "06905", "CT"), *openapiclient.NewShipmentDomesticParcel(float32(2), float32(1), float32(1), "IN", "OZ", float32(1)), "asas2223", "PKG", "PM", *openapiclient.NewShipmentDomesticToAddress("27 Waterview Dr", "US", "06905", "CT"))} // CreateShipmentRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.CreateShipment(context.Background()).CreateShipmentRequest(createShipmentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.CreateShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipment`: CreateShipment200Response
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.CreateShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShipmentRequest** | [**CreateShipmentRequest**](CreateShipmentRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**CreateShipment200Response**](CreateShipment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllShipments

> GetAllShipments GetAllShipments(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).StartDate(startDate).EndDate(endDate).Page(page).Size(size).Execute()

Get All Shipments



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	startDate := "startDate_example" // string | While searching shipments, user set a date range to get all created shipments. This indicatesthe start date of the set date range under shipment search criteria. The date format must be: YYYY-MM-DD. (optional)
	endDate := "endDate_example" // string | While searching shipments, user set a date range to get all created shipments. This indicatesthe end date of the set date range under shipment search criteria. The date format must be: YYYY-MM-DD. (optional)
	page := "page_example" // string | The page of the Shipments search result list. (optional)
	size := "size_example" // string | The size/count of the searched result list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetAllShipments(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).StartDate(startDate).EndDate(endDate).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetAllShipments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllShipments`: GetAllShipments
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetAllShipments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **startDate** | **string** | While searching shipments, user set a date range to get all created shipments. This indicatesthe start date of the set date range under shipment search criteria. The date format must be: YYYY-MM-DD. | 
 **endDate** | **string** | While searching shipments, user set a date range to get all created shipments. This indicatesthe end date of the set date range under shipment search criteria. The date format must be: YYYY-MM-DD. | 
 **page** | **string** | The page of the Shipments search result list. | 
 **size** | **string** | The size/count of the searched result list. | 

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


## GetCarrierAccount

> GetCarrierAccount200Response GetCarrierAccount(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Get Carrier Accounts



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetCarrierAccount(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetCarrierAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCarrierAccount`: GetCarrierAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetCarrierAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCarrierAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetCarrierAccount200Response**](GetCarrierAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCarriers

> Carriers GetCarriers(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Get Carriers



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetCarriers(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetCarriers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCarriers`: Carriers
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetCarriers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCarriersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**Carriers**](Carriers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> []CountriesInner GetCountries(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).Execute()

Get Countries



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	carrier := "carrier_example" // string | The carrierID, a unique identifier given to an individual carrier. (optional)
	originCountryCode := "originCountryCode_example" // string | The Source Country. The two-character ISO country code for the country where the Shipment originates. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetCountries(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []CountriesInner
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **carrier** | **string** | The carrierID, a unique identifier given to an individual carrier. | 
 **originCountryCode** | **string** | The Source Country. The two-character ISO country code for the country where the Shipment originates. | 

### Return type

[**[]CountriesInner**](CountriesInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParcelTypes

> []ParcelTypesInner GetParcelTypes(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()

Get Parcel Types



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	carrier := "carrier_example" // string | The CarrierID, a unique identifier given to an individual carrier. It can be referred from the response of Get Carriers API (optional)
	originCountryCode := "originCountryCode_example" // string | The Source Country. The two-character ISO country code for the country where the Shipment originates. (optional)
	destinationCountryCode := "destinationCountryCode_example" // string | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetParcelTypes(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetParcelTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParcelTypes`: []ParcelTypesInner
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetParcelTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParcelTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **carrier** | **string** | The CarrierID, a unique identifier given to an individual carrier. It can be referred from the response of Get Carriers API | 
 **originCountryCode** | **string** | The Source Country. The two-character ISO country code for the country where the Shipment originates. | 
 **destinationCountryCode** | **string** | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. | 

### Return type

[**[]ParcelTypesInner**](ParcelTypesInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRates

> GetRates200Response GetRates(ctx).GetRatesRequest(getRatesRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).CompactResponse(compactResponse).Execute()

Rate Shop and Get Single Rate



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
	getRatesRequest := openapiclient.getRates_request{RateShop: openapiclient.NewRateShop(*openapiclient.NewRateShopFromAddress("27 Waterview Dr", "US", "04611-3214"), *openapiclient.NewRateShopParcel(int32(2), int32(1), int32(1), "IN", "OZ", float32(1)), *openapiclient.NewSingleRateToAddress("27 Waterview Dr", "US", "01886-6502"))} // GetRatesRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	compactResponse := true // bool | The header defines if the response required is detailed or compact. When value is set to true, it will only return rates object in response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetRates(context.Background()).GetRatesRequest(getRatesRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).CompactResponse(compactResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRates`: GetRates200Response
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getRatesRequest** | [**GetRatesRequest**](GetRatesRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **compactResponse** | **bool** | The header defines if the response required is detailed or compact. When value is set to true, it will only return rates object in response. | 

### Return type

[**GetRates200Response**](GetRates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []ServicesInner GetServices(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()

Get Services



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	carrier := "carrier_example" // string | The CarrierID, a unique identifier provided to an individual carrier. It can be referred from the response of Get Carriers API (optional)
	originCountryCode := "originCountryCode_example" // string | The Source Country. The two-character ISO country code for the country where the Shipment originates. (optional)
	destinationCountryCode := "destinationCountryCode_example" // string | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetServices(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServices`: []ServicesInner
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **carrier** | **string** | The CarrierID, a unique identifier provided to an individual carrier. It can be referred from the response of Get Carriers API | 
 **originCountryCode** | **string** | The Source Country. The two-character ISO country code for the country where the Shipment originates. | 
 **destinationCountryCode** | **string** | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. | 

### Return type

[**[]ServicesInner**](ServicesInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecialServices

> SpecialServices GetSpecialServices(ctx).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Service(service).Parcel(parcel).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()

Get Special Services



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
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	service := "service_example" // string | The serviceId. It can be referred from response of `Get Services` API (optional)
	parcel := "parcel_example" // string | The parcel Id, a unique identifier named to individual package. It can be referred from response of `Get Parcel Types` API (optional)
	carrier := "carrier_example" // string | The CarrierID, a unique identifier given to an individual carrier. It can be referred from response of `Get Carriers` API (optional)
	originCountryCode := "originCountryCode_example" // string | The Source Country. The two-character ISO country code for the country where the Shipment originates. (optional)
	destinationCountryCode := "destinationCountryCode_example" // string | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.GetSpecialServices(context.Background()).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Service(service).Parcel(parcel).Carrier(carrier).OriginCountryCode(originCountryCode).DestinationCountryCode(destinationCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.GetSpecialServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecialServices`: SpecialServices
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.GetSpecialServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecialServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **service** | **string** | The serviceId. It can be referred from response of &#x60;Get Services&#x60; API | 
 **parcel** | **string** | The parcel Id, a unique identifier named to individual package. It can be referred from response of &#x60;Get Parcel Types&#x60; API | 
 **carrier** | **string** | The CarrierID, a unique identifier given to an individual carrier. It can be referred from response of &#x60;Get Carriers&#x60; API | 
 **originCountryCode** | **string** | The Source Country. The two-character ISO country code for the country where the Shipment originates. | 
 **destinationCountryCode** | **string** | The Destination Country for the Shipment. The two-character ISO country code for the country where the shipment is to be delivered. | 

### Return type

[**SpecialServices**](SpecialServices.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprintShipmentById

> ReprintShipment ReprintShipmentById(ctx, shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).CompactResponse(compactResponse).Execute()

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
	shipmentId := "shipmentId_example" // string | The shipmentId, a unique identifier assigned to the Shipment.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	compactResponse := false // bool | The header defines if the response required is detailed or compact. When value is set to true, it will only return label layout details and parcel tracking number object in response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.ReprintShipmentById(context.Background(), shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).CompactResponse(compactResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ReprintShipmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReprintShipmentById`: ReprintShipment
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ReprintShipmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | The shipmentId, a unique identifier assigned to the Shipment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReprintShipmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **compactResponse** | **bool** | The header defines if the response required is detailed or compact. When value is set to true, it will only return label layout details and parcel tracking number object in response. | 

### Return type

[**ReprintShipment**](ReprintShipment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentById

> GetSingleShipment ShipmentById(ctx, shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()

Get Shipment by Id



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
	shipmentId := "shipmentId_example" // string | The shipmentId, a unique identifier for an individual Shipment.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners. When the developer is the only business partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.ShipmentById(context.Background(), shipmentId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.ShipmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShipmentById`: GetSingleShipment
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.ShipmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | The shipmentId, a unique identifier for an individual Shipment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners. When the developer is the only business partner, this field is not required. | 

### Return type

[**GetSingleShipment**](GetSingleShipment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

