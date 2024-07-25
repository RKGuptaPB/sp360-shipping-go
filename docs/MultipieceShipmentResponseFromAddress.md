# MultipieceShipmentResponseFromAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O. Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the address belongs. | [optional] 
**CountryCode** | **string** | The country in which the address is located. Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines. | 
**Name** | Pointer to **string** | Name of the sender to which this address points. | [optional] 
**Phone** | Pointer to **string** | This is sender&#39;s phone number. Enter the digits with or without spaces or hyphens. The maximum limit of characters for Phone number is 10 digits.  | [optional] 
**PostalCode** | **string** | The Postal Code or ZIP Code of the address. For US addresses, use either the 5-digit or 9-digit ZIP Code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | 
**StateProvince** | **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | 

## Methods

### NewMultipieceShipmentResponseFromAddress

`func NewMultipieceShipmentResponseFromAddress(addressLine1 string, countryCode string, postalCode string, stateProvince string, ) *MultipieceShipmentResponseFromAddress`

NewMultipieceShipmentResponseFromAddress instantiates a new MultipieceShipmentResponseFromAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceShipmentResponseFromAddressWithDefaults

`func NewMultipieceShipmentResponseFromAddressWithDefaults() *MultipieceShipmentResponseFromAddress`

NewMultipieceShipmentResponseFromAddressWithDefaults instantiates a new MultipieceShipmentResponseFromAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *MultipieceShipmentResponseFromAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *MultipieceShipmentResponseFromAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *MultipieceShipmentResponseFromAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *MultipieceShipmentResponseFromAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *MultipieceShipmentResponseFromAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *MultipieceShipmentResponseFromAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *MultipieceShipmentResponseFromAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *MultipieceShipmentResponseFromAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *MultipieceShipmentResponseFromAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *MultipieceShipmentResponseFromAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCountryCode

`func (o *MultipieceShipmentResponseFromAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MultipieceShipmentResponseFromAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MultipieceShipmentResponseFromAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetName

`func (o *MultipieceShipmentResponseFromAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultipieceShipmentResponseFromAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultipieceShipmentResponseFromAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MultipieceShipmentResponseFromAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *MultipieceShipmentResponseFromAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MultipieceShipmentResponseFromAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MultipieceShipmentResponseFromAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MultipieceShipmentResponseFromAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *MultipieceShipmentResponseFromAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MultipieceShipmentResponseFromAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MultipieceShipmentResponseFromAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetStateProvince

`func (o *MultipieceShipmentResponseFromAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *MultipieceShipmentResponseFromAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *MultipieceShipmentResponseFromAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


