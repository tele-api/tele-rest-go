# StoryAreaTypeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the area, always “location” | [default to "location"]
**Latitude** | **float32** | Location latitude in degrees | 
**Longitude** | **float32** | Location longitude in degrees | 
**Address** | Pointer to [**LocationAddress**](LocationAddress.md) |  | [optional] 

## Methods

### NewStoryAreaTypeLocation

`func NewStoryAreaTypeLocation(type_ string, latitude float32, longitude float32, ) *StoryAreaTypeLocation`

NewStoryAreaTypeLocation instantiates a new StoryAreaTypeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryAreaTypeLocationWithDefaults

`func NewStoryAreaTypeLocationWithDefaults() *StoryAreaTypeLocation`

NewStoryAreaTypeLocationWithDefaults instantiates a new StoryAreaTypeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoryAreaTypeLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoryAreaTypeLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoryAreaTypeLocation) SetType(v string)`

SetType sets Type field to given value.


### GetLatitude

`func (o *StoryAreaTypeLocation) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *StoryAreaTypeLocation) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *StoryAreaTypeLocation) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *StoryAreaTypeLocation) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *StoryAreaTypeLocation) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *StoryAreaTypeLocation) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetAddress

`func (o *StoryAreaTypeLocation) GetAddress() LocationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StoryAreaTypeLocation) GetAddressOk() (*LocationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StoryAreaTypeLocation) SetAddress(v LocationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StoryAreaTypeLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


