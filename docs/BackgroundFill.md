# BackgroundFill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background fill, always “freeform\\_gradient” | [default to "freeform_gradient"]
**Color** | **int32** | The color of the background fill in the RGB24 format | 
**TopColor** | **int32** | Top color of the gradient in the RGB24 format | 
**BottomColor** | **int32** | Bottom color of the gradient in the RGB24 format | 
**RotationAngle** | **int32** | Clockwise rotation angle of the background fill in degrees; 0-359 | 
**Colors** | **[]int32** | A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format | 

## Methods

### NewBackgroundFill

`func NewBackgroundFill(type_ string, color int32, topColor int32, bottomColor int32, rotationAngle int32, colors []int32, ) *BackgroundFill`

NewBackgroundFill instantiates a new BackgroundFill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundFillWithDefaults

`func NewBackgroundFillWithDefaults() *BackgroundFill`

NewBackgroundFillWithDefaults instantiates a new BackgroundFill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundFill) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundFill) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundFill) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *BackgroundFill) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BackgroundFill) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BackgroundFill) SetColor(v int32)`

SetColor sets Color field to given value.


### GetTopColor

`func (o *BackgroundFill) GetTopColor() int32`

GetTopColor returns the TopColor field if non-nil, zero value otherwise.

### GetTopColorOk

`func (o *BackgroundFill) GetTopColorOk() (*int32, bool)`

GetTopColorOk returns a tuple with the TopColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopColor

`func (o *BackgroundFill) SetTopColor(v int32)`

SetTopColor sets TopColor field to given value.


### GetBottomColor

`func (o *BackgroundFill) GetBottomColor() int32`

GetBottomColor returns the BottomColor field if non-nil, zero value otherwise.

### GetBottomColorOk

`func (o *BackgroundFill) GetBottomColorOk() (*int32, bool)`

GetBottomColorOk returns a tuple with the BottomColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomColor

`func (o *BackgroundFill) SetBottomColor(v int32)`

SetBottomColor sets BottomColor field to given value.


### GetRotationAngle

`func (o *BackgroundFill) GetRotationAngle() int32`

GetRotationAngle returns the RotationAngle field if non-nil, zero value otherwise.

### GetRotationAngleOk

`func (o *BackgroundFill) GetRotationAngleOk() (*int32, bool)`

GetRotationAngleOk returns a tuple with the RotationAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationAngle

`func (o *BackgroundFill) SetRotationAngle(v int32)`

SetRotationAngle sets RotationAngle field to given value.


### GetColors

`func (o *BackgroundFill) GetColors() []int32`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *BackgroundFill) GetColorsOk() (*[]int32, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *BackgroundFill) SetColors(v []int32)`

SetColors sets Colors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


