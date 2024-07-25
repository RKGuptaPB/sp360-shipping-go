# RateShopResponseParcel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **int32** | Length is always the greatest of the three dimensions. The other two dimensions are used in the calculation of the girth. | 
**Width** | **int32** | There is no strict rule as to which element is the width or the height, but typically, by convention the width is the second greatest dimension of a parcel. | 
**Height** | **int32** | By convention the height is the smallest dimension of the parcel. | 
**DimUnit** | **string** | DimUnit is a standard for measuring the physical quantities of specified dimension parameters. | 
**WeightUnit** | **string** | WeightUnit is a standard for measuring the physical quantities of specified weight. | 
**Weight** | **float32** | Weight is the measure of how heavy an object is. | 

## Methods

### NewRateShopResponseParcel

`func NewRateShopResponseParcel(length int32, width int32, height int32, dimUnit string, weightUnit string, weight float32, ) *RateShopResponseParcel`

NewRateShopResponseParcel instantiates a new RateShopResponseParcel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateShopResponseParcelWithDefaults

`func NewRateShopResponseParcelWithDefaults() *RateShopResponseParcel`

NewRateShopResponseParcelWithDefaults instantiates a new RateShopResponseParcel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *RateShopResponseParcel) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RateShopResponseParcel) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RateShopResponseParcel) SetLength(v int32)`

SetLength sets Length field to given value.


### GetWidth

`func (o *RateShopResponseParcel) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *RateShopResponseParcel) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *RateShopResponseParcel) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *RateShopResponseParcel) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *RateShopResponseParcel) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *RateShopResponseParcel) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDimUnit

`func (o *RateShopResponseParcel) GetDimUnit() string`

GetDimUnit returns the DimUnit field if non-nil, zero value otherwise.

### GetDimUnitOk

`func (o *RateShopResponseParcel) GetDimUnitOk() (*string, bool)`

GetDimUnitOk returns a tuple with the DimUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimUnit

`func (o *RateShopResponseParcel) SetDimUnit(v string)`

SetDimUnit sets DimUnit field to given value.


### GetWeightUnit

`func (o *RateShopResponseParcel) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *RateShopResponseParcel) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *RateShopResponseParcel) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetWeight

`func (o *RateShopResponseParcel) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RateShopResponseParcel) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RateShopResponseParcel) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


