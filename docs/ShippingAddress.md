# ShippingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | Two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code | 
**State** | **string** | State, if applicable | 
**City** | **string** | City | 
**StreetLine1** | **string** | First line for the address | 
**StreetLine2** | **string** | Second line for the address | 
**PostCode** | **string** | Address post code | 

## Methods

### NewShippingAddress

`func NewShippingAddress(countryCode string, state string, city string, streetLine1 string, streetLine2 string, postCode string, ) *ShippingAddress`

NewShippingAddress instantiates a new ShippingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingAddressWithDefaults

`func NewShippingAddressWithDefaults() *ShippingAddress`

NewShippingAddressWithDefaults instantiates a new ShippingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *ShippingAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShippingAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShippingAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetState

`func (o *ShippingAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShippingAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShippingAddress) SetState(v string)`

SetState sets State field to given value.


### GetCity

`func (o *ShippingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShippingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShippingAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetStreetLine1

`func (o *ShippingAddress) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *ShippingAddress) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *ShippingAddress) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.


### GetStreetLine2

`func (o *ShippingAddress) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *ShippingAddress) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *ShippingAddress) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.


### GetPostCode

`func (o *ShippingAddress) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *ShippingAddress) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *ShippingAddress) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


