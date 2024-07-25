# ShipmentParcel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **float32** | Height is a part of Dimension objet where it helps determine a parcel’s girth. | 
**Length** | **float32** | Length is a part of Dimension objet having highest numeric value out of three required parameters (length, width, and height) of Dimension. It helps determine a parcel’s girth. | 
**Width** | **float32** | Width is a part of Dimension object having lowest numeric value out of three required parameters of dimension (length, width, and height). This helps determine a parcel’s girth. | 
**DimUnit** | **string** | DimUnit is a standard for measuring the physical quantities of specified dimension parameters. | 
**WeightUnit** | **string** | WeightUnit is a standard for measuring the physical quantities of specified weight. | 
**Weight** | **float32** | Weight is the measure of how heavy an object is. | 

## Methods

### NewShipmentParcel

`func NewShipmentParcel(height float32, length float32, width float32, dimUnit string, weightUnit string, weight float32, ) *ShipmentParcel`

NewShipmentParcel instantiates a new ShipmentParcel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentParcelWithDefaults

`func NewShipmentParcelWithDefaults() *ShipmentParcel`

NewShipmentParcelWithDefaults instantiates a new ShipmentParcel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ShipmentParcel) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ShipmentParcel) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ShipmentParcel) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetLength

`func (o *ShipmentParcel) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ShipmentParcel) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ShipmentParcel) SetLength(v float32)`

SetLength sets Length field to given value.


### GetWidth

`func (o *ShipmentParcel) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ShipmentParcel) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ShipmentParcel) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetDimUnit

`func (o *ShipmentParcel) GetDimUnit() string`

GetDimUnit returns the DimUnit field if non-nil, zero value otherwise.

### GetDimUnitOk

`func (o *ShipmentParcel) GetDimUnitOk() (*string, bool)`

GetDimUnitOk returns a tuple with the DimUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimUnit

`func (o *ShipmentParcel) SetDimUnit(v string)`

SetDimUnit sets DimUnit field to given value.


### GetWeightUnit

`func (o *ShipmentParcel) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ShipmentParcel) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ShipmentParcel) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetWeight

`func (o *ShipmentParcel) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ShipmentParcel) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ShipmentParcel) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


