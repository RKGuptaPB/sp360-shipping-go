# NotFoundErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Error code(s) that appear due HTTP 404 Page or File not found. | [optional] 
**ErrorDescription** | Pointer to **string** | The HTTP 404 Not Found response status code indicates that the server cannot find the requested resource. | [optional] 
**AdditionalCode** | Pointer to **string** | A unique identifier for the error, for example 0100025, 1110017, or 1090001. | [optional] 
**AdditionalInfo** | Pointer to **string** | Additional information about the error. This error &#39;Not Found&#39; might appear due to &#x60;Shipment Not Found&#x60;, &#x60;No Shipments to close&#x60;, or &#x60;Original Transaction not found&#x60;. | [optional] 
**AdditionalParameters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotFoundErrorsInner

`func NewNotFoundErrorsInner() *NotFoundErrorsInner`

NewNotFoundErrorsInner instantiates a new NotFoundErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorsInnerWithDefaults

`func NewNotFoundErrorsInnerWithDefaults() *NotFoundErrorsInner`

NewNotFoundErrorsInnerWithDefaults instantiates a new NotFoundErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *NotFoundErrorsInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NotFoundErrorsInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NotFoundErrorsInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NotFoundErrorsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *NotFoundErrorsInner) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *NotFoundErrorsInner) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *NotFoundErrorsInner) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *NotFoundErrorsInner) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetAdditionalCode

`func (o *NotFoundErrorsInner) GetAdditionalCode() string`

GetAdditionalCode returns the AdditionalCode field if non-nil, zero value otherwise.

### GetAdditionalCodeOk

`func (o *NotFoundErrorsInner) GetAdditionalCodeOk() (*string, bool)`

GetAdditionalCodeOk returns a tuple with the AdditionalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCode

`func (o *NotFoundErrorsInner) SetAdditionalCode(v string)`

SetAdditionalCode sets AdditionalCode field to given value.

### HasAdditionalCode

`func (o *NotFoundErrorsInner) HasAdditionalCode() bool`

HasAdditionalCode returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *NotFoundErrorsInner) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *NotFoundErrorsInner) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *NotFoundErrorsInner) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *NotFoundErrorsInner) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *NotFoundErrorsInner) GetAdditionalParameters() []string`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *NotFoundErrorsInner) GetAdditionalParametersOk() (*[]string, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *NotFoundErrorsInner) SetAdditionalParameters(v []string)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *NotFoundErrorsInner) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


