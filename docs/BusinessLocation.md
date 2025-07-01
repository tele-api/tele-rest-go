# BusinessLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address of the business | 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 

## Methods

### NewBusinessLocation

`func NewBusinessLocation(address string, ) *BusinessLocation`

NewBusinessLocation instantiates a new BusinessLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessLocationWithDefaults

`func NewBusinessLocationWithDefaults() *BusinessLocation`

NewBusinessLocationWithDefaults instantiates a new BusinessLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BusinessLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessLocation) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLocation

`func (o *BusinessLocation) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BusinessLocation) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BusinessLocation) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BusinessLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


