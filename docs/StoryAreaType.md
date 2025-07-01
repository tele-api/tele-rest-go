# StoryAreaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the area, always “unique\\_gift” | [default to "unique_gift"]
**Latitude** | **float32** | Location latitude in degrees | 
**Longitude** | **float32** | Location longitude in degrees | 
**Address** | Pointer to [**LocationAddress**](LocationAddress.md) |  | [optional] 
**ReactionType** | [**ReactionType**](ReactionType.md) |  | 
**IsDark** | Pointer to **bool** | *Optional*. Pass *True* if the reaction area has a dark background | [optional] 
**IsFlipped** | Pointer to **bool** | *Optional*. Pass *True* if reaction area corner is flipped | [optional] 
**Url** | **string** | HTTP or tg:// URL to be opened when the area is clicked | 
**Temperature** | **float32** | Temperature, in degree Celsius | 
**Emoji** | **string** | Emoji representing the weather | 
**BackgroundColor** | **int32** | A color of the area background in the ARGB format | 
**Name** | **string** | Unique name of the gift | 

## Methods

### NewStoryAreaType

`func NewStoryAreaType(type_ string, latitude float32, longitude float32, reactionType ReactionType, url string, temperature float32, emoji string, backgroundColor int32, name string, ) *StoryAreaType`

NewStoryAreaType instantiates a new StoryAreaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryAreaTypeWithDefaults

`func NewStoryAreaTypeWithDefaults() *StoryAreaType`

NewStoryAreaTypeWithDefaults instantiates a new StoryAreaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoryAreaType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoryAreaType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoryAreaType) SetType(v string)`

SetType sets Type field to given value.


### GetLatitude

`func (o *StoryAreaType) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *StoryAreaType) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *StoryAreaType) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *StoryAreaType) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *StoryAreaType) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *StoryAreaType) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetAddress

`func (o *StoryAreaType) GetAddress() LocationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StoryAreaType) GetAddressOk() (*LocationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StoryAreaType) SetAddress(v LocationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StoryAreaType) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetReactionType

`func (o *StoryAreaType) GetReactionType() ReactionType`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *StoryAreaType) GetReactionTypeOk() (*ReactionType, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *StoryAreaType) SetReactionType(v ReactionType)`

SetReactionType sets ReactionType field to given value.


### GetIsDark

`func (o *StoryAreaType) GetIsDark() bool`

GetIsDark returns the IsDark field if non-nil, zero value otherwise.

### GetIsDarkOk

`func (o *StoryAreaType) GetIsDarkOk() (*bool, bool)`

GetIsDarkOk returns a tuple with the IsDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDark

`func (o *StoryAreaType) SetIsDark(v bool)`

SetIsDark sets IsDark field to given value.

### HasIsDark

`func (o *StoryAreaType) HasIsDark() bool`

HasIsDark returns a boolean if a field has been set.

### GetIsFlipped

`func (o *StoryAreaType) GetIsFlipped() bool`

GetIsFlipped returns the IsFlipped field if non-nil, zero value otherwise.

### GetIsFlippedOk

`func (o *StoryAreaType) GetIsFlippedOk() (*bool, bool)`

GetIsFlippedOk returns a tuple with the IsFlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlipped

`func (o *StoryAreaType) SetIsFlipped(v bool)`

SetIsFlipped sets IsFlipped field to given value.

### HasIsFlipped

`func (o *StoryAreaType) HasIsFlipped() bool`

HasIsFlipped returns a boolean if a field has been set.

### GetUrl

`func (o *StoryAreaType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoryAreaType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoryAreaType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTemperature

`func (o *StoryAreaType) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *StoryAreaType) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *StoryAreaType) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.


### GetEmoji

`func (o *StoryAreaType) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StoryAreaType) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StoryAreaType) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetBackgroundColor

`func (o *StoryAreaType) GetBackgroundColor() int32`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *StoryAreaType) GetBackgroundColorOk() (*int32, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *StoryAreaType) SetBackgroundColor(v int32)`

SetBackgroundColor sets BackgroundColor field to given value.


### GetName

`func (o *StoryAreaType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoryAreaType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoryAreaType) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


