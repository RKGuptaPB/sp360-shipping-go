# ShipmentInternationalBatchFromAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** |  | 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressLine3** | Pointer to **string** |  | [optional] 
**CityTown** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**PostalCode** | **string** |  | 
**StateProvince** | **string** |  | 

## Methods

### NewShipmentInternationalBatchFromAddress

`func NewShipmentInternationalBatchFromAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string, ) *ShipmentInternationalBatchFromAddress`

NewShipmentInternationalBatchFromAddress instantiates a new ShipmentInternationalBatchFromAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalBatchFromAddressWithDefaults

`func NewShipmentInternationalBatchFromAddressWithDefaults() *ShipmentInternationalBatchFromAddress`

NewShipmentInternationalBatchFromAddressWithDefaults instantiates a new ShipmentInternationalBatchFromAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ShipmentInternationalBatchFromAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ShipmentInternationalBatchFromAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ShipmentInternationalBatchFromAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *ShipmentInternationalBatchFromAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *ShipmentInternationalBatchFromAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *ShipmentInternationalBatchFromAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *ShipmentInternationalBatchFromAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *ShipmentInternationalBatchFromAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *ShipmentInternationalBatchFromAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *ShipmentInternationalBatchFromAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShipmentInternationalBatchFromAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShipmentInternationalBatchFromAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShipmentInternationalBatchFromAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *ShipmentInternationalBatchFromAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentInternationalBatchFromAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentInternationalBatchFromAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentInternationalBatchFromAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ShipmentInternationalBatchFromAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ShipmentInternationalBatchFromAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ShipmentInternationalBatchFromAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ShipmentInternationalBatchFromAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *ShipmentInternationalBatchFromAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ShipmentInternationalBatchFromAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ShipmentInternationalBatchFromAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetStateProvince

`func (o *ShipmentInternationalBatchFromAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *ShipmentInternationalBatchFromAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *ShipmentInternationalBatchFromAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


