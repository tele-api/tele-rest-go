# LocationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | The two-letter ISO 3166-1 alpha-2 country code of the country where the location is located | 
**State** | Pointer to **string** | *Optional*. State of the location | [optional] 
**City** | Pointer to **string** | *Optional*. City of the location | [optional] 
**Street** | Pointer to **string** | *Optional*. Street address of the location | [optional] 

## Methods

### NewLocationAddress

`func NewLocationAddress(countryCode string, ) *LocationAddress`

NewLocationAddress instantiates a new LocationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAddressWithDefaults

`func NewLocationAddressWithDefaults() *LocationAddress`

NewLocationAddressWithDefaults instantiates a new LocationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *LocationAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LocationAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LocationAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetState

`func (o *LocationAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *LocationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LocationAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreet

`func (o *LocationAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *LocationAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *LocationAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *LocationAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


