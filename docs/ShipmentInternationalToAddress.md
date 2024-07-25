# ShipmentInternationalToAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O. Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the address belongs. | [optional] 
**CountryCode** | **string** | The two-character ISO Code of the destination country from this ISO country list.  The country in which the address is located. Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines. | 
**Name** | Pointer to **string** | Name of the recipient to which this address points. | [optional] 
**Phone** | Pointer to **string** | This is recipient&#39;s phone number. Enter the digits with or without spaces or hyphens. The maximum limit of characters for Phone number is 10 digits.  | [optional] 
**PostalCode** | **string** | The Postal Code or ZIP Code of the address. For US addresses, use either the 5-digit or 9-digit ZIP Code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | 
**StateProvince** | **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | 

## Methods

### NewShipmentInternationalToAddress

`func NewShipmentInternationalToAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string, ) *ShipmentInternationalToAddress`

NewShipmentInternationalToAddress instantiates a new ShipmentInternationalToAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalToAddressWithDefaults

`func NewShipmentInternationalToAddressWithDefaults() *ShipmentInternationalToAddress`

NewShipmentInternationalToAddressWithDefaults instantiates a new ShipmentInternationalToAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ShipmentInternationalToAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ShipmentInternationalToAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ShipmentInternationalToAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ShipmentInternationalToAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ShipmentInternationalToAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ShipmentInternationalToAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ShipmentInternationalToAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *ShipmentInternationalToAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *ShipmentInternationalToAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *ShipmentInternationalToAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *ShipmentInternationalToAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *ShipmentInternationalToAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *ShipmentInternationalToAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *ShipmentInternationalToAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *ShipmentInternationalToAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShipmentInternationalToAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShipmentInternationalToAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShipmentInternationalToAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *ShipmentInternationalToAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentInternationalToAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentInternationalToAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentInternationalToAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ShipmentInternationalToAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ShipmentInternationalToAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ShipmentInternationalToAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ShipmentInternationalToAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *ShipmentInternationalToAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ShipmentInternationalToAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ShipmentInternationalToAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetStateProvince

`func (o *ShipmentInternationalToAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *ShipmentInternationalToAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *ShipmentInternationalToAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


