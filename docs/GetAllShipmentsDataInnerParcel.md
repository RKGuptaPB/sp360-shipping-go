# GetAllShipmentsDataInnerParcel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **float32** | Length is a part of Dimension object having highest numeric value out of three required parameters (length, width, and height) of Dimension. It helps determine a parcel’s girth. | 
**Width** | **float32** | Width is a part of Dimension object having lowest numeric value out of three required parameters of dimension (length, width, and height). This helps determine a parcel’s girth. | 
**Height** | **float32** | Height is a part of Dimension object where it helps determine a parcel’s girth. | 
**DimUnit** | **string** | DimUnit is a standard for measuring the physical quantities of specified dimension parameters. | 
**WeightUnit** | **string** | WeightUnit is a standard for measuring the physical quantities of specified weight. | 
**Weight** | **float32** |  | 

## Methods

### NewGetAllShipmentsDataInnerParcel

`func NewGetAllShipmentsDataInnerParcel(length float32, width float32, height float32, dimUnit string, weightUnit string, weight float32, ) *GetAllShipmentsDataInnerParcel`

NewGetAllShipmentsDataInnerParcel instantiates a new GetAllShipmentsDataInnerParcel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllShipmentsDataInnerParcelWithDefaults

`func NewGetAllShipmentsDataInnerParcelWithDefaults() *GetAllShipmentsDataInnerParcel`

NewGetAllShipmentsDataInnerParcelWithDefaults instantiates a new GetAllShipmentsDataInnerParcel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *GetAllShipmentsDataInnerParcel) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *GetAllShipmentsDataInnerParcel) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *GetAllShipmentsDataInnerParcel) SetLength(v float32)`

SetLength sets Length field to given value.


### GetWidth

`func (o *GetAllShipmentsDataInnerParcel) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GetAllShipmentsDataInnerParcel) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GetAllShipmentsDataInnerParcel) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *GetAllShipmentsDataInnerParcel) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetAllShipmentsDataInnerParcel) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetAllShipmentsDataInnerParcel) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetDimUnit

`func (o *GetAllShipmentsDataInnerParcel) GetDimUnit() string`

GetDimUnit returns the DimUnit field if non-nil, zero value otherwise.

### GetDimUnitOk

`func (o *GetAllShipmentsDataInnerParcel) GetDimUnitOk() (*string, bool)`

GetDimUnitOk returns a tuple with the DimUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimUnit

`func (o *GetAllShipmentsDataInnerParcel) SetDimUnit(v string)`

SetDimUnit sets DimUnit field to given value.


### GetWeightUnit

`func (o *GetAllShipmentsDataInnerParcel) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *GetAllShipmentsDataInnerParcel) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *GetAllShipmentsDataInnerParcel) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetWeight

`func (o *GetAllShipmentsDataInnerParcel) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetAllShipmentsDataInnerParcel) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetAllShipmentsDataInnerParcel) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


