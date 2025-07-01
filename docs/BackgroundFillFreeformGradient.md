# BackgroundFillFreeformGradient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background fill, always “freeform\\_gradient” | [default to "freeform_gradient"]
**Colors** | **[]int32** | A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format | 

## Methods

### NewBackgroundFillFreeformGradient

`func NewBackgroundFillFreeformGradient(type_ string, colors []int32, ) *BackgroundFillFreeformGradient`

NewBackgroundFillFreeformGradient instantiates a new BackgroundFillFreeformGradient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundFillFreeformGradientWithDefaults

`func NewBackgroundFillFreeformGradientWithDefaults() *BackgroundFillFreeformGradient`

NewBackgroundFillFreeformGradientWithDefaults instantiates a new BackgroundFillFreeformGradient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundFillFreeformGradient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundFillFreeformGradient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundFillFreeformGradient) SetType(v string)`

SetType sets Type field to given value.


### GetColors

`func (o *BackgroundFillFreeformGradient) GetColors() []int32`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *BackgroundFillFreeformGradient) GetColorsOk() (*[]int32, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *BackgroundFillFreeformGradient) SetColors(v []int32)`

SetColors sets Colors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


