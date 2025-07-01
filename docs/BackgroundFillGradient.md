# BackgroundFillGradient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background fill, always “gradient” | [default to "gradient"]
**TopColor** | **int32** | Top color of the gradient in the RGB24 format | 
**BottomColor** | **int32** | Bottom color of the gradient in the RGB24 format | 
**RotationAngle** | **int32** | Clockwise rotation angle of the background fill in degrees; 0-359 | 

## Methods

### NewBackgroundFillGradient

`func NewBackgroundFillGradient(type_ string, topColor int32, bottomColor int32, rotationAngle int32, ) *BackgroundFillGradient`

NewBackgroundFillGradient instantiates a new BackgroundFillGradient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundFillGradientWithDefaults

`func NewBackgroundFillGradientWithDefaults() *BackgroundFillGradient`

NewBackgroundFillGradientWithDefaults instantiates a new BackgroundFillGradient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundFillGradient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundFillGradient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundFillGradient) SetType(v string)`

SetType sets Type field to given value.


### GetTopColor

`func (o *BackgroundFillGradient) GetTopColor() int32`

GetTopColor returns the TopColor field if non-nil, zero value otherwise.

### GetTopColorOk

`func (o *BackgroundFillGradient) GetTopColorOk() (*int32, bool)`

GetTopColorOk returns a tuple with the TopColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopColor

`func (o *BackgroundFillGradient) SetTopColor(v int32)`

SetTopColor sets TopColor field to given value.


### GetBottomColor

`func (o *BackgroundFillGradient) GetBottomColor() int32`

GetBottomColor returns the BottomColor field if non-nil, zero value otherwise.

### GetBottomColorOk

`func (o *BackgroundFillGradient) GetBottomColorOk() (*int32, bool)`

GetBottomColorOk returns a tuple with the BottomColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomColor

`func (o *BackgroundFillGradient) SetBottomColor(v int32)`

SetBottomColor sets BottomColor field to given value.


### GetRotationAngle

`func (o *BackgroundFillGradient) GetRotationAngle() int32`

GetRotationAngle returns the RotationAngle field if non-nil, zero value otherwise.

### GetRotationAngleOk

`func (o *BackgroundFillGradient) GetRotationAngleOk() (*int32, bool)`

GetRotationAngleOk returns a tuple with the RotationAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationAngle

`func (o *BackgroundFillGradient) SetRotationAngle(v int32)`

SetRotationAngle sets RotationAngle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


