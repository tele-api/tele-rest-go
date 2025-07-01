# StoryAreaTypeWeather

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the area, always “weather” | [default to "weather"]
**Temperature** | **float32** | Temperature, in degree Celsius | 
**Emoji** | **string** | Emoji representing the weather | 
**BackgroundColor** | **int32** | A color of the area background in the ARGB format | 

## Methods

### NewStoryAreaTypeWeather

`func NewStoryAreaTypeWeather(type_ string, temperature float32, emoji string, backgroundColor int32, ) *StoryAreaTypeWeather`

NewStoryAreaTypeWeather instantiates a new StoryAreaTypeWeather object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryAreaTypeWeatherWithDefaults

`func NewStoryAreaTypeWeatherWithDefaults() *StoryAreaTypeWeather`

NewStoryAreaTypeWeatherWithDefaults instantiates a new StoryAreaTypeWeather object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoryAreaTypeWeather) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoryAreaTypeWeather) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoryAreaTypeWeather) SetType(v string)`

SetType sets Type field to given value.


### GetTemperature

`func (o *StoryAreaTypeWeather) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *StoryAreaTypeWeather) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *StoryAreaTypeWeather) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.


### GetEmoji

`func (o *StoryAreaTypeWeather) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StoryAreaTypeWeather) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StoryAreaTypeWeather) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetBackgroundColor

`func (o *StoryAreaTypeWeather) GetBackgroundColor() int32`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *StoryAreaTypeWeather) GetBackgroundColorOk() (*int32, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *StoryAreaTypeWeather) SetBackgroundColor(v int32)`

SetBackgroundColor sets BackgroundColor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


