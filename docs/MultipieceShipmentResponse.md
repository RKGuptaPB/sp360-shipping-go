# MultipieceShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | Key assigned by the shipping system to the transaction. | [optional] 
**LabelLayout** | Pointer to [**[]MultipieceShipmentResponseLabelLayoutInner**](MultipieceShipmentResponseLabelLayoutInner.md) | description | [optional] 
**FromAddress** | Pointer to [**MultipieceShipmentResponseFromAddress**](MultipieceShipmentResponseFromAddress.md) |  | [optional] 
**MultiPieceRates** | Pointer to [**[]MultipieceShipmentResponseMultiPieceRatesInner**](MultipieceShipmentResponseMultiPieceRatesInner.md) | description | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | description | [optional] 
**ShipmentId** | Pointer to **string** | description | [optional] 
**ShipmentOptions** | Pointer to [**MultipieceShipmentRequestShipmentOptions**](MultipieceShipmentRequestShipmentOptions.md) |  | [optional] 
**ToAddress** | Pointer to [**MultipieceShipmentRequestToAddress**](MultipieceShipmentRequestToAddress.md) |  | [optional] 

## Methods

### NewMultipieceShipmentResponse

`func NewMultipieceShipmentResponse() *MultipieceShipmentResponse`

NewMultipieceShipmentResponse instantiates a new MultipieceShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceShipmentResponseWithDefaults

`func NewMultipieceShipmentResponseWithDefaults() *MultipieceShipmentResponse`

NewMultipieceShipmentResponseWithDefaults instantiates a new MultipieceShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *MultipieceShipmentResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MultipieceShipmentResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MultipieceShipmentResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MultipieceShipmentResponse) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetLabelLayout

`func (o *MultipieceShipmentResponse) GetLabelLayout() []MultipieceShipmentResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *MultipieceShipmentResponse) GetLabelLayoutOk() (*[]MultipieceShipmentResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *MultipieceShipmentResponse) SetLabelLayout(v []MultipieceShipmentResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *MultipieceShipmentResponse) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceShipmentResponse) GetFromAddress() MultipieceShipmentResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceShipmentResponse) GetFromAddressOk() (*MultipieceShipmentResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceShipmentResponse) SetFromAddress(v MultipieceShipmentResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceShipmentResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceRates

`func (o *MultipieceShipmentResponse) GetMultiPieceRates() []MultipieceShipmentResponseMultiPieceRatesInner`

GetMultiPieceRates returns the MultiPieceRates field if non-nil, zero value otherwise.

### GetMultiPieceRatesOk

`func (o *MultipieceShipmentResponse) GetMultiPieceRatesOk() (*[]MultipieceShipmentResponseMultiPieceRatesInner, bool)`

GetMultiPieceRatesOk returns a tuple with the MultiPieceRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceRates

`func (o *MultipieceShipmentResponse) SetMultiPieceRates(v []MultipieceShipmentResponseMultiPieceRatesInner)`

SetMultiPieceRates sets MultiPieceRates field to given value.

### HasMultiPieceRates

`func (o *MultipieceShipmentResponse) HasMultiPieceRates() bool`

HasMultiPieceRates returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *MultipieceShipmentResponse) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *MultipieceShipmentResponse) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *MultipieceShipmentResponse) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *MultipieceShipmentResponse) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetShipmentId

`func (o *MultipieceShipmentResponse) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *MultipieceShipmentResponse) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *MultipieceShipmentResponse) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *MultipieceShipmentResponse) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceShipmentResponse) GetShipmentOptions() MultipieceShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceShipmentResponse) GetShipmentOptionsOk() (*MultipieceShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceShipmentResponse) SetShipmentOptions(v MultipieceShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceShipmentResponse) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceShipmentResponse) GetToAddress() MultipieceShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceShipmentResponse) GetToAddressOk() (*MultipieceShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceShipmentResponse) SetToAddress(v MultipieceShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceShipmentResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


