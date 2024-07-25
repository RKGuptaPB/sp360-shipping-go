# SchedulePickupCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupIds** | Pointer to **[]string** | description | [optional] 
**Options** | Pointer to [**[]SchedulePickupCancelRequestOptionsInner**](SchedulePickupCancelRequestOptionsInner.md) | description | [optional] 

## Methods

### NewSchedulePickupCancelRequest

`func NewSchedulePickupCancelRequest() *SchedulePickupCancelRequest`

NewSchedulePickupCancelRequest instantiates a new SchedulePickupCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupCancelRequestWithDefaults

`func NewSchedulePickupCancelRequestWithDefaults() *SchedulePickupCancelRequest`

NewSchedulePickupCancelRequestWithDefaults instantiates a new SchedulePickupCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupIds

`func (o *SchedulePickupCancelRequest) GetPickupIds() []string`

GetPickupIds returns the PickupIds field if non-nil, zero value otherwise.

### GetPickupIdsOk

`func (o *SchedulePickupCancelRequest) GetPickupIdsOk() (*[]string, bool)`

GetPickupIdsOk returns a tuple with the PickupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupIds

`func (o *SchedulePickupCancelRequest) SetPickupIds(v []string)`

SetPickupIds sets PickupIds field to given value.

### HasPickupIds

`func (o *SchedulePickupCancelRequest) HasPickupIds() bool`

HasPickupIds returns a boolean if a field has been set.

### GetOptions

`func (o *SchedulePickupCancelRequest) GetOptions() []SchedulePickupCancelRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SchedulePickupCancelRequest) GetOptionsOk() (*[]SchedulePickupCancelRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SchedulePickupCancelRequest) SetOptions(v []SchedulePickupCancelRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SchedulePickupCancelRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


