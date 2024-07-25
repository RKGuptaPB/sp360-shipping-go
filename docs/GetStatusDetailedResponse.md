# GetStatusDetailedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  This is the system generated Batch ID. | [optional] 
**GroupName** | Pointer to **string** | This indicates group name for the Batch. | [optional] 
**Name** | Pointer to **string** |  The name of the Batch. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**Import** | Pointer to [**ImportCounterStatus**](ImportCounterStatus.md) |  | [optional] 
**AddressValidation** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**Rating** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**LabelGeneration** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**VoidLabel** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**LabelDetails** | Pointer to [**GetStatusDetailedResponseLabelDetails**](GetStatusDetailedResponseLabelDetails.md) |  | [optional] 

## Methods

### NewGetStatusDetailedResponse

`func NewGetStatusDetailedResponse() *GetStatusDetailedResponse`

NewGetStatusDetailedResponse instantiates a new GetStatusDetailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusDetailedResponseWithDefaults

`func NewGetStatusDetailedResponseWithDefaults() *GetStatusDetailedResponse`

NewGetStatusDetailedResponseWithDefaults instantiates a new GetStatusDetailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *GetStatusDetailedResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *GetStatusDetailedResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *GetStatusDetailedResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *GetStatusDetailedResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetGroupName

`func (o *GetStatusDetailedResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetStatusDetailedResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetStatusDetailedResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetStatusDetailedResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetName

`func (o *GetStatusDetailedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStatusDetailedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStatusDetailedResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStatusDetailedResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetStatusDetailedResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStatusDetailedResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStatusDetailedResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStatusDetailedResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImport

`func (o *GetStatusDetailedResponse) GetImport() ImportCounterStatus`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *GetStatusDetailedResponse) GetImportOk() (*ImportCounterStatus, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *GetStatusDetailedResponse) SetImport(v ImportCounterStatus)`

SetImport sets Import field to given value.

### HasImport

`func (o *GetStatusDetailedResponse) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetAddressValidation

`func (o *GetStatusDetailedResponse) GetAddressValidation() CounterStatus`

GetAddressValidation returns the AddressValidation field if non-nil, zero value otherwise.

### GetAddressValidationOk

`func (o *GetStatusDetailedResponse) GetAddressValidationOk() (*CounterStatus, bool)`

GetAddressValidationOk returns a tuple with the AddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressValidation

`func (o *GetStatusDetailedResponse) SetAddressValidation(v CounterStatus)`

SetAddressValidation sets AddressValidation field to given value.

### HasAddressValidation

`func (o *GetStatusDetailedResponse) HasAddressValidation() bool`

HasAddressValidation returns a boolean if a field has been set.

### GetRating

`func (o *GetStatusDetailedResponse) GetRating() CounterStatus`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetStatusDetailedResponse) GetRatingOk() (*CounterStatus, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetStatusDetailedResponse) SetRating(v CounterStatus)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GetStatusDetailedResponse) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLabelGeneration

`func (o *GetStatusDetailedResponse) GetLabelGeneration() CounterStatus`

GetLabelGeneration returns the LabelGeneration field if non-nil, zero value otherwise.

### GetLabelGenerationOk

`func (o *GetStatusDetailedResponse) GetLabelGenerationOk() (*CounterStatus, bool)`

GetLabelGenerationOk returns a tuple with the LabelGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelGeneration

`func (o *GetStatusDetailedResponse) SetLabelGeneration(v CounterStatus)`

SetLabelGeneration sets LabelGeneration field to given value.

### HasLabelGeneration

`func (o *GetStatusDetailedResponse) HasLabelGeneration() bool`

HasLabelGeneration returns a boolean if a field has been set.

### GetVoidLabel

`func (o *GetStatusDetailedResponse) GetVoidLabel() CounterStatus`

GetVoidLabel returns the VoidLabel field if non-nil, zero value otherwise.

### GetVoidLabelOk

`func (o *GetStatusDetailedResponse) GetVoidLabelOk() (*CounterStatus, bool)`

GetVoidLabelOk returns a tuple with the VoidLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidLabel

`func (o *GetStatusDetailedResponse) SetVoidLabel(v CounterStatus)`

SetVoidLabel sets VoidLabel field to given value.

### HasVoidLabel

`func (o *GetStatusDetailedResponse) HasVoidLabel() bool`

HasVoidLabel returns a boolean if a field has been set.

### GetLabelDetails

`func (o *GetStatusDetailedResponse) GetLabelDetails() GetStatusDetailedResponseLabelDetails`

GetLabelDetails returns the LabelDetails field if non-nil, zero value otherwise.

### GetLabelDetailsOk

`func (o *GetStatusDetailedResponse) GetLabelDetailsOk() (*GetStatusDetailedResponseLabelDetails, bool)`

GetLabelDetailsOk returns a tuple with the LabelDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelDetails

`func (o *GetStatusDetailedResponse) SetLabelDetails(v GetStatusDetailedResponseLabelDetails)`

SetLabelDetails sets LabelDetails field to given value.

### HasLabelDetails

`func (o *GetStatusDetailedResponse) HasLabelDetails() bool`

HasLabelDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


