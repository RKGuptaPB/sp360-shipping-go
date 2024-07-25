# ReprintShipmentFromAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | [optional] 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O. Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the address belongs. | [optional] 
**Company** | Pointer to **string** | The name of the company, in case if the sender address is not residential. | [optional] 
**CountryCode** | Pointer to **string** |  The country in which the address is located. Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines. | [optional] 
**Email** | Pointer to **string** | The email address of the sender. It can be person&#39;s email address or company email address (for non-residential). | [optional] 
**Name** | Pointer to **string** | Name of the sender to which this address points. | [optional] 
**Phone** | Pointer to **string** | This is sender&#39;s phone number. Enter the digits with or without spaces or hyphens. The maximum limit of characters for Phone number is 10 digits.  | [optional] 
**PostalCode** | Pointer to **string** | The Postal Code or ZIP Code of the address. For US addresses, use either the 5-digit or 9-digit ZIP Code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | [optional] 
**Residential** | Pointer to **bool** | The specified address can be Residential or Official. In case if the address is Residential, the boolean value will be &#39;true&#39;, else it will take &#39;false&#39;. | [optional] 
**StateProvince** | Pointer to **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | [optional] 

## Methods

### NewReprintShipmentFromAddress

`func NewReprintShipmentFromAddress() *ReprintShipmentFromAddress`

NewReprintShipmentFromAddress instantiates a new ReprintShipmentFromAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprintShipmentFromAddressWithDefaults

`func NewReprintShipmentFromAddressWithDefaults() *ReprintShipmentFromAddress`

NewReprintShipmentFromAddressWithDefaults instantiates a new ReprintShipmentFromAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ReprintShipmentFromAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ReprintShipmentFromAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ReprintShipmentFromAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ReprintShipmentFromAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ReprintShipmentFromAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ReprintShipmentFromAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ReprintShipmentFromAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ReprintShipmentFromAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *ReprintShipmentFromAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *ReprintShipmentFromAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *ReprintShipmentFromAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *ReprintShipmentFromAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *ReprintShipmentFromAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *ReprintShipmentFromAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *ReprintShipmentFromAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *ReprintShipmentFromAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCompany

`func (o *ReprintShipmentFromAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ReprintShipmentFromAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ReprintShipmentFromAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ReprintShipmentFromAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *ReprintShipmentFromAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ReprintShipmentFromAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ReprintShipmentFromAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ReprintShipmentFromAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *ReprintShipmentFromAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReprintShipmentFromAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReprintShipmentFromAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ReprintShipmentFromAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *ReprintShipmentFromAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReprintShipmentFromAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReprintShipmentFromAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReprintShipmentFromAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ReprintShipmentFromAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ReprintShipmentFromAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ReprintShipmentFromAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ReprintShipmentFromAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *ReprintShipmentFromAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ReprintShipmentFromAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ReprintShipmentFromAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ReprintShipmentFromAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetResidential

`func (o *ReprintShipmentFromAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *ReprintShipmentFromAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *ReprintShipmentFromAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *ReprintShipmentFromAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetStateProvince

`func (o *ReprintShipmentFromAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *ReprintShipmentFromAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *ReprintShipmentFromAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *ReprintShipmentFromAddress) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


