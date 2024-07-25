# ShipmentInternationalBatchToAddress

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

### NewShipmentInternationalBatchToAddress

`func NewShipmentInternationalBatchToAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string, ) *ShipmentInternationalBatchToAddress`

NewShipmentInternationalBatchToAddress instantiates a new ShipmentInternationalBatchToAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalBatchToAddressWithDefaults

`func NewShipmentInternationalBatchToAddressWithDefaults() *ShipmentInternationalBatchToAddress`

NewShipmentInternationalBatchToAddressWithDefaults instantiates a new ShipmentInternationalBatchToAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ShipmentInternationalBatchToAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ShipmentInternationalBatchToAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ShipmentInternationalBatchToAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *ShipmentInternationalBatchToAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *ShipmentInternationalBatchToAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *ShipmentInternationalBatchToAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *ShipmentInternationalBatchToAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *ShipmentInternationalBatchToAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *ShipmentInternationalBatchToAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *ShipmentInternationalBatchToAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShipmentInternationalBatchToAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShipmentInternationalBatchToAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShipmentInternationalBatchToAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *ShipmentInternationalBatchToAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentInternationalBatchToAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentInternationalBatchToAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentInternationalBatchToAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ShipmentInternationalBatchToAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ShipmentInternationalBatchToAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ShipmentInternationalBatchToAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ShipmentInternationalBatchToAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *ShipmentInternationalBatchToAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ShipmentInternationalBatchToAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ShipmentInternationalBatchToAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetStateProvince

`func (o *ShipmentInternationalBatchToAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *ShipmentInternationalBatchToAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *ShipmentInternationalBatchToAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


