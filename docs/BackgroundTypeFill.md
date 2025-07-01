# BackgroundTypeFill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background, always “fill” | [default to "fill"]
**Fill** | [**BackgroundFill**](BackgroundFill.md) |  | 
**DarkThemeDimming** | **int32** | Dimming of the background in dark themes, as a percentage; 0-100 | 

## Methods

### NewBackgroundTypeFill

`func NewBackgroundTypeFill(type_ string, fill BackgroundFill, darkThemeDimming int32, ) *BackgroundTypeFill`

NewBackgroundTypeFill instantiates a new BackgroundTypeFill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTypeFillWithDefaults

`func NewBackgroundTypeFillWithDefaults() *BackgroundTypeFill`

NewBackgroundTypeFillWithDefaults instantiates a new BackgroundTypeFill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundTypeFill) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundTypeFill) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundTypeFill) SetType(v string)`

SetType sets Type field to given value.


### GetFill

`func (o *BackgroundTypeFill) GetFill() BackgroundFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *BackgroundTypeFill) GetFillOk() (*BackgroundFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *BackgroundTypeFill) SetFill(v BackgroundFill)`

SetFill sets Fill field to given value.


### GetDarkThemeDimming

`func (o *BackgroundTypeFill) GetDarkThemeDimming() int32`

GetDarkThemeDimming returns the DarkThemeDimming field if non-nil, zero value otherwise.

### GetDarkThemeDimmingOk

`func (o *BackgroundTypeFill) GetDarkThemeDimmingOk() (*int32, bool)`

GetDarkThemeDimmingOk returns a tuple with the DarkThemeDimming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkThemeDimming

`func (o *BackgroundTypeFill) SetDarkThemeDimming(v int32)`

SetDarkThemeDimming sets DarkThemeDimming field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


